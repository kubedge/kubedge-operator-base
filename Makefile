
# Image URL to use all building/pushing image targets
COMPONENT        ?= kubedge-base-operator
VERSION          ?= 0.1.36-kubedge.20260816
DHUBREPO         ?= kubedge1/${COMPONENT}-dev
DOCKER_NAMESPACE ?= kubedge1
IMG              ?= ${DHUBREPO}:v${VERSION}
# Target platforms for the multi-arch image. arm64 is primary (Pi armv8 + Apple Silicon);
# amd64 kept for x86 clusters. Override with `make docker-buildx PLATFORMS=linux/arm64`.
PLATFORMS        ?= linux/arm64,linux/amd64

BINDIR           := bin
TOOLS_DIR        := tools
TOOLS_BIN_DIR    := $(TOOLS_DIR)/bin

# Binaries.
CONTROLLER_GEN   := $(TOOLS_BIN_DIR)/controller-gen

all: docker-build

setup:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif
	# dep ensure

clean:
	rm -fr vendor
	rm -fr cover.out
	rm -fr build/_output
	rm -fr config/crds


## --------------------------------------
## Tooling Binaries
## --------------------------------------

$(CONTROLLER_GEN): $(TOOLS_DIR)/go.mod # Build controller-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BINDIR)/controller-gen sigs.k8s.io/controller-tools/cmd/controller-gen

.PHONY: install-tools
install-tools: $(CONTROLLER_GEN)

## --------------------------------------
## Testing
## --------------------------------------

# Run tests
unittest: setup fmt vet-v1
	echo "sudo systemctl stop kubelet"
	echo -e 'docker stop $$(docker ps -qa)'
	echo -e 'export PATH=$${PATH}:/usr/local/kubebuilder/bin'
	mkdir -p config/crds
	cp chart/templates/*v1alpha1* config/crds/
	go test ./pkg/... ./cmd/... -coverprofile cover.out

## --------------------------------------
## Linting
## --------------------------------------

# Run go fmt against code
fmt: setup
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet-v1: fmt
	go vet -composites=false -tags=v1 ./pkg/... ./cmd/...

## --------------------------------------
## Generate
## --------------------------------------

.PHONY: modules
modules: ## Runs go mod to ensure proper vendoring.
	go mod tidy
	cd $(TOOLS_DIR); go mod tidy

.PHONY: generate
generate: ## Generate code
	$(MAKE) generate-go
	$(MAKE) generate-manifests

.PHONY: generate-go
generate-go: $(CONTROLLER_GEN)
	GO111MODULE=on $(CONTROLLER_GEN) object paths=./pkg/apis/kubedgeoperators/... output:object:dir=./pkg/apis/kubedgeoperators/v1alpha1 output:none

.PHONY: generate-manifests
generate-manifests: $(CONTROLLER_GEN) ## Generate manifests e.g. CRD, RBAC etc.
	mkdir -p chart/templates/
	# GO111MODULE=on $(CONTROLLER_GEN) crd paths=./pkg/apis/kubedgeoperators/... crd:trivialVersions=true output:crd:dir=./chart/templates/ output:none
	GO111MODULE=on $(CONTROLLER_GEN) crd:generateEmbeddedObjectMeta=true paths=./pkg/apis/kubedgeoperators/... output:crd:dir=./chart/templates/ output:none

# Build a multi-arch image (manifest list) and push it. The binary is compiled per
# TARGETARCH inside build/Dockerfile, so a single command covers every platform.
.PHONY: docker-buildx
docker-buildx: vet-v1
	docker buildx build --platform ${PLATFORMS} -f build/Dockerfile -t ${IMG} -t ${DHUBREPO}:latest --push .

# Local single-arch image for dev (loads into the local docker; no push).
.PHONY: docker-build
docker-build: vet-v1
	docker buildx build --load -f build/Dockerfile -t ${IMG} .

# Push is folded into docker-buildx (--push); kept as an alias for muscle memory.
.PHONY: docker-push
docker-push: docker-buildx

# Run against the configured Kubernetes cluster in ~/.kube/config
.PHONY: install
install: docker-buildx
	helm install kubedge-base-operator chart --set images.tags.operator=${IMG}

.PHONY: purge
purge:
	helm uninstall kubedge-base-operator
