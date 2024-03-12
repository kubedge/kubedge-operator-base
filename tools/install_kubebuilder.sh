#!/usr/bin/env bash

[[ -f bin/kubebuilder ]] && exit 0

version=3.14.0
arch=amd64

mkdir -p ./bin
cd ./bin
curl -L -O "https://github.com/kubernetes-sigs/kubebuilder/releases/download/v${version}/kubebuilder_linux_${arch}"
chmod u+x kubebuilder_linux_${arch}

# curl -L -O "https://github.com/kubernetes-sigs/kubebuilder/releases/download/v${version}/kubebuilder_${version}_linux_${arch}.tar.gz"

# tar -zxvf kubebuilder_${version}_linux_${arch}.tar.gz
# mv kubebuilder_${version}_linux_${arch}/bin/* bin

# rm kubebuilder_${version}_linux_${arch}.tar.gz
# rm -r kubebuilder_${version}_linux_${arch}
