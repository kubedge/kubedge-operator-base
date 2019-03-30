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

package basecontroller

import (
	"time"

	services "github.com/kubedge/kubedge-operator-base/pkg/services"

	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/client-go/tools/record"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// PhaseReconciler reconciles custom resources as Workflow, Jobs....
type PhaseReconciler struct {
	client                  client.Client
	scheme                  *runtime.Scheme
	recorder                record.EventRecorder
	managerFactory          services.PhaseManagerFactory
	reconcilePeriod         time.Duration
	depResourceWatchUpdater services.DependentResourceWatchUpdater
}

func (r *PhaseReconciler) contains(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}
