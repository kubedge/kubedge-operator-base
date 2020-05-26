package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// ECDSClusterSpec defines the desired state of ECDSCluster
type ECDSClusterSpec struct {
	KubedgeSpec `json:",inline"`

	// Platforms describes the set of Platform execution context deployed in the cluster
	Platforms *KubedgeSetSpec `json:"platforms,omitempty"`

	// FrontEnds describes the set of FrontEnd execution context deployed in the cluster
	FrontEnds *KubedgeSetSpec `json:"frontEnds,omitempty"`

	// Enrichments describes the set of Enrichment execution context deployed in the cluster
	Enrichments *KubedgeSetSpec `json:"enrichments,omitempty"`

	// BusinessLogics describes the set of BusinessLogic execution context deployed in the cluster
	BusinessLogics *KubedgeSetSpec `json:"businessLogics,omitempty"`

	// LoadBalancers describes the set of LoadBalancer deployed in the cluster
	LoadBalancers *KubedgeSetSpec `json:"loadbalancers,omitempty"`
}

// ECDSClusterStatus defines the observed state of ECDSCluster
type ECDSClusterStatus struct {
	KubedgeStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ECDSCluster is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=ecdsclusters,shortName=ecds
// +kubebuilder:printcolumn:name="Satisfied",type="boolean",JSONPath=".status.satisfied",description="Satisfied"
type ECDSCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ECDSClusterSpec   `json:"spec,omitempty"`
	Status ECDSClusterStatus `json:"status,omitempty"`
}

// Init is used to initialize an ECDSCluster. Namely, if the state has not been
// specified, it will be set
func (obj *ECDSCluster) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialized
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateUninitialized
	}
	obj.Status.Satisfied = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *ECDSCluster) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed ECDSCluster
func ToECDSCluster(u *unstructured.Unstructured) *ECDSCluster {
	var obj *ECDSCluster
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &ECDSCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed ECDSCluster into an unstructured.Unstructured
func (obj *ECDSCluster) FromECDSCluster() *unstructured.Unstructured {
	u := NewECDSClusterVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the chart group has been deleted
func (obj *ECDSCluster) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsSatisfied returns true if the chart's actual state meets its target state
func (obj *ECDSCluster) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

func (obj *ECDSCluster) GetName() string {
	return obj.ObjectMeta.Name
}

// Returns a GKV for ECDSCluster
func NewECDSClusterVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("kubedgeoperators.kubedge.cloud/v1alpha1")
	u.SetKind("ECDSCluster")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ECDSClusterList contains a list of ECDSCluster
type ECDSClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ECDSCluster `json:"items"`
}

// Convert an unstructured.Unstructured into a typed ECDSClusterList
func ToECDSClusterList(u *unstructured.Unstructured) *ECDSClusterList {
	var obj *ECDSClusterList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &ECDSClusterList{}
	}
	return obj
}

// Convert a typed ECDSClusterList into an unstructured.Unstructured
func (obj *ECDSClusterList) FromECDSClusterList() *unstructured.Unstructured {
	u := NewECDSClusterListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *ECDSClusterList) Equivalent(other *ECDSClusterList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for ECDSClusterList
func NewECDSClusterListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("kubedgeoperators.kubedge.cloud/v1alpha1")
	u.SetKind("ECDSClusterList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&ECDSCluster{}, &ECDSClusterList{})
}
