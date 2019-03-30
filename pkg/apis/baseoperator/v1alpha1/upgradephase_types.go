package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// UpgradePhaseSpec defines the desired state of UpgradePhase
type UpgradePhaseSpec struct {
	PhaseSpec `json:",inline"`
}

// UpgradePhaseStatus defines the observed state of UpgradePhase
type UpgradePhaseStatus struct {
	PhaseStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UpgradePhase is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=upgradephases,shortName=osupg
// +kubebuilder:printcolumn:name="Succeeded",type="boolean",JSONPath=".status.succeeded",description="Succeeded"
type UpgradePhase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UpgradePhaseSpec   `json:"spec,omitempty"`
	Status UpgradePhaseStatus `json:"status,omitempty"`
}

// Init is used to initialize an UpgradePhase. Namely, if the state has not been
// specified, it will be set
func (obj *UpgradePhase) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialied
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateUninitialied
	}
	obj.Status.Succeeded = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *UpgradePhase) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed UpgradePhase
func ToUpgradePhase(u *unstructured.Unstructured) *UpgradePhase {
	var obj *UpgradePhase
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &UpgradePhase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed UpgradePhase into an unstructured.Unstructured
func (obj *UpgradePhase) FromUpgradePhase() *unstructured.Unstructured {
	u := NewUpgradePhaseVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the chart group has been deleted
func (obj *UpgradePhase) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsSatisfied returns true if the chart's actual state meets its target state
func (obj *UpgradePhase) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

func (obj *UpgradePhase) GetName() string {
	return obj.ObjectMeta.Name
}

// Returns a GKV for UpgradePhase
func NewUpgradePhaseVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("baseoperator.kubedge.cloud/v1alpha1")
	u.SetKind("UpgradePhase")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UpgradePhaseList contains a list of UpgradePhase
type UpgradePhaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UpgradePhase `json:"items"`
}

// Convert an unstructured.Unstructured into a typed UpgradePhaseList
func ToUpgradePhaseList(u *unstructured.Unstructured) *UpgradePhaseList {
	var obj *UpgradePhaseList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &UpgradePhaseList{}
	}
	return obj
}

// Convert a typed UpgradePhaseList into an unstructured.Unstructured
func (obj *UpgradePhaseList) FromUpgradePhaseList() *unstructured.Unstructured {
	u := NewUpgradePhaseListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *UpgradePhaseList) Equivalent(other *UpgradePhaseList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for UpgradePhaseList
func NewUpgradePhaseListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("baseoperator.kubedge.cloud/v1alpha1")
	u.SetKind("UpgradePhaseList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&UpgradePhase{}, &UpgradePhaseList{})
}
