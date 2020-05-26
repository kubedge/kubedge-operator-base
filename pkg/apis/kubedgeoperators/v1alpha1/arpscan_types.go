package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// ArpscanSpec defines the desired state of Arpscan
type ArpscanSpec struct {
	KubedgeSpec `json:",inline"`

	// Scanners describes the set of arp scanners deployed in the kubedge cluster
	Scanners *KubedgeSetSpec `json:"scanners,omitempty"`
}

// ArpscanStatus defines the observed state of Arpscan
type ArpscanStatus struct {
	KubedgeStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Arpscan is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=arpscans,shortName=arp
// +kubebuilder:printcolumn:name="Satisfied",type="boolean",JSONPath=".status.satisfied",description="Satisfied"
type Arpscan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArpscanSpec   `json:"spec,omitempty"`
	Status ArpscanStatus `json:"status,omitempty"`
}

// Init is used to initialize an Arpscan. Namely, if the state has not been
// specified, it will be set
func (obj *Arpscan) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialized
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateUninitialized
	}
	obj.Status.Satisfied = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *Arpscan) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed Arpscan
func ToArpscan(u *unstructured.Unstructured) *Arpscan {
	var obj *Arpscan
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &Arpscan{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed Arpscan into an unstructured.Unstructured
func (obj *Arpscan) FromArpscan() *unstructured.Unstructured {
	u := NewArpscanVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the chart group has been deleted
func (obj *Arpscan) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsSatisfied returns true if the chart's actual state meets its target state
func (obj *Arpscan) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

func (obj *Arpscan) GetName() string {
	return obj.ObjectMeta.Name
}

// Returns a GKV for Arpscan
func NewArpscanVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("kubedgeoperators.kubedge.cloud/v1alpha1")
	u.SetKind("Arpscan")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArpscanList contains a list of Arpscan
type ArpscanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Arpscan `json:"items"`
}

// Convert an unstructured.Unstructured into a typed ArpscanList
func ToArpscanList(u *unstructured.Unstructured) *ArpscanList {
	var obj *ArpscanList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &ArpscanList{}
	}
	return obj
}

// Convert a typed ArpscanList into an unstructured.Unstructured
func (obj *ArpscanList) FromArpscanList() *unstructured.Unstructured {
	u := NewArpscanListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *ArpscanList) Equivalent(other *ArpscanList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for ArpscanList
func NewArpscanListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("kubedgeoperators.kubedge.cloud/v1alpha1")
	u.SetKind("ArpscanList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&Arpscan{}, &ArpscanList{})
}
