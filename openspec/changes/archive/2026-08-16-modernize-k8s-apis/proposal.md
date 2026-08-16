# Modernize deprecated Kubernetes APIs

## Why

The k8s deps bump (to v0.36.3 / controller-runtime v0.24.1) surfaced deprecations that
were suppressed with documented `//nolint:staticcheck` + `TODO(deps-bump)` to keep the
release green. They still work but are on borrowed time and should be migrated properly.

## What changes

- `corev1.Endpoints` → EndpointSlice (this changes what the operator watches/consumes for
  endpoint data; v1.33+ direction).
- controller-runtime `scheme.Builder` (SA1019 in `register.go`) → apimachinery
  `runtime.NewSchemeBuilder`.
- Remove the corresponding `//nolint` suppressions once migrated.

## Non-goals

- Any behavior change beyond API-surface migration; the reconcile semantics stay the same.

## Impact

Base is the shared API/registration surface, so the `scheme.Builder` change in particular
ripples to consumers — coordinate the version bump. Fleet-wide follow-up.
