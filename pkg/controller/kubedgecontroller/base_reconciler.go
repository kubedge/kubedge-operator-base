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

package kubedgecontroller

import (
	"time"

	mgr "github.com/kubedge/kubedge-operator-base/pkg/kubedgemanager"

	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/client-go/tools/record"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// KubedgeBaseReconciler reconciles custom resources as Workflow, Jobs....
type KubedgeBaseReconciler struct {
	Client                  client.Client
	Scheme                  *runtime.Scheme
	Recorder                record.EventRecorder
	ManagerFactory          mgr.KubedgeResourceManagerFactory
	ReconcilePeriod         time.Duration
	DepResourceWatchUpdater mgr.DependentResourceWatchUpdater
}

func (r *KubedgeBaseReconciler) Contains(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}
