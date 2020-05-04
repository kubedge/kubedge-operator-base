package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// EMBBSliceSpec defines the desired state of EMBBSlice
type EMBBSliceSpec struct {
	KubedgeSpec `json:",inline"`

	// UPFs describes the set of UPF deployed in the simulator
	UPFs *KubedgeSetSpec `json:"upfs,omitempty"`

	SMFs *KubedgeSetSpec `json:"smfs,omitempty"`
}

// EMBBSliceStatus defines the observed state of EMBBSlice
type EMBBSliceStatus struct {
	KubedgeStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EMBBSlice is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=embbslices,shortName=embb
// +kubebuilder:printcolumn:name="Succeeded",type="boolean",JSONPath=".status.succeeded",description="Succeeded"
type EMBBSlice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EMBBSliceSpec   `json:"spec,omitempty"`
	Status EMBBSliceStatus `json:"status,omitempty"`
}

// Init is used to initialize an EMBBSlice. Namely, if the state has not been
// specified, it will be set
func (obj *EMBBSlice) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialied
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateUninitialied
	}
	obj.Status.Succeeded = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *EMBBSlice) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed EMBBSlice
func ToEMBBSlice(u *unstructured.Unstructured) *EMBBSlice {
	var obj *EMBBSlice
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &EMBBSlice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed EMBBSlice into an unstructured.Unstructured
func (obj *EMBBSlice) FromEMBBSlice() *unstructured.Unstructured {
	u := NewEMBBSliceVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the chart group has been deleted
func (obj *EMBBSlice) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsSatisfied returns true if the chart's actual state meets its target state
func (obj *EMBBSlice) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

func (obj *EMBBSlice) GetName() string {
	return obj.ObjectMeta.Name
}

// Returns a GKV for EMBBSlice
func NewEMBBSliceVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("kubedgeoperators.kubedge.cloud/v1alpha1")
	u.SetKind("EMBBSlice")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EMBBSliceList contains a list of EMBBSlice
type EMBBSliceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EMBBSlice `json:"items"`
}

// Convert an unstructured.Unstructured into a typed EMBBSliceList
func ToEMBBSliceList(u *unstructured.Unstructured) *EMBBSliceList {
	var obj *EMBBSliceList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &EMBBSliceList{}
	}
	return obj
}

// Convert a typed EMBBSliceList into an unstructured.Unstructured
func (obj *EMBBSliceList) FromEMBBSliceList() *unstructured.Unstructured {
	u := NewEMBBSliceListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *EMBBSliceList) Equivalent(other *EMBBSliceList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for EMBBSliceList
func NewEMBBSliceListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("kubedgeoperators.kubedge.cloud/v1alpha1")
	u.SetKind("EMBBSliceList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&EMBBSlice{}, &EMBBSliceList{})
}
