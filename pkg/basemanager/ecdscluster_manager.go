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

	av1 "github.com/kubedge/kubedge-operator-base/pkg/apis/kubedgeoperators/v1alpha1"
)

type upgrademanager struct {
	phasemanager

	spec   av1.ECDSClusterSpec
	status *av1.ECDSClusterStatus
}

// Sync retrieves from K8s the sub resources (Workflow, Job, ....) attached to this ECDSCluster CR
func (m *upgrademanager) Sync(ctx context.Context) error {

	m.deployedSubResourceList = av1.NewSubResourceList(m.phaseNamespace, m.phaseName)

	rendered, deployed, err := m.sync(ctx)
	if err != nil {
		return err
	}

	m.deployedSubResourceList = deployed
	if len(rendered.Items) != len(deployed.Items) {
		m.isInstalled = false
		m.isUpdateRequired = false
	} else {
		m.isInstalled = true
		m.isUpdateRequired = false
	}

	return nil
}

// InstallResource creates K8s sub resources (Workflow, Job, ....) attached to this ECDSCluster CR
func (m upgrademanager) InstallResource(ctx context.Context) (*av1.SubResourceList, error) {
	return m.installResource(ctx)
}

// InstallResource updates K8s sub resources (Workflow, Job, ....) attached to this ECDSCluster CR
func (m upgrademanager) UpdateResource(ctx context.Context) (*av1.SubResourceList, *av1.SubResourceList, error) {
	return m.updateResource(ctx)
}

// ReconcileResource creates or patches resources as necessary to match this ECDSCluster CR
func (m upgrademanager) ReconcileResource(ctx context.Context) (*av1.SubResourceList, error) {
	return m.reconcileResource(ctx)
}

// UninstallResource delete K8s sub resources (Workflow, Job, ....) attached to this ECDSCluster CR
func (m upgrademanager) UninstallResource(ctx context.Context) (*av1.SubResourceList, error) {
	return m.uninstallResource(ctx)
}
