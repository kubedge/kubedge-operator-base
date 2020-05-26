// +build tools

package tools

import (
	// These imports are all tools used in the building and testing process
	_ "k8s.io/kube-openapi/cmd/openapi-gen"
        _ "sigs.k8s.io/controller-tools/cmd/controller-gen"
        _ "github.com/golangci/golangci-lint/cmd/golangci-lint"
        _ "sigs.k8s.io/kind"
        _ "github.com/instrumenta/kubeval"
        _ "k8s.io/kube-openapi/cmd/openapi-gen"
)
