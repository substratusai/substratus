package controller

import (
	"context"
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ptr "k8s.io/utils/pointer"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	apiv1 "github.com/substratusai/substratus/api/v1"
	"github.com/substratusai/substratus/internal/cloud"
)

const (
	ReasonLoading = "Loading"
	ReasonLoaded  = "Loaded"
)

// DatasetReconciler reconciles a Dataset object.
type DatasetReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	*ContainerReconciler

	CloudContext *cloud.Context
}

//+kubebuilder:rbac:groups=substratus.ai,resources=datasets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=substratus.ai,resources=datasets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=substratus.ai,resources=datasets/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete

func (r *DatasetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	log.Info("Reconciling Dataset")
	defer log.Info("Done reconciling Dataset")

	var dataset apiv1.Dataset
	if err := r.Get(ctx, req.NamespacedName, &dataset); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if result, err := r.ReconcileContainer(ctx, &dataset); !result.success {
		return result.Result, err
	}

	if result, err := r.reconcileData(ctx, &dataset); !result.success {
		return result.Result, err
	}

	result, err := reconcileReadiness(ctx, r.Client, &dataset, map[string]bool{
		apiv1.ConditionContainerReady: true,
		apiv1.ConditionDataReady:      true,
	})
	if result.success {
		log.Info("Dataset is ready")
	}

	return result.Result, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *DatasetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1.Dataset{}).
		Owns(&batchv1.Job{}).
		Complete(r)
}

func (r *DatasetReconciler) reconcileData(ctx context.Context, dataset *apiv1.Dataset) (result, error) {
	log := log.FromContext(ctx)

	if dataset.Status.URL != "" {
		return result{success: true}, nil
	}

	// Service account used for loading the data.
	loaderSA := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dataLoaderServiceAccountName,
			Namespace: dataset.Namespace,
		},
	}
	if _, err := controllerutil.CreateOrUpdate(ctx, r.Client, loaderSA, func() error {
		return r.authNServiceAccount(loaderSA)
	}); err != nil {
		return result{}, fmt.Errorf("failed to create or update service account: %w", err)
	}

	// Job that will run the data-loader image that was built by the previous Job.
	loadJob, err := r.loadJob(ctx, dataset)
	if err != nil {
		log.Error(err, "unable to construct data-loader Job")
		// No use in retrying...
		return result{}, nil
	}

	if result, err := reconcileJob(ctx, r.Client, dataset, loadJob, "loading"); !result.success {
		return result, err
	}

	return result{success: true}, nil
}

const (
	dataLoaderServiceAccountName = "data-loader"
)

func (r *DatasetReconciler) authNServiceAccount(sa *corev1.ServiceAccount) error {
	if sa.Annotations == nil {
		sa.Annotations = make(map[string]string)
	}
	switch name := r.CloudContext.Name; name {
	case cloud.GCP:
		sa.Annotations["iam.gke.io/gcp-service-account"] = fmt.Sprintf("substratus-%s@%s.iam.gserviceaccount.com", sa.GetName(), r.CloudContext.GCP.ProjectID)
	default:
		return fmt.Errorf("unsupported cloud type: %q", name)
	}
	return nil
}

func (r *DatasetReconciler) loadJob(ctx context.Context, dataset *apiv1.Dataset) (*batchv1.Job, error) {
	const loaderContainerName = "loader"
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: dataset.Name + "-data-loader",
			// Cross-Namespace owners not allowed, must be same as dataset:
			Namespace: dataset.Namespace,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"kubectl.kubernetes.io/default-container": loaderContainerName,
					},
				},
				Spec: corev1.PodSpec{
					SecurityContext: &corev1.PodSecurityContext{
						RunAsUser:  ptr.Int64(1001),
						RunAsGroup: ptr.Int64(2002),
						FSGroup:    ptr.Int64(3003),
					},
					ServiceAccountName: dataLoaderServiceAccountName,
					Containers: []corev1.Container{
						{
							Name:  loaderContainerName,
							Image: dataset.Spec.Container.Image,
							Args:  []string{"load.sh"},
							Env: []corev1.EnvVar{
								{
									Name:  "LOAD_DATA_PATH",
									Value: "/data/" + dataset.Spec.Filename,
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "data",
									MountPath: "/data",
									SubPath:   string(dataset.UID) + "/data",
								},
								{
									Name:      "data",
									MountPath: "/dataset/logs",
									SubPath:   string(dataset.UID) + "/logs",
								},
							},
						},
					},
					Volumes:       []corev1.Volume{},
					RestartPolicy: "Never",
				},
			},
		},
	}

	switch r.CloudContext.Name {
	case cloud.GCP:
		// GKE will injects GCS Fuse sidecar based on this annotation.
		job.Spec.Template.Annotations["gke-gcsfuse/volumes"] = "true"
		job.Spec.Template.Spec.Volumes = append(job.Spec.Template.Spec.Volumes, corev1.Volume{
			Name: "data",
			VolumeSource: corev1.VolumeSource{
				CSI: &corev1.CSIVolumeSource{
					Driver: "gcsfuse.csi.storage.gke.io",
					VolumeAttributes: map[string]string{
						"bucketName":   r.CloudContext.GCP.ProjectID + "-substratus-datasets",
						"mountOptions": "implicit-dirs,uid=1001,gid=3003",
					},
				},
			},
		})
		dataset.Status.URL = "gcs://" + r.CloudContext.GCP.ProjectID + "-substratus-datasets" +
			"/" + string(dataset.UID) + "/data/" + dataset.Spec.Filename
	}

	if err := controllerutil.SetControllerReference(dataset, job, r.Scheme); err != nil {
		return nil, fmt.Errorf("setting owner reference: %w", err)
	}

	return job, nil
}
