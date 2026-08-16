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
	"errors"
)

var (
	// ErrNotFound indicates the resource was not found.
	ErrNotFound = errors.New("resource not found")

	// ErrOwnershipMismatch indicates that one of the subresources does
	// not have the right ownership.
	ErrOwnershipMismatch = errors.New("ownership mismatch")

	// ErrSync detected during SyncResource
	ErrSync = errors.New("sync error")

	// ErrInstall detected during InstallResource
	ErrInstall = errors.New("install error")

	// ErrUninstall detected during UninstallResource
	ErrUninstall = errors.New("uninstall error")

	// ErrUpdate detected during UpdateResource
	ErrUpdate = errors.New("update error")

	// ErrReconcile detected during ReconcileResource
	ErrReconcile = errors.New("reconcile error")
)
