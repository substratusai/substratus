//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dataset) DeepCopyInto(out *Dataset) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dataset.
func (in *Dataset) DeepCopy() *Dataset {
	if in == nil {
		return nil
	}
	out := new(Dataset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Dataset) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetList) DeepCopyInto(out *DatasetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Dataset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetList.
func (in *DatasetList) DeepCopy() *DatasetList {
	if in == nil {
		return nil
	}
	out := new(DatasetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetSource) DeepCopyInto(out *DatasetSource) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitSource)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetSource.
func (in *DatasetSource) DeepCopy() *DatasetSource {
	if in == nil {
		return nil
	}
	out := new(DatasetSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetSpec) DeepCopyInto(out *DatasetSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetSpec.
func (in *DatasetSpec) DeepCopy() *DatasetSpec {
	if in == nil {
		return nil
	}
	out := new(DatasetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetStatus) DeepCopyInto(out *DatasetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetStatus.
func (in *DatasetStatus) DeepCopy() *DatasetStatus {
	if in == nil {
		return nil
	}
	out := new(DatasetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitSource) DeepCopyInto(out *GitSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitSource.
func (in *GitSource) DeepCopy() *GitSource {
	if in == nil {
		return nil
	}
	out := new(GitSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Model) DeepCopyInto(out *Model) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Model.
func (in *Model) DeepCopy() *Model {
	if in == nil {
		return nil
	}
	out := new(Model)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Model) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelCompute) DeepCopyInto(out *ModelCompute) {
	*out = *in
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]ComputeType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelCompute.
func (in *ModelCompute) DeepCopy() *ModelCompute {
	if in == nil {
		return nil
	}
	out := new(ModelCompute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelList) DeepCopyInto(out *ModelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Model, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelList.
func (in *ModelList) DeepCopy() *ModelList {
	if in == nil {
		return nil
	}
	out := new(ModelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelServer) DeepCopyInto(out *ModelServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelServer.
func (in *ModelServer) DeepCopy() *ModelServer {
	if in == nil {
		return nil
	}
	out := new(ModelServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelServerList) DeepCopyInto(out *ModelServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ModelServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelServerList.
func (in *ModelServerList) DeepCopy() *ModelServerList {
	if in == nil {
		return nil
	}
	out := new(ModelServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelServerSpec) DeepCopyInto(out *ModelServerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelServerSpec.
func (in *ModelServerSpec) DeepCopy() *ModelServerSpec {
	if in == nil {
		return nil
	}
	out := new(ModelServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelServerStatus) DeepCopyInto(out *ModelServerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelServerStatus.
func (in *ModelServerStatus) DeepCopy() *ModelServerStatus {
	if in == nil {
		return nil
	}
	out := new(ModelServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelSize) DeepCopyInto(out *ModelSize) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelSize.
func (in *ModelSize) DeepCopy() *ModelSize {
	if in == nil {
		return nil
	}
	out := new(ModelSize)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelSource) DeepCopyInto(out *ModelSource) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitSource)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelSource.
func (in *ModelSource) DeepCopy() *ModelSource {
	if in == nil {
		return nil
	}
	out := new(ModelSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelSpec) DeepCopyInto(out *ModelSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	if in.Training != nil {
		in, out := &in.Training, &out.Training
		*out = new(Training)
		**out = **in
	}
	out.Size = in.Size
	in.Compute.DeepCopyInto(&out.Compute)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelSpec.
func (in *ModelSpec) DeepCopy() *ModelSpec {
	if in == nil {
		return nil
	}
	out := new(ModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelStatus) DeepCopyInto(out *ModelStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Servers != nil {
		in, out := &in.Servers, &out.Servers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelStatus.
func (in *ModelStatus) DeepCopy() *ModelStatus {
	if in == nil {
		return nil
	}
	out := new(ModelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Notebook) DeepCopyInto(out *Notebook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Notebook.
func (in *Notebook) DeepCopy() *Notebook {
	if in == nil {
		return nil
	}
	out := new(Notebook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Notebook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookList) DeepCopyInto(out *NotebookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Notebook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookList.
func (in *NotebookList) DeepCopy() *NotebookList {
	if in == nil {
		return nil
	}
	out := new(NotebookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NotebookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookSpec) DeepCopyInto(out *NotebookSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookSpec.
func (in *NotebookSpec) DeepCopy() *NotebookSpec {
	if in == nil {
		return nil
	}
	out := new(NotebookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookStatus) DeepCopyInto(out *NotebookStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookStatus.
func (in *NotebookStatus) DeepCopy() *NotebookStatus {
	if in == nil {
		return nil
	}
	out := new(NotebookStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Training) DeepCopyInto(out *Training) {
	*out = *in
	out.Params = in.Params
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Training.
func (in *Training) DeepCopy() *Training {
	if in == nil {
		return nil
	}
	out := new(Training)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainingParams) DeepCopyInto(out *TrainingParams) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainingParams.
func (in *TrainingParams) DeepCopy() *TrainingParams {
	if in == nil {
		return nil
	}
	out := new(TrainingParams)
	in.DeepCopyInto(out)
	return out
}
