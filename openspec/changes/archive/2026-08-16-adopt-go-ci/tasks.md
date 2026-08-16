# Tasks — adopt-go-ci

- [x] From the meta session, run `/alemax:update-skills` so the class-M set (incl. `ci.yml`) is staged for this repo. — done meta-side; delivered branch `meta-broadcast/deliver-class-m-templates-ciyml--hygiene` + `.local/HANDOFF.md`.
- [x] In this repo's session, run `/alemax:complete-update` to apply the update branch onto the working branch. — cherry-picked onto `apply/adopt-go-ci`; all 10 files clean adds.
- [x] Confirm `.github/workflows/ci.yml` present and its jobs gate on `go.mod`. — present; go jobs `if: needs.detect.outputs.go == 'true'`; trimmed `type-check`; added `gomod` to dependabot.
- [x] Trial push the branch; confirm `go-build`/`go-vet`/`go-test`/`golangci-lint` are green. — PR #7, run 31960755943: **all 8 checks green**. Required fixing the shipped golangci-lint pin (v1.61.0→v2.12.2, action v6→v7) — v1.61 can't lint a go 1.26 module; flagged for meta re-broadcast.
- [x] Confirm the rest of class-M landed: `.editorconfig`, `.gitattributes`, `.github/*`, `dependabot.yml`, `.pre-commit-config.yaml`, `bin/set-secret.sh`. — all present on `main` (9519f7f).
