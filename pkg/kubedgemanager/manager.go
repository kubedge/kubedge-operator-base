// Copyright 2019 The Kubedge Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package basemanager

import (
	"context"
	"fmt"

	av1 "github.com/kubedge/kubedge-operator-base/pkg/apis/kubedgeoperators/v1alpha1"

	v1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	yaml "sigs.k8s.io/yaml"
)

// KubedgeResourceManager provides the interface for base manaager.
type KubedgeResourceManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Render(ctx context.Context) (*av1.SubResourceList, error)
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.SubResourceList, error)
	UpdateResource(context.Context) (*av1.SubResourceList, *av1.SubResourceList, error)
	ReconcileResource(context.Context) (*av1.SubResourceList, error)
	UninstallResource(context.Context) (*av1.SubResourceList, error)
}

// KubedgeBaseManager provides the default implementation.
type KubedgeBaseManager struct {
	KubeClient     client.Client
	Renderer       KubedgeResourceRenderer
	OwnerRefs      []metav1.OwnerReference
	PhaseName      string
	PhaseNamespace string
	Source         *av1.KubedgeSource

	IsInstalledFlag         bool
	IsUpdateRequiredFlag    bool
	DeployedSubResourceList *av1.SubResourceList
}

// ResourceName returns the name of the release.
func (m KubedgeBaseManager) ResourceName() string {
	return m.PhaseName
}

// IsInstalled indicates with the resources have been installed.
func (m KubedgeBaseManager) IsInstalled() bool {
	return m.IsInstalledFlag
}

// IsUpdateRequired indicates with the resources have been installed.
func (m KubedgeBaseManager) IsUpdateRequired() bool {
	return m.IsUpdateRequiredFlag
}

// Render a chart or just a file
func (m KubedgeBaseManager) Render(ctx context.Context) (*av1.SubResourceList, error) {
	return m.Renderer.RenderFile(m.PhaseName, m.PhaseNamespace, m.Source.Location)
}

// BaseSync retrieves from K8s the sub resources (Workflow, Job, ....) attached to this Oslc CR
func (m *KubedgeBaseManager) BaseSync(ctx context.Context) error {
	m.DeployedSubResourceList = av1.NewSubResourceList(m.PhaseNamespace, m.PhaseName)

	rendered, alreadyDeployed, err := m.internalSync(ctx)
	if err != nil {
		return err
	}

	m.DeployedSubResourceList = alreadyDeployed
	if len(rendered.GetDependentResources()) != len(alreadyDeployed.GetDependentResources()) {
		m.IsInstalledFlag = false
		m.IsUpdateRequiredFlag = false
	} else {
		m.IsInstalledFlag = true
		m.IsUpdateRequiredFlag = false
	}

	return nil
}

// internalSync attempts to compare the K8s object present with the rendered objects
func (m KubedgeBaseManager) internalSync(ctx context.Context) (*av1.SubResourceList, *av1.SubResourceList, error) {
	deployed := av1.NewSubResourceList(m.PhaseNamespace, m.PhaseName)

	rendered, err := m.Render(ctx)
	if err != nil {
		return nil, deployed, err
	}

	errs := make([]error, 0)

	for _, renderedResource := range rendered.Items {
		existingResource := unstructured.Unstructured{}
		existingResource.SetAPIVersion(renderedResource.GetAPIVersion())
		existingResource.SetKind(renderedResource.GetKind())
		existingResource.SetName(renderedResource.GetName())
		existingResource.SetNamespace(renderedResource.GetNamespace())

		err := m.KubeClient.Get(context.TODO(), types.NamespacedName{Name: existingResource.GetName(), Namespace: existingResource.GetNamespace()}, &existingResource)
		if err != nil {
			if !apierrors.IsNotFound(err) {
				log.Error(err, "Can't Retrieve Resource", "kind", existingResource.GetKind(), "name", existingResource.GetName())
				errs = append(errs, err)
			}
		} else {
			deployed.Items = append(deployed.Items, existingResource)
		}
	}

	if !deployed.CheckOwnerReference(m.OwnerRefs) {
		return rendered, nil, ErrOwnershipMismatch
	}

	// TODO(jeb): not sure this is right
	// if len(errs) != 0 {
	//	return rendered, deployed, errs[0]
	// }
	return rendered, deployed, nil

}

// BaseInstallResource performs a "helm install" equivalent
// It creates K8s sub resources (Workflow, Job, ....) attached to this Kubedge CR
func (m KubedgeBaseManager) BaseInstallResource(ctx context.Context) (*av1.SubResourceList, error) {

	errs := make([]error, 0)
	created := av1.NewSubResourceList(m.PhaseNamespace, m.PhaseName)

	if m.DeployedSubResourceList == nil {
		// There was an error during SyncResource
		return created, ErrInstall
	}

	rendered, err := m.Render(ctx)
	if err != nil {
		return m.DeployedSubResourceList, err
	}

	for _, toCreate := range rendered.Items {
		err := m.KubeClient.Create(context.TODO(), &toCreate)
		if err != nil {
			if !apierrors.IsAlreadyExists(err) {
				blob, _ := yaml.Marshal(toCreate)
				thestr := fmt.Sprintf("[%s]", string(blob))
				log.Error(err, "Can't Create Resource", "kind", toCreate.GetKind(), "name", toCreate.GetName(), "resource", thestr)
				errs = append(errs, err)
			} else {
				// Should consider as just created by us
				log.Info("Created Resource", "kind", toCreate.GetKind(), "name", toCreate.GetName())
				created.Items = append(created.Items, toCreate)
			}
		} else {
			created.Items = append(created.Items, toCreate)
		}
	}

	if len(errs) != 0 {
		return created, errs[0]
	}
	return created, nil
}

// BasedUpdateResource performs a "helm upgrade" equivalent. Most likely the Values field in the Kubedge Resource.
// It updates K8s sub resources (Workflow, Job, ....) attached to this CR
func (m KubedgeBaseManager) BaseUpdateResource(ctx context.Context) (*av1.SubResourceList, *av1.SubResourceList, error) {
	updated := av1.NewSubResourceList(m.PhaseNamespace, m.PhaseName)

	if m.DeployedSubResourceList == nil {
		// There was an error during SyncResource
		return m.DeployedSubResourceList, updated, ErrUpdate
	}

	// TODO(JEB): Big hack. ReconcileResource should do more
	m.DeployedSubResourceList.DeepCopyInto(updated)

	return m.DeployedSubResourceList, updated, nil
}

// BaseReconcileResource creates or patches resources as necessary to match this deployed Kubedge CR
func (m KubedgeBaseManager) BaseReconcileResource(ctx context.Context) (*av1.SubResourceList, error) {
	errs := make([]error, 0)
	reconciled := av1.NewSubResourceList(m.PhaseNamespace, m.PhaseName)

	if m.DeployedSubResourceList == nil {
		// There was an error during SyncResource
		return reconciled, ErrReconcile
	}

	rendered, err := m.Render(ctx)
	if err != nil {
		return m.DeployedSubResourceList, err
	}

	for _, renderedResource := range rendered.Items {
		existingResource := unstructured.Unstructured{}
		existingResource.SetAPIVersion(renderedResource.GetAPIVersion())
		existingResource.SetKind(renderedResource.GetKind())
		existingResource.SetName(renderedResource.GetName())
		existingResource.SetNamespace(renderedResource.GetNamespace())

		err := m.KubeClient.Get(context.TODO(), types.NamespacedName{Name: existingResource.GetName(), Namespace: existingResource.GetNamespace()}, &existingResource)
		if err != nil {
			if !apierrors.IsNotFound(err) {
				log.Error(err, "Can't Retrieve Resource", "kind", existingResource.GetKind(), "name", existingResource.GetName())
				errs = append(errs, err)
			}
		} else {
			if renderedResource.GetKind() == "StatefulSet" {

				existingStatefulSet := v1.StatefulSet{}
				err1 := runtime.DefaultUnstructuredConverter.FromUnstructured(existingResource.UnstructuredContent(), &existingStatefulSet)
				if err1 != nil {
					log.Error(err1, "error converting existingResource from Unstructured")
				}

				renderedStatefulSet := v1.StatefulSet{}
				err2 := runtime.DefaultUnstructuredConverter.FromUnstructured(renderedResource.UnstructuredContent(), &renderedStatefulSet)
				if err2 != nil {
					log.Error(err2, "error converting from renderedResource Unstructured")
				}

				if (existingStatefulSet.Spec.Replicas != nil) && (renderedStatefulSet.Spec.Replicas != nil) && (*existingStatefulSet.Spec.Replicas != *renderedStatefulSet.Spec.Replicas) {
					// A merge patch will preserve other fields modified at runtime.
					patch := client.MergeFrom(existingResource.DeepCopy())

					// JEB: Seems convolutedy, probably needs to go through golang training again
					existingStatefulSet.Spec.Replicas = new(int32)
					*existingStatefulSet.Spec.Replicas = *renderedStatefulSet.Spec.Replicas
					unst, err3 := runtime.DefaultUnstructuredConverter.ToUnstructured(&existingStatefulSet)
					if err3 != nil {
						log.Error(err3, "error converting to Unstructured")
					}
					existingResource.SetUnstructuredContent(unst)

					err := m.KubeClient.Patch(context.TODO(), &existingResource, patch)
					if err != nil {
						if !apierrors.IsNotFound(err) {
							log.Error(err, "Can't Patch Resource", "kind", existingResource.GetKind(), "name", existingResource.GetName())
							errs = append(errs, err)
						}
					} else {
						log.Info("Patched Resource", "kind", existingResource.GetKind(), "name", existingResource.GetName())
						reconciled.Items = append(reconciled.Items, existingResource)
					}
				} else {
					reconciled.Items = append(reconciled.Items, existingResource)
				}
			} else {
				reconciled.Items = append(reconciled.Items, existingResource)
			}
		}
	}

	if len(errs) != 0 {
		return reconciled, errs[0]
	}
	return reconciled, nil
}

// BaseUninstallResource performs a "helm delete" equivalent
// It delete K8s sub resources (Workflow, Job, ....) attached to this Kubedge CR
func (m KubedgeBaseManager) BaseUninstallResource(ctx context.Context) (*av1.SubResourceList, error) {
	errs := make([]error, 0)
	notdeleted := av1.NewSubResourceList(m.PhaseNamespace, m.PhaseName)

	if m.DeployedSubResourceList == nil {
		// There was an error during SyncResource
		return notdeleted, ErrUninstall
	}

	for _, toDelete := range m.DeployedSubResourceList.Items {
		err := m.KubeClient.Delete(context.TODO(), &toDelete)
		if err != nil {
			if !apierrors.IsNotFound(err) {
				log.Error(err, "Can't Delete Resource", "kind", toDelete.GetKind(), "name", toDelete.GetName())
				errs = append(errs, err)
				notdeleted.Items = append(notdeleted.Items, toDelete)
			}
		}
	}

	if len(errs) != 0 {
		return notdeleted, errs[0]
	}
	return notdeleted, nil
}
