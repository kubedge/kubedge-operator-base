// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Arpscan) DeepCopyInto(out *Arpscan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Arpscan.
func (in *Arpscan) DeepCopy() *Arpscan {
	if in == nil {
		return nil
	}
	out := new(Arpscan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Arpscan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArpscanList) DeepCopyInto(out *ArpscanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Arpscan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArpscanList.
func (in *ArpscanList) DeepCopy() *ArpscanList {
	if in == nil {
		return nil
	}
	out := new(ArpscanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArpscanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArpscanSpec) DeepCopyInto(out *ArpscanSpec) {
	*out = *in
	in.KubedgeSpec.DeepCopyInto(&out.KubedgeSpec)
	if in.Scanners != nil {
		in, out := &in.Scanners, &out.Scanners
		*out = new(KubedgeSetSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArpscanSpec.
func (in *ArpscanSpec) DeepCopy() *ArpscanSpec {
	if in == nil {
		return nil
	}
	out := new(ArpscanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArpscanStatus) DeepCopyInto(out *ArpscanStatus) {
	*out = *in
	in.KubedgeStatus.DeepCopyInto(&out.KubedgeStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArpscanStatus.
func (in *ArpscanStatus) DeepCopy() *ArpscanStatus {
	if in == nil {
		return nil
	}
	out := new(ArpscanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRevision) DeepCopyInto(out *ControllerRevision) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Data.DeepCopyInto(&out.Data)
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
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerRevision, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
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
func (in *ECDSCluster) DeepCopyInto(out *ECDSCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECDSCluster.
func (in *ECDSCluster) DeepCopy() *ECDSCluster {
	if in == nil {
		return nil
	}
	out := new(ECDSCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ECDSCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECDSClusterList) DeepCopyInto(out *ECDSClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ECDSCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECDSClusterList.
func (in *ECDSClusterList) DeepCopy() *ECDSClusterList {
	if in == nil {
		return nil
	}
	out := new(ECDSClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ECDSClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECDSClusterSpec) DeepCopyInto(out *ECDSClusterSpec) {
	*out = *in
	in.KubedgeSpec.DeepCopyInto(&out.KubedgeSpec)
	if in.Platforms != nil {
		in, out := &in.Platforms, &out.Platforms
		*out = new(KubedgeSetSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.FrontEnds != nil {
		in, out := &in.FrontEnds, &out.FrontEnds
		*out = new(KubedgeSetSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Enrichments != nil {
		in, out := &in.Enrichments, &out.Enrichments
		*out = new(KubedgeSetSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.BusinessLogics != nil {
		in, out := &in.BusinessLogics, &out.BusinessLogics
		*out = new(KubedgeSetSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LoadBalancers != nil {
		in, out := &in.LoadBalancers, &out.LoadBalancers
		*out = new(KubedgeSetSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECDSClusterSpec.
func (in *ECDSClusterSpec) DeepCopy() *ECDSClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ECDSClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECDSClusterStatus) DeepCopyInto(out *ECDSClusterStatus) {
	*out = *in
	in.KubedgeStatus.DeepCopyInto(&out.KubedgeStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECDSClusterStatus.
func (in *ECDSClusterStatus) DeepCopy() *ECDSClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ECDSClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EMBBSlice) DeepCopyInto(out *EMBBSlice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EMBBSlice.
func (in *EMBBSlice) DeepCopy() *EMBBSlice {
	if in == nil {
		return nil
	}
	out := new(EMBBSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EMBBSlice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EMBBSliceList) DeepCopyInto(out *EMBBSliceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EMBBSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EMBBSliceList.
func (in *EMBBSliceList) DeepCopy() *EMBBSliceList {
	if in == nil {
		return nil
	}
	out := new(EMBBSliceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EMBBSliceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EMBBSliceSpec) DeepCopyInto(out *EMBBSliceSpec) {
	*out = *in
	in.KubedgeSpec.DeepCopyInto(&out.KubedgeSpec)
	if in.UPFs != nil {
		in, out := &in.UPFs, &out.UPFs
		*out = new(KubedgeSetSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.SMFs != nil {
		in, out := &in.SMFs, &out.SMFs
		*out = new(KubedgeSetSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EMBBSliceSpec.
func (in *EMBBSliceSpec) DeepCopy() *EMBBSliceSpec {
	if in == nil {
		return nil
	}
	out := new(EMBBSliceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EMBBSliceStatus) DeepCopyInto(out *EMBBSliceStatus) {
	*out = *in
	in.KubedgeStatus.DeepCopyInto(&out.KubedgeStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EMBBSliceStatus.
func (in *EMBBSliceStatus) DeepCopy() *EMBBSliceStatus {
	if in == nil {
		return nil
	}
	out := new(EMBBSliceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubedgeCondition) DeepCopyInto(out *KubedgeCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubedgeCondition.
func (in *KubedgeCondition) DeepCopy() *KubedgeCondition {
	if in == nil {
		return nil
	}
	out := new(KubedgeCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubedgeConditionListHelper) DeepCopyInto(out *KubedgeConditionListHelper) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubedgeCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubedgeConditionListHelper.
func (in *KubedgeConditionListHelper) DeepCopy() *KubedgeConditionListHelper {
	if in == nil {
		return nil
	}
	out := new(KubedgeConditionListHelper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubedgeSetSpec) DeepCopyInto(out *KubedgeSetSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubedgeSetSpec.
func (in *KubedgeSetSpec) DeepCopy() *KubedgeSetSpec {
	if in == nil {
		return nil
	}
	out := new(KubedgeSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubedgeSource) DeepCopyInto(out *KubedgeSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubedgeSource.
func (in *KubedgeSource) DeepCopy() *KubedgeSource {
	if in == nil {
		return nil
	}
	out := new(KubedgeSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubedgeSpec) DeepCopyInto(out *KubedgeSpec) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(KubedgeSource)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubedgeSpec.
func (in *KubedgeSpec) DeepCopy() *KubedgeSpec {
	if in == nil {
		return nil
	}
	out := new(KubedgeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubedgeStatus) DeepCopyInto(out *KubedgeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]KubedgeCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubedgeStatus.
func (in *KubedgeStatus) DeepCopy() *KubedgeStatus {
	if in == nil {
		return nil
	}
	out := new(KubedgeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MMESim) DeepCopyInto(out *MMESim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MMESim.
func (in *MMESim) DeepCopy() *MMESim {
	if in == nil {
		return nil
	}
	out := new(MMESim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MMESim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MMESimList) DeepCopyInto(out *MMESimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MMESim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MMESimList.
func (in *MMESimList) DeepCopy() *MMESimList {
	if in == nil {
		return nil
	}
	out := new(MMESimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MMESimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MMESimSpec) DeepCopyInto(out *MMESimSpec) {
	*out = *in
	in.KubedgeSpec.DeepCopyInto(&out.KubedgeSpec)
	if in.LCs != nil {
		in, out := &in.LCs, &out.LCs
		*out = new(KubedgeSetSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.GPBs != nil {
		in, out := &in.GPBs, &out.GPBs
		*out = new(KubedgeSetSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NCBs != nil {
		in, out := &in.NCBs, &out.NCBs
		*out = new(KubedgeSetSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.FSBs != nil {
		in, out := &in.FSBs, &out.FSBs
		*out = new(KubedgeSetSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MMESimSpec.
func (in *MMESimSpec) DeepCopy() *MMESimSpec {
	if in == nil {
		return nil
	}
	out := new(MMESimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MMESimStatus) DeepCopyInto(out *MMESimStatus) {
	*out = *in
	in.KubedgeStatus.DeepCopyInto(&out.KubedgeStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MMESimStatus.
func (in *MMESimStatus) DeepCopy() *MMESimStatus {
	if in == nil {
		return nil
	}
	out := new(MMESimStatus)
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
