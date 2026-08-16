# Tasks — modernize-k8s-apis

- [ ] Grep the tree for the deferred suppressions: `//nolint:staticcheck` + `TODO(deps-bump)`.
- [ ] Migrate `scheme.Builder` → `runtime.NewSchemeBuilder` in `pkg/apis/kubedgeoperators/v1alpha1/register.go`; drop the SA1019 nolint.
- [ ] Migrate `corev1.Endpoints` usage → EndpointSlice; adjust the watched types/RBAC.
- [ ] `go build ./... && go vet ./... && go test ./... -race && golangci-lint run --max-same-issues 0 --max-issues-per-linter 0 ./...` all green.
- [ ] Re-tag base (new `v0.1.<k8s-minor>-kubedge.<buildday>`) and have consumers realign, since `register.go` is part of the shared surface.
