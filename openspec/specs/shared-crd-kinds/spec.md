# shared-crd-kinds Specification

## Purpose

Defines the Custom Resource API types that the kubedge operator family shares. This is
the contract the `-arpscan/-ecds/-embb/-mme` operators import; each operator reconciles
one of these kinds. Implemented in `pkg/apis/kubedgeoperators/v1alpha1/`.
## Requirements
### Requirement: One registered API group/version for all kubedge operator kinds

The package SHALL register all kubedge operator CRD kinds under a single
GroupVersion `kubedgeoperators.kubedge.cloud/v1alpha1` (`register.go`,
`SchemeGroupVersion`), exposing an `AddToScheme` builder that consumers add to their
controller-runtime scheme.

#### Scenario: consumer registers the shared kinds
- **WHEN** a consumer operator adds this package's scheme builder to its manager scheme
- **THEN** the kinds Arpscan, ECDSCluster, MMESim, EMBBSlice, OSLC become available under `kubedgeoperators.kubedge.cloud/v1alpha1`

### Requirement: Each kind follows a common declarative target/actual state model

Every kind SHALL embed the shared `KubedgeSpec` (carrying `targetState`) and
`KubedgeStatus` (carrying `actualState`, `satisfied`, and a `conditions` list), so all
kinds share one reconcile contract: the controller drives `actualState` toward
`spec.targetState` and reports `satisfied` + conditions. Resource lifecycle states are
the shared `KubedgeResourceState` enum (Uninitialized/Unknown/Initialized/Deployed/
Uninstalled/Failed/Pending) in `common_types.go`.

#### Scenario: a kind exposes target vs actual state
- **WHEN** any kubedge CR is inspected
- **THEN** `.spec.targetState`, `.status.actualState`, and `.status.satisfied` are present and printed as kubectl columns

#### Scenario: shared condition vocabulary
- **WHEN** a controller sets a condition on any kind
- **THEN** it uses the shared `KubedgeConditionType/Status/Reason/Severity` types from `common_types.go`

### Requirement: Each domain kind adds its own set-spec fields over the common base

Each kind SHALL extend the common spec with domain-specific `KubedgeSetSpec` groups —
e.g. `ECDSCluster` adds `Platforms`, `FrontEnds`, `Enrichments`, `BusinessLogics`,
`LoadBalancers`. Kinds carry kubebuilder markers for `path`, `shortName`, `status`
subresource, and print columns.

#### Scenario: ECDSCluster domain shape
- **WHEN** an ECDSCluster is applied
- **THEN** its spec accepts the platform/frontend/enrichment/businesslogic/loadbalancer sets and it is addressable via shortName `ecds`

### Requirement: Deepcopy is generated, not hand-maintained

The package SHALL provide generated `zz_generated.deepcopy.go` for all kinds, produced
by `make generate` (controller-gen `object`). Consumers rely on these deepcopy methods
for the runtime.Object contract.

#### Scenario: regeneration after a type change
- **WHEN** a kind's Go struct changes and `make generate` runs
- **THEN** `zz_generated.deepcopy.go` is regenerated to match, with no hand edits

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

