# kubedgemanager-engine Specification

## Purpose

The shared rendering/lifecycle engine the operators use to turn a kubedge CR into
managed Kubernetes resources. Implemented in `pkg/kubedgemanager/` (manager factory,
manager, helm-derived owner-reference renderer, watches, patch helpers).

## Requirements

### Requirement: A factory produces a per-kind resource manager

The package SHALL expose `KubedgeResourceManagerFactory` with a constructor per kind
(`NewArpscanManager`, `NewECDSClusterManager`, `NewMMESimManager`, `NewEMBBSliceManager`,
…) created from a controller-runtime `manager.Manager` via `NewManagerFactory`. Each
returns a `KubedgeResourceManager` that controls that kind's phase of the service
lifecycle.

#### Scenario: consumer builds a manager for its kind
- **WHEN** an operator constructs the factory with its controller manager and calls the constructor for its kind
- **THEN** it receives a `KubedgeResourceManager` bound to that CR instance

### Requirement: Resources are rendered with owner references to the CR

The engine SHALL render managed resources through a `KubedgeResourceRenderer`
(`ownerref_renderer.go`, derived from Helm's renderer) that stamps owner references
back to the owning kubedge CR, so managed objects are garbage-collected with the CR.

#### Scenario: rendered objects are owned
- **WHEN** the engine renders resources for a CR
- **THEN** each rendered object carries an owner reference to that CR

### Requirement: The manager reports lifecycle state back onto the CR status

The manager SHALL drive and report the resource lifecycle using the shared
`KubedgeResourceState` values (Deployed/Failed/Pending/…) so the reconciler can set
`status.actualState` and `satisfied` from the manager's result.

#### Scenario: state surfaces after an apply
- **WHEN** the manager applies rendered resources
- **THEN** it returns a lifecycle state the reconciler maps onto `.status.actualState`
