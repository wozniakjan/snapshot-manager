// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hooks) DeepCopyInto(out *Hooks) {
	*out = *in
	if in.Init != nil {
		in, out := &in.Init, &out.Init
		*out = new(v1.JobSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PreValidation != nil {
		in, out := &in.PreValidation, &out.PreValidation
		*out = new(v1.JobSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Validation != nil {
		in, out := &in.Validation, &out.Validation
		*out = new(v1.JobSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hooks.
func (in *Hooks) DeepCopy() *Hooks {
	if in == nil {
		return nil
	}
	out := new(Hooks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kustomization) DeepCopyInto(out *Kustomization) {
	*out = *in
	if in.CommonLabels != nil {
		in, out := &in.CommonLabels, &out.CommonLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CommonAnnotations != nil {
		in, out := &in.CommonAnnotations, &out.CommonAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Patches != nil {
		in, out := &in.Patches, &out.Patches
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kustomization.
func (in *Kustomization) DeepCopy() *Kustomization {
	if in == nil {
		return nil
	}
	out := new(Kustomization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceName) DeepCopyInto(out *ResourceName) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceName.
func (in *ResourceName) DeepCopy() *ResourceName {
	if in == nil {
		return nil
	}
	out := new(ResourceName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetType) DeepCopyInto(out *StatefulSetType) {
	*out = *in
	if in.SnapshotClaimStorageClass != nil {
		in, out := &in.SnapshotClaimStorageClass, &out.SnapshotClaimStorageClass
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetType.
func (in *StatefulSetType) DeepCopy() *StatefulSetType {
	if in == nil {
		return nil
	}
	out := new(StatefulSetType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationRun) DeepCopyInto(out *ValidationRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationRun.
func (in *ValidationRun) DeepCopy() *ValidationRun {
	if in == nil {
		return nil
	}
	out := new(ValidationRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ValidationRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationRunList) DeepCopyInto(out *ValidationRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ValidationRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationRunList.
func (in *ValidationRunList) DeepCopy() *ValidationRunList {
	if in == nil {
		return nil
	}
	out := new(ValidationRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ValidationRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationRunObjects) DeepCopyInto(out *ValidationRunObjects) {
	*out = *in
	if in.Claims != nil {
		in, out := &in.Claims, &out.Claims
		*out = make([]corev1.PersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kustomized != nil {
		in, out := &in.Kustomized, &out.Kustomized
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationRunObjects.
func (in *ValidationRunObjects) DeepCopy() *ValidationRunObjects {
	if in == nil {
		return nil
	}
	out := new(ValidationRunObjects)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationRunSpec) DeepCopyInto(out *ValidationRunSpec) {
	*out = *in
	if in.ClaimsToSnapshots != nil {
		in, out := &in.ClaimsToSnapshots, &out.ClaimsToSnapshots
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Objects.DeepCopyInto(&out.Objects)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationRunSpec.
func (in *ValidationRunSpec) DeepCopy() *ValidationRunSpec {
	if in == nil {
		return nil
	}
	out := new(ValidationRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationRunStatus) DeepCopyInto(out *ValidationRunStatus) {
	*out = *in
	if in.KustStarted != nil {
		in, out := &in.KustStarted, &out.KustStarted
		*out = (*in).DeepCopy()
	}
	if in.KustFinished != nil {
		in, out := &in.KustFinished, &out.KustFinished
		*out = (*in).DeepCopy()
	}
	if in.InitStarted != nil {
		in, out := &in.InitStarted, &out.InitStarted
		*out = (*in).DeepCopy()
	}
	if in.InitFinished != nil {
		in, out := &in.InitFinished, &out.InitFinished
		*out = (*in).DeepCopy()
	}
	if in.PreValidationStarted != nil {
		in, out := &in.PreValidationStarted, &out.PreValidationStarted
		*out = (*in).DeepCopy()
	}
	if in.PreValidationFinished != nil {
		in, out := &in.PreValidationFinished, &out.PreValidationFinished
		*out = (*in).DeepCopy()
	}
	if in.ValidationStarted != nil {
		in, out := &in.ValidationStarted, &out.ValidationStarted
		*out = (*in).DeepCopy()
	}
	if in.ValidationFinished != nil {
		in, out := &in.ValidationFinished, &out.ValidationFinished
		*out = (*in).DeepCopy()
	}
	if in.CleanupStarted != nil {
		in, out := &in.CleanupStarted, &out.CleanupStarted
		*out = (*in).DeepCopy()
	}
	if in.CleanupFinished != nil {
		in, out := &in.CleanupFinished, &out.CleanupFinished
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationRunStatus.
func (in *ValidationRunStatus) DeepCopy() *ValidationRunStatus {
	if in == nil {
		return nil
	}
	out := new(ValidationRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationStrategy) DeepCopyInto(out *ValidationStrategy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationStrategy.
func (in *ValidationStrategy) DeepCopy() *ValidationStrategy {
	if in == nil {
		return nil
	}
	out := new(ValidationStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ValidationStrategy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationStrategyList) DeepCopyInto(out *ValidationStrategyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ValidationStrategy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationStrategyList.
func (in *ValidationStrategyList) DeepCopy() *ValidationStrategyList {
	if in == nil {
		return nil
	}
	out := new(ValidationStrategyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ValidationStrategyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationStrategySpec) DeepCopyInto(out *ValidationStrategySpec) {
	*out = *in
	if in.StsType != nil {
		in, out := &in.StsType, &out.StsType
		*out = new(StatefulSetType)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalResources != nil {
		in, out := &in.AdditionalResources, &out.AdditionalResources
		*out = make([]ResourceName, len(*in))
		copy(*out, *in)
	}
	in.Kustomization.DeepCopyInto(&out.Kustomization)
	if in.Hooks != nil {
		in, out := &in.Hooks, &out.Hooks
		*out = new(Hooks)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationStrategySpec.
func (in *ValidationStrategySpec) DeepCopy() *ValidationStrategySpec {
	if in == nil {
		return nil
	}
	out := new(ValidationStrategySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationStrategyStatus) DeepCopyInto(out *ValidationStrategyStatus) {
	*out = *in
	if in.finishedRuns != nil {
		in, out := &in.finishedRuns, &out.finishedRuns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.activeRuns != nil {
		in, out := &in.activeRuns, &out.activeRuns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationStrategyStatus.
func (in *ValidationStrategyStatus) DeepCopy() *ValidationStrategyStatus {
	if in == nil {
		return nil
	}
	out := new(ValidationStrategyStatus)
	in.DeepCopyInto(out)
	return out
}
