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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRevision) DeepCopyInto(out *ControllerRevision) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Data.DeepCopyInto(&out.Data)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRevision.
func (in *ControllerRevision) DeepCopy() *ControllerRevision {
	if in == nil {
		return nil
	}
	out := new(ControllerRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerRevision) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRevisionList) DeepCopyInto(out *ControllerRevisionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerRevision, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRevisionList.
func (in *ControllerRevisionList) DeepCopy() *ControllerRevisionList {
	if in == nil {
		return nil
	}
	out := new(ControllerRevisionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerRevisionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LcmResourceCondition) DeepCopyInto(out *LcmResourceCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LcmResourceCondition.
func (in *LcmResourceCondition) DeepCopy() *LcmResourceCondition {
	if in == nil {
		return nil
	}
	out := new(LcmResourceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LcmResourceConditionListHelper) DeepCopyInto(out *LcmResourceConditionListHelper) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LcmResourceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LcmResourceConditionListHelper.
func (in *LcmResourceConditionListHelper) DeepCopy() *LcmResourceConditionListHelper {
	if in == nil {
		return nil
	}
	out := new(LcmResourceConditionListHelper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackLcmStatus) DeepCopyInto(out *OpenstackLcmStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]LcmResourceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackLcmStatus.
func (in *OpenstackLcmStatus) DeepCopy() *OpenstackLcmStatus {
	if in == nil {
		return nil
	}
	out := new(OpenstackLcmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhaseList) DeepCopyInto(out *PhaseList) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]unstructured.Unstructured, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhaseList.
func (in *PhaseList) DeepCopy() *PhaseList {
	if in == nil {
		return nil
	}
	out := new(PhaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhaseSource) DeepCopyInto(out *PhaseSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhaseSource.
func (in *PhaseSource) DeepCopy() *PhaseSource {
	if in == nil {
		return nil
	}
	out := new(PhaseSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhaseSpec) DeepCopyInto(out *PhaseSpec) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(PhaseSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhaseSpec.
func (in *PhaseSpec) DeepCopy() *PhaseSpec {
	if in == nil {
		return nil
	}
	out := new(PhaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhaseStatus) DeepCopyInto(out *PhaseStatus) {
	*out = *in
	in.OpenstackLcmStatus.DeepCopyInto(&out.OpenstackLcmStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhaseStatus.
func (in *PhaseStatus) DeepCopy() *PhaseStatus {
	if in == nil {
		return nil
	}
	out := new(PhaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollbackPhase) DeepCopyInto(out *RollbackPhase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollbackPhase.
func (in *RollbackPhase) DeepCopy() *RollbackPhase {
	if in == nil {
		return nil
	}
	out := new(RollbackPhase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RollbackPhase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollbackPhaseList) DeepCopyInto(out *RollbackPhaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RollbackPhase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollbackPhaseList.
func (in *RollbackPhaseList) DeepCopy() *RollbackPhaseList {
	if in == nil {
		return nil
	}
	out := new(RollbackPhaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RollbackPhaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollbackPhaseSpec) DeepCopyInto(out *RollbackPhaseSpec) {
	*out = *in
	in.PhaseSpec.DeepCopyInto(&out.PhaseSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollbackPhaseSpec.
func (in *RollbackPhaseSpec) DeepCopy() *RollbackPhaseSpec {
	if in == nil {
		return nil
	}
	out := new(RollbackPhaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollbackPhaseStatus) DeepCopyInto(out *RollbackPhaseStatus) {
	*out = *in
	in.PhaseStatus.DeepCopyInto(&out.PhaseStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollbackPhaseStatus.
func (in *RollbackPhaseStatus) DeepCopy() *RollbackPhaseStatus {
	if in == nil {
		return nil
	}
	out := new(RollbackPhaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubResourceList) DeepCopyInto(out *SubResourceList) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]unstructured.Unstructured, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubResourceList.
func (in *SubResourceList) DeepCopy() *SubResourceList {
	if in == nil {
		return nil
	}
	out := new(SubResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestPhase) DeepCopyInto(out *TestPhase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestPhase.
func (in *TestPhase) DeepCopy() *TestPhase {
	if in == nil {
		return nil
	}
	out := new(TestPhase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestPhase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestPhaseList) DeepCopyInto(out *TestPhaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TestPhase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestPhaseList.
func (in *TestPhaseList) DeepCopy() *TestPhaseList {
	if in == nil {
		return nil
	}
	out := new(TestPhaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestPhaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestPhaseSpec) DeepCopyInto(out *TestPhaseSpec) {
	*out = *in
	in.PhaseSpec.DeepCopyInto(&out.PhaseSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestPhaseSpec.
func (in *TestPhaseSpec) DeepCopy() *TestPhaseSpec {
	if in == nil {
		return nil
	}
	out := new(TestPhaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestPhaseStatus) DeepCopyInto(out *TestPhaseStatus) {
	*out = *in
	in.PhaseStatus.DeepCopyInto(&out.PhaseStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestPhaseStatus.
func (in *TestPhaseStatus) DeepCopy() *TestPhaseStatus {
	if in == nil {
		return nil
	}
	out := new(TestPhaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradePhase) DeepCopyInto(out *UpgradePhase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradePhase.
func (in *UpgradePhase) DeepCopy() *UpgradePhase {
	if in == nil {
		return nil
	}
	out := new(UpgradePhase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpgradePhase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradePhaseList) DeepCopyInto(out *UpgradePhaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UpgradePhase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradePhaseList.
func (in *UpgradePhaseList) DeepCopy() *UpgradePhaseList {
	if in == nil {
		return nil
	}
	out := new(UpgradePhaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpgradePhaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradePhaseSpec) DeepCopyInto(out *UpgradePhaseSpec) {
	*out = *in
	in.PhaseSpec.DeepCopyInto(&out.PhaseSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradePhaseSpec.
func (in *UpgradePhaseSpec) DeepCopy() *UpgradePhaseSpec {
	if in == nil {
		return nil
	}
	out := new(UpgradePhaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradePhaseStatus) DeepCopyInto(out *UpgradePhaseStatus) {
	*out = *in
	in.PhaseStatus.DeepCopyInto(&out.PhaseStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradePhaseStatus.
func (in *UpgradePhaseStatus) DeepCopy() *UpgradePhaseStatus {
	if in == nil {
		return nil
	}
	out := new(UpgradePhaseStatus)
	in.DeepCopyInto(out)
	return out
}
