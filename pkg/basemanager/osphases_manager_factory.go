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
	av1 "github.com/kubedge/kubedge-operator-base/pkg/apis/kubedgeoperators/v1alpha1"
)

// ManagerFactory creates Managers that are specific to custom resources.
type KubedgeResourceManagerFactory interface {
	NewArpscanManager(r *av1.Arpscan) KubedgeResourceManager
	NewECDSClusterManager(r *av1.ECDSCluster) KubedgeResourceManager
	NewMMESimManager(r *av1.MMESim) KubedgeResourceManager
}
