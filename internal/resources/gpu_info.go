package resources

import (
	apiv1 "github.com/substratusai/substratus/api/v1"
	"github.com/substratusai/substratus/internal/cloud"
	corev1 "k8s.io/api/core/v1"
)

type GPUInfo struct {
	Memory       int64
	ResourceName corev1.ResourceName
	NodeSelector map[string]string
}

var cloudGPUs = map[cloud.Name]map[apiv1.GPUType]*GPUInfo{
	cloud.GCP: {
		// https://cloud.google.com/compute/docs/gpus#nvidia_t4_gpus
		apiv1.GPUTypeNvidiaTeslaT4: {
			Memory:       16 * gigabyte,
			ResourceName: corev1.ResourceName("nvidia.com/gpu"),
			NodeSelector: map[string]string{
				"cloud.google.com/gke-accelerator": "nvidia-tesla-t4",
			},
		},
		// https://cloud.google.com/compute/docs/gpus#l4-gpus
		apiv1.GPUTypeNvidiaL4: {
			Memory:       24 * gigabyte,
			ResourceName: corev1.ResourceName("nvidia.com/gpu"),
			NodeSelector: map[string]string{
				"cloud.google.com/gke-accelerator": "nvidia-l4",
			},
		},
	},
}
