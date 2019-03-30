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
	av1 "github.com/kubedge/kubedge-operator-base/pkg/apis/baseoperator/v1alpha1"
)

// PlanningPhaseManager manages the PlanningPhase Phase of an OpenstackServiceLifeCycle
type PlanningPhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.SubResourceList, error)
	UpdateResource(context.Context) (*av1.SubResourceList, *av1.SubResourceList, error)
	ReconcileResource(context.Context) (*av1.SubResourceList, error)
	UninstallResource(context.Context) (*av1.SubResourceList, error)
}

// InstallPhaseManager manages the InstallPhase Phase of an OpenstackServiceLifeCycle
type InstallPhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.SubResourceList, error)
	UpdateResource(context.Context) (*av1.SubResourceList, *av1.SubResourceList, error)
	ReconcileResource(context.Context) (*av1.SubResourceList, error)
	UninstallResource(context.Context) (*av1.SubResourceList, error)
}

// TestPhaseManager manages the TestPhase Phase of an OpenstackServiceLifeCycle
type TestPhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.SubResourceList, error)
	UpdateResource(context.Context) (*av1.SubResourceList, *av1.SubResourceList, error)
	ReconcileResource(context.Context) (*av1.SubResourceList, error)
	UninstallResource(context.Context) (*av1.SubResourceList, error)
}

// TrafficRolloutPhaseManager manages the TrafficRolloutPhase Phase of an OpenstackServiceLifeCycle
type TrafficRolloutPhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.SubResourceList, error)
	UpdateResource(context.Context) (*av1.SubResourceList, *av1.SubResourceList, error)
	ReconcileResource(context.Context) (*av1.SubResourceList, error)
	UninstallResource(context.Context) (*av1.SubResourceList, error)
}

// OperationalPhaseManager manages the OperationalPhase Phase of an OpenstackServiceLifeCycle
type OperationalPhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.SubResourceList, error)
	UpdateResource(context.Context) (*av1.SubResourceList, *av1.SubResourceList, error)
	ReconcileResource(context.Context) (*av1.SubResourceList, error)
	UninstallResource(context.Context) (*av1.SubResourceList, error)
}

// TrafficDrainPhaseManager manages the TrafficDrainPhase Phase of an OpenstackServiceLifeCycle
type TrafficDrainPhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.SubResourceList, error)
	UpdateResource(context.Context) (*av1.SubResourceList, *av1.SubResourceList, error)
	ReconcileResource(context.Context) (*av1.SubResourceList, error)
	UninstallResource(context.Context) (*av1.SubResourceList, error)
}

// UpgradePhaseManager manages the UpgradePhase Phase of an OpenstackServiceLifeCycle
type UpgradePhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.SubResourceList, error)
	UpdateResource(context.Context) (*av1.SubResourceList, *av1.SubResourceList, error)
	ReconcileResource(context.Context) (*av1.SubResourceList, error)
	UninstallResource(context.Context) (*av1.SubResourceList, error)
}

// RollbackPhaseManager manages the RollbackPhase Phase of an OpenstackServiceLifeCycle
type RollbackPhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.SubResourceList, error)
	UpdateResource(context.Context) (*av1.SubResourceList, *av1.SubResourceList, error)
	ReconcileResource(context.Context) (*av1.SubResourceList, error)
	UninstallResource(context.Context) (*av1.SubResourceList, error)
}

// DeletePhaseManager manages the DeletePhase Phase of an OpenstackServiceLifeCycle
type DeletePhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.SubResourceList, error)
	UpdateResource(context.Context) (*av1.SubResourceList, *av1.SubResourceList, error)
	ReconcileResource(context.Context) (*av1.SubResourceList, error)
	UninstallResource(context.Context) (*av1.SubResourceList, error)
}
