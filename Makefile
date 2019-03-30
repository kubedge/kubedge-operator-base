
# Image URL to use all building/pushing image targets
COMPONENT        ?= baseoperator-operator
VERSION_V1       ?= 0.1.0
DHUBREPO         ?= kubedge1/${COMPONENT}-dev
DOCKER_NAMESPACE ?= kubedge1
IMG_V1           ?= ${DHUBREPO}:v${VERSION_V1}

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

# Run tests
unittest: setup fmt vet-v1
	echo "sudo systemctl stop kubelet"
	echo -e 'docker stop $$(docker ps -qa)'
	echo -e 'export PATH=$${PATH}:/usr/local/kubebuilder/bin'
	mkdir -p config/crds
	cp chart/templates/*v1alpha1* config/crds/
	go test ./pkg/... ./cmd/... -coverprofile cover.out

# Run go fmt against code
fmt: setup
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet-v1: fmt
	go vet -composites=false -tags=v1 ./pkg/... ./cmd/...

# Generate code
generate: setup
	go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd --output-dir ./chart/templates/ --domain kubedge.cloud
	go run vendor/k8s.io/code-generator/cmd/deepcopy-gen/main.go --input-dirs github.com/kubedge/kubedge-operator-base/pkg/apis/baseoperator/v1alpha1 -O zz_generated.deepcopy --bounding-dirs github.com/kubedge/kubedge-operator-base/pkg/apis

# Build the docker image
docker-build: fmt docker-build-v1

docker-build-v1: vet-v1
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/_output/bin/baseoperator-operator -gcflags all=-trimpath=${GOPATH} -asmflags all=-trimpath=${GOPATH} -tags=v1 ./cmd/...
	docker build . -f build/Dockerfile -t ${IMG_V1}
	docker tag ${IMG_V1} ${DHUBREPO}:latest


# Push the docker image
docker-push: docker-push-v1

docker-push-v1:
	docker push ${IMG_V1}

# Run against the configured Kubernetes cluster in ~/.kube/config
install: install-v1

purge: setup
	helm delete --purge baseoperator-operator

install-v1: docker-build-v1
	helm install --name baseoperator-operator chart --set images.tags.operator=${IMG_V1}

# Deploy and purge procedure which do not rely on helm itself
install-kubectl: docker-build
	kubectl apply -f ./chart/templates/baseoperator_v1alpha1_openstackbackup.yaml
	kubectl apply -f ./chart/templates/baseoperator_v1alpha1_openstackdeployment.yaml
	kubectl apply -f ./chart/templates/baseoperator_v1alpha1_openstackrestore.yaml
	kubectl apply -f ./chart/templates/baseoperator_v1alpha1_openstackrollback.yaml
	kubectl apply -f ./chart/templates/baseoperator_v1alpha1_openstackupgrade.yaml
	kubectl apply -f ./chart/templates/baseoperator_v1alpha1_oslc.yaml
	kubectl apply -f ./chart/templates/role_binding.yaml
	kubectl apply -f ./chart/templates/role.yaml
	kubectl apply -f ./chart/templates/service_account.yaml
	kubectl apply -f ./chart/templates/argo_baseoperator_role.yaml
	kubectl create -f deploy/operator.yaml

purge-kubectl: setup
	kubectl delete -f deploy/operator.yaml
	kubectl delete -f ./chart/templates/baseoperator_v1alpha1_openstackbackup.yaml
	kubectl delete -f ./chart/templates/baseoperator_v1alpha1_openstackdeployment.yaml
	kubectl delete -f ./chart/templates/baseoperator_v1alpha1_openstackrestore.yaml
	kubectl delete -f ./chart/templates/baseoperator_v1alpha1_openstackrollback.yaml
	kubectl delete -f ./chart/templates/baseoperator_v1alpha1_openstackupgrade.yaml
	kubectl delete -f ./chart/templates/baseoperator_v1alpha1_oslc.yaml
	kubectl delete -f ./chart/templates/role_binding.yaml
	kubectl delete -f ./chart/templates/role.yaml
	kubectl delete -f ./chart/templates/service_account.yaml
	kubectl delete -f ./chart/templates/argo_baseoperator_role.yaml
