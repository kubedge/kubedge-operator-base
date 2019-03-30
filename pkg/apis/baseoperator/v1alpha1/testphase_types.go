package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// TestPhaseSpec defines the desired state of TestPhase
type TestPhaseSpec struct {
	PhaseSpec `json:",inline"`
}

// TestPhaseStatus defines the observed state of TestPhase
type TestPhaseStatus struct {
	PhaseStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TestPhase is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=testphases,shortName=ostest
// +kubebuilder:printcolumn:name="Succeeded",type="boolean",JSONPath=".status.succeeded",description="Succeeded"
type TestPhase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestPhaseSpec   `json:"spec,omitempty"`
	Status TestPhaseStatus `json:"status,omitempty"`
}

// Init is used to initialize an TestPhase. Namely, if the state has not been
// specified, it will be set
func (obj *TestPhase) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialied
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateUninitialied
	}
	obj.Status.Succeeded = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *TestPhase) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed TestPhase
func ToTestPhase(u *unstructured.Unstructured) *TestPhase {
	var obj *TestPhase
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &TestPhase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed TestPhase into an unstructured.Unstructured
func (obj *TestPhase) FromTestPhase() *unstructured.Unstructured {
	u := NewTestPhaseVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the chart group has been deleted
func (obj *TestPhase) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsSatisfied returns true if the chart's actual state meets its target state
func (obj *TestPhase) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

func (obj *TestPhase) GetName() string {
	return obj.ObjectMeta.Name
}

// Returns a GKV for TestPhase
func NewTestPhaseVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("baseoperator.kubedge.cloud/v1alpha1")
	u.SetKind("TestPhase")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TestPhaseList contains a list of TestPhase
type TestPhaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestPhase `json:"items"`
}

// Convert an unstructured.Unstructured into a typed TestPhaseList
func ToTestPhaseList(u *unstructured.Unstructured) *TestPhaseList {
	var obj *TestPhaseList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &TestPhaseList{}
	}
	return obj
}

// Convert a typed TestPhaseList into an unstructured.Unstructured
func (obj *TestPhaseList) FromTestPhaseList() *unstructured.Unstructured {
	u := NewTestPhaseListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *TestPhaseList) Equivalent(other *TestPhaseList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for TestPhaseList
func NewTestPhaseListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("baseoperator.kubedge.cloud/v1alpha1")
	u.SetKind("TestPhaseList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&TestPhase{}, &TestPhaseList{})
}
