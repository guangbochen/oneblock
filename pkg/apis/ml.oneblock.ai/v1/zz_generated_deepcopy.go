//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 1block.ai.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1

import (
	managementoneblockaiv1 "github.com/oneblock-ai/oneblock/pkg/apis/management.oneblock.ai/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dataset) DeepCopyInto(out *Dataset) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
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
	return
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
func (in *DatasetSpec) DeepCopyInto(out *DatasetSpec) {
	*out = *in
	return
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
	return
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
func (in *DeploymentConfig) DeepCopyInto(out *DeploymentConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConfig.
func (in *DeploymentConfig) DeepCopy() *DeploymentConfig {
	if in == nil {
		return nil
	}
	out := new(DeploymentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EngineConfig) DeepCopyInto(out *EngineConfig) {
	*out = *in
	in.Generation.DeepCopyInto(&out.Generation)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EngineConfig.
func (in *EngineConfig) DeepCopy() *EngineConfig {
	if in == nil {
		return nil
	}
	out := new(EngineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenerationConfig) DeepCopyInto(out *GenerationConfig) {
	*out = *in
	out.PromptFormat = in.PromptFormat
	if in.StoppingSequences != nil {
		in, out := &in.StoppingSequences, &out.StoppingSequences
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenerationConfig.
func (in *GenerationConfig) DeepCopy() *GenerationConfig {
	if in == nil {
		return nil
	}
	out := new(GenerationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HFSecretRef) DeepCopyInto(out *HFSecretRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HFSecretRef.
func (in *HFSecretRef) DeepCopy() *HFSecretRef {
	if in == nil {
		return nil
	}
	out := new(HFSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeadGroupSpec) DeepCopyInto(out *HeadGroupSpec) {
	*out = *in
	if in.RayStartParams != nil {
		in, out := &in.RayStartParams, &out.RayStartParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(Volume)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeadGroupSpec.
func (in *HeadGroupSpec) DeepCopy() *HeadGroupSpec {
	if in == nil {
		return nil
	}
	out := new(HeadGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MLClusterRef) DeepCopyInto(out *MLClusterRef) {
	*out = *in
	in.RayClusterSpec.DeepCopyInto(&out.RayClusterSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MLClusterRef.
func (in *MLClusterRef) DeepCopy() *MLClusterRef {
	if in == nil {
		return nil
	}
	out := new(MLClusterRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MLService) DeepCopyInto(out *MLService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MLService.
func (in *MLService) DeepCopy() *MLService {
	if in == nil {
		return nil
	}
	out := new(MLService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MLService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MLServiceList) DeepCopyInto(out *MLServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MLService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MLServiceList.
func (in *MLServiceList) DeepCopy() *MLServiceList {
	if in == nil {
		return nil
	}
	out := new(MLServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MLServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MLServiceSpec) DeepCopyInto(out *MLServiceSpec) {
	*out = *in
	if in.ModelTemplateVersionRef != nil {
		in, out := &in.ModelTemplateVersionRef, &out.ModelTemplateVersionRef
		*out = new(ModelTemplateVersionRef)
		**out = **in
	}
	if in.HFSecretRef != nil {
		in, out := &in.HFSecretRef, &out.HFSecretRef
		*out = new(HFSecretRef)
		**out = **in
	}
	if in.MLClusterRef != nil {
		in, out := &in.MLClusterRef, &out.MLClusterRef
		*out = new(MLClusterRef)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MLServiceSpec.
func (in *MLServiceSpec) DeepCopy() *MLServiceSpec {
	if in == nil {
		return nil
	}
	out := new(MLServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MLServiceStatus) DeepCopyInto(out *MLServiceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]managementoneblockaiv1.Condition, len(*in))
		copy(*out, *in)
	}
	in.RayServiceStatuses.DeepCopyInto(&out.RayServiceStatuses)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MLServiceStatus.
func (in *MLServiceStatus) DeepCopy() *MLServiceStatus {
	if in == nil {
		return nil
	}
	out := new(MLServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelTemplateVersion) DeepCopyInto(out *ModelTemplateVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelTemplateVersion.
func (in *ModelTemplateVersion) DeepCopy() *ModelTemplateVersion {
	if in == nil {
		return nil
	}
	out := new(ModelTemplateVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelTemplateVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelTemplateVersionList) DeepCopyInto(out *ModelTemplateVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ModelTemplateVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelTemplateVersionList.
func (in *ModelTemplateVersionList) DeepCopy() *ModelTemplateVersionList {
	if in == nil {
		return nil
	}
	out := new(ModelTemplateVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelTemplateVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelTemplateVersionRef) DeepCopyInto(out *ModelTemplateVersionRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelTemplateVersionRef.
func (in *ModelTemplateVersionRef) DeepCopy() *ModelTemplateVersionRef {
	if in == nil {
		return nil
	}
	out := new(ModelTemplateVersionRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelTemplateVersionSpec) DeepCopyInto(out *ModelTemplateVersionSpec) {
	*out = *in
	in.EngineConfig.DeepCopyInto(&out.EngineConfig)
	out.DeploymentConfig = in.DeploymentConfig
	in.ScalingConfig.DeepCopyInto(&out.ScalingConfig)
	in.WorkerGroupSpec.DeepCopyInto(&out.WorkerGroupSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelTemplateVersionSpec.
func (in *ModelTemplateVersionSpec) DeepCopy() *ModelTemplateVersionSpec {
	if in == nil {
		return nil
	}
	out := new(ModelTemplateVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelTemplateVersionStatus) DeepCopyInto(out *ModelTemplateVersionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]managementoneblockaiv1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelTemplateVersionStatus.
func (in *ModelTemplateVersionStatus) DeepCopy() *ModelTemplateVersionStatus {
	if in == nil {
		return nil
	}
	out := new(ModelTemplateVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Notebook) DeepCopyInto(out *Notebook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
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
	return
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
	in.Template.DeepCopyInto(&out.Template)
	return
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
		*out = make([]managementoneblockaiv1.Condition, len(*in))
		copy(*out, *in)
	}
	in.State.DeepCopyInto(&out.State)
	return
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
func (in *NotebookTemplateSpec) DeepCopyInto(out *NotebookTemplateSpec) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookTemplateSpec.
func (in *NotebookTemplateSpec) DeepCopy() *NotebookTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(NotebookTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromptFormat) DeepCopyInto(out *PromptFormat) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromptFormat.
func (in *PromptFormat) DeepCopy() *PromptFormat {
	if in == nil {
		return nil
	}
	out := new(PromptFormat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayClusterSpec) DeepCopyInto(out *RayClusterSpec) {
	*out = *in
	in.HeadGroupSpec.DeepCopyInto(&out.HeadGroupSpec)
	if in.WorkerGroupSpec != nil {
		in, out := &in.WorkerGroupSpec, &out.WorkerGroupSpec
		*out = make([]WorkerGroupSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayClusterSpec.
func (in *RayClusterSpec) DeepCopy() *RayClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RayClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingConfig) DeepCopyInto(out *ScalingConfig) {
	*out = *in
	if in.ResourcesPerWorker != nil {
		in, out := &in.ResourcesPerWorker, &out.ResourcesPerWorker
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingConfig.
func (in *ScalingConfig) DeepCopy() *ScalingConfig {
	if in == nil {
		return nil
	}
	out := new(ScalingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerGroupSpec) DeepCopyInto(out *WorkerGroupSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
	if in.RayStartParams != nil {
		in, out := &in.RayStartParams, &out.RayStartParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AcceleratorTypes != nil {
		in, out := &in.AcceleratorTypes, &out.AcceleratorTypes
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(Volume)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerGroupSpec.
func (in *WorkerGroupSpec) DeepCopy() *WorkerGroupSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerGroupSpec)
	in.DeepCopyInto(out)
	return out
}
