#!/usr/bin/env bash
# release.sh — cut a NEW, immutable Go-module release tag for kubedge-operator-base.
#
# WHY THIS EXISTS (read before "just retagging"):
#   This module is PUBLIC and goes through the Go module proxy (proxy.golang.org) and
#   checksum database (sum.golang.org). Once a version tag has been fetched even once, the
#   proxy caches its zip and the sumdb records its hash *permanently*. Therefore:
#
#     * You can NOT update the code under an existing version. Moving a tag to a new
#       commit does not reach consumers — the proxy keeps serving the cached zip, and a
#       fresh `go get` that bypasses the cache hits a "checksum mismatch" and FAILS the
#       build (the sumdb hash no longer matches the new content).
#     * The ONLY way to ship updated code is a NEW version string.
#
#   House convention encodes the k8s minor + the build date:  v0.1.<k8s-minor>-kubedge.<YYYYMMDD>
#   Same k8s minor, new day => a new, unique tag. That is the mechanism by which consumers
#   learn "the code changed": they bump to the new tag. There is no "same tag, new code".
#
# TO TELL CONSUMERS A PUBLISHED VERSION IS BAD (do not silently move a tag):
#   Add a `retract` block to go.mod and release a newer version that carries it, e.g.
#       retract v0.1.36-kubedge.20260815 // superseded: <reason>
#   `go get` / `go list -m -u all` then warns consumers and steers them off it.
#   See `--print-retract` below for a ready-to-paste snippet.
#
# USAGE:
#   bin/release.sh                 # cut v0.1.<minor>-kubedge.<today> at HEAD of main, push
#   bin/release.sh --dry-run       # show what it would do, touch nothing
#   bin/release.sh --print-retract v0.1.36-kubedge.20260815 "reason"
#
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

DRY_RUN=0
case "${1:-}" in
  --dry-run) DRY_RUN=1 ;;
  --print-retract)
    ver="${2:?usage: --print-retract <version> [reason]}"; reason="${3:-superseded}"
    echo "Add this to go.mod (then release a NEW version that includes it):"
    echo ""
    echo "retract $ver // $reason"
    exit 0
    ;;
  "" ) : ;;
  * ) echo "unknown arg: $1" >&2; exit 2 ;;
esac

# --- derive the tag from the pinned k8s minor + today -----------------------------------
kapi="$(awk '$1=="k8s.io/api"{print $2; exit}' go.mod)"     # e.g. v0.36.3
[ -n "$kapi" ] || { echo "could not read k8s.io/api version from go.mod" >&2; exit 1; }
minor="$(echo "$kapi" | cut -d. -f2)"                        # e.g. 36
today="$(date +%Y%m%d)"
TAG="v0.1.${minor}-kubedge.${today}"

echo "k8s.io/api pin : $kapi  -> minor $minor"
echo "candidate tag  : $TAG"

# --- guardrails -------------------------------------------------------------------------
branch="$(git rev-parse --abbrev-ref HEAD)"
[ "$branch" = "main" ] || { echo "refusing: not on main (on '$branch'). Release from main." >&2; exit 1; }
git diff --quiet && git diff --cached --quiet || { echo "refusing: working tree not clean." >&2; exit 1; }
git fetch origin --quiet --tags
[ "$(git rev-parse HEAD)" = "$(git rev-parse origin/main)" ] || { echo "refusing: local main != origin/main. Pull/push first." >&2; exit 1; }

# NEVER reuse or move a tag. Refuse if it exists locally, on origin, or on the proxy.
if git rev-parse -q --verify "refs/tags/$TAG" >/dev/null; then
  echo "refusing: tag $TAG already exists locally. Tags are immutable — bump the date or add a suffix." >&2; exit 1
fi
if git ls-remote --exit-code --tags origin "refs/tags/$TAG" >/dev/null 2>&1; then
  echo "refusing: tag $TAG already exists on origin. Never move a published tag." >&2; exit 1
fi
mod="$(awk 'NR==1{print $2}' go.mod)"
if curl -fs "https://proxy.golang.org/${mod}/@v/list" 2>/dev/null | grep -qx "$TAG"; then
  echo "refusing: $TAG is already known to proxy.golang.org — it is permanently immutable." >&2; exit 1
fi

# --- verify green before tagging --------------------------------------------------------
echo "verifying build/vet/test..."
go build ./... && go vet ./... && go test ./... -race >/dev/null

if [ "$DRY_RUN" = "1" ]; then
  echo "[dry-run] would: git tag -a $TAG -m ... && git push origin $TAG"
  exit 0
fi

# --- tag + push -------------------------------------------------------------------------
git tag -a "$TAG" -m "kubedge-operator-base $TAG

Kubernetes 1.${minor} ($kapi). Released from main @ $(git rev-parse --short HEAD)."
git push origin "$TAG"

echo ""
echo "Published $TAG"
echo "Consumers upgrade with:"
echo "  go get ${mod}@${TAG} && go mod tidy"
