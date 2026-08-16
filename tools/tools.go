//go:build tools
// +build tools

package tools

import (
	// controller-gen is the only build tool this module vendors; it is invoked by
	// the Makefile (generate-go / generate-manifests). golangci-lint is provided by
	// the system toolchain; kind/kubeval/openapi-gen were unused and dropped during
	// the k8s v0.36 bump (kubeval pulled a conflicting legacy genproto).
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
