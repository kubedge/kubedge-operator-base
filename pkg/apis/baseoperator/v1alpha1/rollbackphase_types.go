package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// RollbackPhaseSpec defines the desired state of RollbackPhase
type RollbackPhaseSpec struct {
	PhaseSpec `json:",inline"`
}

// RollbackPhaseStatus defines the observed state of RollbackPhase
type RollbackPhaseStatus struct {
	PhaseStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RollbackPhase is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=rollbackphases,shortName=osrbck
// +kubebuilder:printcolumn:name="Succeeded",type="boolean",JSONPath=".status.succeeded",description="Succeeded"
type RollbackPhase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RollbackPhaseSpec   `json:"spec,omitempty"`
	Status RollbackPhaseStatus `json:"status,omitempty"`
}

// Init is used to initialize an RollbackPhase. Namely, if the state has not been
// specified, it will be set
func (obj *RollbackPhase) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialied
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateUninitialied
	}
	obj.Status.Succeeded = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *RollbackPhase) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed RollbackPhase
func ToRollbackPhase(u *unstructured.Unstructured) *RollbackPhase {
	var obj *RollbackPhase
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &RollbackPhase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed RollbackPhase into an unstructured.Unstructured
func (obj *RollbackPhase) FromRollbackPhase() *unstructured.Unstructured {
	u := NewRollbackPhaseVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the chart group has been deleted
func (obj *RollbackPhase) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsSatisfied returns true if the chart's actual state meets its target state
func (obj *RollbackPhase) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

func (obj *RollbackPhase) GetName() string {
	return obj.ObjectMeta.Name
}

// Returns a GKV for RollbackPhase
func NewRollbackPhaseVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("baseoperator.kubedge.cloud/v1alpha1")
	u.SetKind("RollbackPhase")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RollbackPhaseList contains a list of RollbackPhase
type RollbackPhaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RollbackPhase `json:"items"`
}

// Convert an unstructured.Unstructured into a typed RollbackPhaseList
func ToRollbackPhaseList(u *unstructured.Unstructured) *RollbackPhaseList {
	var obj *RollbackPhaseList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &RollbackPhaseList{}
	}
	return obj
}

// Convert a typed RollbackPhaseList into an unstructured.Unstructured
func (obj *RollbackPhaseList) FromRollbackPhaseList() *unstructured.Unstructured {
	u := NewRollbackPhaseListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *RollbackPhaseList) Equivalent(other *RollbackPhaseList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for RollbackPhaseList
func NewRollbackPhaseListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("baseoperator.kubedge.cloud/v1alpha1")
	u.SetKind("RollbackPhaseList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&RollbackPhase{}, &RollbackPhaseList{})
}
