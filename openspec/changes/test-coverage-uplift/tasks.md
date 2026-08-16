# Tasks — test-coverage-uplift

- [x] Inventory packages with no test files (`go test ./... -cover` → note 0.0% pkgs). — before: every pkg 0% except `pkg/kubedgemanager`.
- [~] Add tests for the manager/renderer: owner-ref stamping + lifecycle-state mapping. — lifecycle-state mapping covered via `KubedgeStatus.ComputeActualState` tests (common_types); owner-ref renderer already had `ownerref_renderer_test.go`. Deeper manager tests deferred to a stricter pass.
- [x] Add tests for the conditions package: Get/Has/IsTrue/IsFalse edge cases (absent condition). — `pkg/conditions/conditions_test.go` (Get/Has/IsTrue/IsFalse/IsUnknown/Set/GetMessage/GetReason).
- [x] Add tests for the shared status/state helpers in common_types. — `common_types_test.go`: ComputeActualState table, SetCondition/RemoveCondition, FindCondition.
- [x] Ensure `go test ./... -race` stays green; record the coverage delta. — green. Delta: `pkg/conditions` 0% → 13.8%, `pkg/apis/.../v1alpha1` 0% → 6.2%.

Note (first-pass): coverage is intentionally seeded, not exhaustive; a stricter pass will
raise it on the manager engine and reconciler.
