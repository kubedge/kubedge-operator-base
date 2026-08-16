// NOTE: Boilerplate only.  Ignore this file.

// Package v1alpha1 contains API Schema definitions for the kubedgeoperator v1alpha1 API group
// +k8s:deepcopy-gen=package,register
// +groupName=kubedgeoperators.kubedge.cloud
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "kubedgeoperators.kubedge.cloud", Version: "v1alpha1"}

	// SchemeBuilder collects the functions that add this group's types to a Scheme.
	// Uses apimachinery's runtime.NewSchemeBuilder so this api package depends only on
	// k8s.io/apimachinery (no controller-runtime), per the SA1019 guidance.
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)

	// AddToScheme adds the types in this group-version to a Scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

// addKnownTypes registers this group's types with the given Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Arpscan{}, &ArpscanList{},
		&ECDSCluster{}, &ECDSClusterList{},
		&MMESim{}, &MMESimList{},
		&EMBBSlice{}, &EMBBSliceList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
