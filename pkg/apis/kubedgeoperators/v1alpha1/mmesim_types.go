package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// MMESimSpec defines the desired state of MMESim
type MMESimSpec struct {
	KubedgeSpec `json:",inline"`

	// LCs describes the set of LC deployed in the simulator
	LCs *KubedgeSetSpec `json:"lcs,omitempty"`

	// GPBs describes the set of GPB deployed in the simulator
	GPBs *KubedgeSetSpec `json:"gpbs,omitempty"`

	// NCBs describes the set of NCB deployed in the simulator
	NCBs *KubedgeSetSpec `json:"ncbs,omitempty"`

	// FSBs describes the set of FSB deployed in the simulator
	FSBs *KubedgeSetSpec `json:"fsbs,omitempty"`
}

// MMESimStatus defines the observed state of MMESim
type MMESimStatus struct {
	KubedgeStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MMESim is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=mmesims,shortName=mme
// +kubebuilder:printcolumn:name="Satisfied",type="boolean",JSONPath=".status.satisfied",description="Satisfied"
type MMESim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MMESimSpec   `json:"spec,omitempty"`
	Status MMESimStatus `json:"status,omitempty"`
}

// Init is used to initialize an MMESim. Namely, if the state has not been
// specified, it will be set
func (obj *MMESim) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialized
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateUninitialized
	}
	obj.Status.Satisfied = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *MMESim) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed MMESim
func ToMMESim(u *unstructured.Unstructured) *MMESim {
	var obj *MMESim
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &MMESim{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed MMESim into an unstructured.Unstructured
func (obj *MMESim) FromMMESim() *unstructured.Unstructured {
	u := NewMMESimVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the chart group has been deleted
func (obj *MMESim) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsSatisfied returns true if the chart's actual state meets its target state
func (obj *MMESim) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

func (obj *MMESim) GetName() string {
	return obj.ObjectMeta.Name
}

// Returns a GKV for MMESim
func NewMMESimVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("kubedgeoperators.kubedge.cloud/v1alpha1")
	u.SetKind("MMESim")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MMESimList contains a list of MMESim
type MMESimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MMESim `json:"items"`
}

// Convert an unstructured.Unstructured into a typed MMESimList
func ToMMESimList(u *unstructured.Unstructured) *MMESimList {
	var obj *MMESimList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &MMESimList{}
	}
	return obj
}

// Convert a typed MMESimList into an unstructured.Unstructured
func (obj *MMESimList) FromMMESimList() *unstructured.Unstructured {
	u := NewMMESimListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *MMESimList) Equivalent(other *MMESimList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for MMESimList
func NewMMESimListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("kubedgeoperators.kubedge.cloud/v1alpha1")
	u.SetKind("MMESimList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&MMESim{}, &MMESimList{})
}
