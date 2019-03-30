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
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	av1 "github.com/kubedge/kubedge-operator-base/pkg/apis/baseoperator/v1alpha1"
	lcmif "github.com/kubedge/kubedge-operator-base/pkg/services"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type managerFactory struct {
	kubeClient client.Client
}

// Simple function to init the renderFiles passed to the helm renderer
func initRenderFiles(stage av1.OslcPhase) []string {
	renderFiles := make([]string, 0)
	return renderFiles
}

// Simple function to init the renderValues passed to the helm renderer
func initRenderValues(stage av1.OslcPhase) map[string]interface{} {
	oslcValues := map[string]interface{}{}
	oslcValues["stage"] = stage.String()
	renderValues := map[string]interface{}{}
	renderValues["oslc"] = oslcValues
	return renderValues
}

// NewManagerFactory returns a new factory.
func NewManagerFactory(mgr manager.Manager) lcmif.PhaseManagerFactory {
	return &managerFactory{kubeClient: mgr.GetClient()}
}

// NewTestPhaseManager returns a new manager capable of controlling TestPhase phase of the service lifecyle
func (f managerFactory) NewTestPhaseManager(r *av1.TestPhase) lcmif.TestPhaseManager {
	controllerRef := metav1.NewControllerRef(r, r.GroupVersionKind())
	ownerRefs := []metav1.OwnerReference{
		*controllerRef,
	}

	renderFiles := initRenderFiles(av1.PhaseTest)
	renderValues := initRenderValues(av1.PhaseTest)

	return &testmanager{
		phasemanager: phasemanager{
			kubeClient:     f.kubeClient,
			renderer:       NewOwnerRefRenderer(ownerRefs, "ostest", renderFiles, renderValues),
			source:         r.Spec.Source,
			phaseName:      r.GetName(),
			phaseNamespace: r.GetNamespace()},

		spec:   r.Spec,
		status: &r.Status,
	}
}

// NewUpgradePhaseManager returns a new manager capable of controlling UpgradePhase phase of the service lifecyle
func (f managerFactory) NewUpgradePhaseManager(r *av1.UpgradePhase) lcmif.UpgradePhaseManager {
	controllerRef := metav1.NewControllerRef(r, r.GroupVersionKind())
	ownerRefs := []metav1.OwnerReference{
		*controllerRef,
	}

	renderFiles := initRenderFiles(av1.PhaseUpgrade)
	renderValues := initRenderValues(av1.PhaseUpgrade)

	return &upgrademanager{
		phasemanager: phasemanager{
			kubeClient:     f.kubeClient,
			renderer:       NewOwnerRefRenderer(ownerRefs, "osupg", renderFiles, renderValues),
			source:         r.Spec.Source,
			phaseName:      r.GetName(),
			phaseNamespace: r.GetNamespace()},

		spec:   r.Spec,
		status: &r.Status,
	}
}

// NewRollbackPhaseManager returns a new manager capable of controlling RollbackPhase phase of the service lifecyle
func (f managerFactory) NewRollbackPhaseManager(r *av1.RollbackPhase) lcmif.RollbackPhaseManager {
	controllerRef := metav1.NewControllerRef(r, r.GroupVersionKind())
	ownerRefs := []metav1.OwnerReference{
		*controllerRef,
	}

	renderFiles := initRenderFiles(av1.PhaseRollback)
	renderValues := initRenderValues(av1.PhaseRollback)

	return &rollbackmanager{
		phasemanager: phasemanager{
			kubeClient:     f.kubeClient,
			renderer:       NewOwnerRefRenderer(ownerRefs, "osrbck", renderFiles, renderValues),
			source:         r.Spec.Source,
			phaseName:      r.GetName(),
			phaseNamespace: r.GetNamespace()},

		spec:   r.Spec,
		status: &r.Status,
	}
}
