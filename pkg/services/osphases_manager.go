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

package services

import (
	"context"
	av1 "github.com/kubedge/kubedge-operator-base/pkg/apis/kubedgeoperators/v1alpha1"
)

// ArpscanManager manages the Arpscan Phase of an OpenstackServiceLifeCycle
type ArpscanManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.SubResourceList, error)
	UpdateResource(context.Context) (*av1.SubResourceList, *av1.SubResourceList, error)
	ReconcileResource(context.Context) (*av1.SubResourceList, error)
	UninstallResource(context.Context) (*av1.SubResourceList, error)
}

// ECDSClusterManager manages the ECDSCluster Phase of an OpenstackServiceLifeCycle
type ECDSClusterManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.SubResourceList, error)
	UpdateResource(context.Context) (*av1.SubResourceList, *av1.SubResourceList, error)
	ReconcileResource(context.Context) (*av1.SubResourceList, error)
	UninstallResource(context.Context) (*av1.SubResourceList, error)
}

// MMESimManager manages the MMESim Phase of an OpenstackServiceLifeCycle
type MMESimManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.SubResourceList, error)
	UpdateResource(context.Context) (*av1.SubResourceList, *av1.SubResourceList, error)
	ReconcileResource(context.Context) (*av1.SubResourceList, error)
	UninstallResource(context.Context) (*av1.SubResourceList, error)
}
