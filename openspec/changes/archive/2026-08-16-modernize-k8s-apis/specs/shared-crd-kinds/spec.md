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

### Requirement: No deprecated corev1.Endpoints usage

The api package SHALL NOT use the deprecated `corev1.Endpoints` type. The
`IsServiceReady`/`IsServiceFailedOrError` helpers that relied on it were **dead code** —
never invoked, because the dependency-readiness dispatch treats `Service` as
ready-by-presence — and SHALL be removed together with the corresponding `//nolint` +
`TODO(deps-bump)` suppression. Service/Deployment/StatefulSet readiness semantics are
unchanged (no behavior change). (No EndpointSlice reader is introduced: nothing consumed
endpoint data, so there is no watch/RBAC surface to migrate.)

#### Scenario: no deprecated Endpoints usage remains
- **WHEN** the tree is linted
- **THEN** no `corev1.Endpoints` deprecation is reported and no nolint suppresses one
