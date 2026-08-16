// NOTE: Boilerplate only.  Ignore this file.

// Package v1alpha1 contains API Schema definitions for the kubedgeoperator v1alpha1 API group
// +k8s:deepcopy-gen=package,register
// +groupName=kubedgeoperators.kubedge.cloud
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "kubedgeoperators.kubedge.cloud", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	// TODO(deps-bump): controller-runtime scheme.Builder is deprecated as of the
	// k8s v0.36 bump; migrate scheme registration to apimachinery's
	// runtime.NewSchemeBuilder as a follow-up. Still functional for now.
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion} //nolint:staticcheck // SA1019: deprecated but functional; migration tracked separately
)
