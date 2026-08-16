## ADDED Requirements

### Requirement: Scheme registration uses apimachinery runtime.NewSchemeBuilder

The scheme registration in `pkg/apis/kubedgeoperators/v1alpha1/register.go` SHALL use
`k8s.io/apimachinery/pkg/runtime`'s `NewSchemeBuilder` instead of the deprecated
controller-runtime `scheme.Builder` (SA1019), and the corresponding
`//nolint:staticcheck` suppression SHALL be removed. Because registration is part of the
shared surface consumers import, this migration SHALL be released as a new base tag.

#### Scenario: registration no longer uses the deprecated builder
- **WHEN** `golangci-lint run` executes over `register.go`
- **THEN** no SA1019 `scheme.Builder` deprecation is reported and no nolint suppresses it

### Requirement: Endpoint data uses EndpointSlice, not the deprecated Endpoints type

Any use of `corev1.Endpoints` SHALL be migrated to EndpointSlice, updating the watched
types and RBAC accordingly, and the corresponding `//nolint` + `TODO(deps-bump)` SHALL be
removed.

#### Scenario: no deprecated Endpoints usage remains
- **WHEN** the tree is linted after migration
- **THEN** no `corev1.Endpoints` deprecation is reported
