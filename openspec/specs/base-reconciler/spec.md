# base-reconciler Specification

## Purpose

The shared reconcile scaffolding every kubedge operator embeds, plus the conditions
helper package. Implemented in `pkg/controller/kubedgecontroller/base_reconciler.go`
and `pkg/conditions/`.

## Requirements

### Requirement: A reusable base reconciler drives kubedge CRs

The package SHALL provide `KubedgeBaseReconciler` that consumer operators embed to
reconcile their kind, reusing common behavior instead of re-implementing the loop.

#### Scenario: consumer embeds the base reconciler
- **WHEN** an operator's reconciler embeds `KubedgeBaseReconciler`
- **THEN** it inherits the shared reconcile helpers and dependent-resource handling

### Requirement: Dependent subresources are watched via a shared predicate

The base reconciler SHALL provide `BuildDependentPredicate()` returning
controller-runtime predicate funcs so operators watch their managed subresources
consistently (create/update/delete filtering).

#### Scenario: subresource change triggers reconcile
- **WHEN** a resource owned by a kubedge CR changes and the operator registered the dependent predicate
- **THEN** the owning CR is re-queued for reconcile

### Requirement: Conditions are read/written through the shared conditions package

The `conditions` package SHALL provide getters/setters (`Get`, `Has`, `IsTrue`,
`IsFalse`, …) over the shared `KubedgeCondition` types, following the Cluster-API
conditions pattern, so all operators manage status conditions uniformly.

#### Scenario: checking a condition
- **WHEN** an operator queries a condition type on its CR via the conditions package
- **THEN** it gets a consistent True/False/absent answer without duplicating condition logic
