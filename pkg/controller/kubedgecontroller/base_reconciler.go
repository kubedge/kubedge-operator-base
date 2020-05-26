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
	"reflect"
	"time"

	av1 "github.com/kubedge/kubedge-operator-base/pkg/apis/kubedgeoperators/v1alpha1"
	mgr "github.com/kubedge/kubedge-operator-base/pkg/kubedgemanager"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	crtpredicate "sigs.k8s.io/controller-runtime/pkg/predicate"
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

// buildDependentPredicate create the predicates used by subresources watches
func (r *KubedgeBaseReconciler) BuildDependentPredicate() *crtpredicate.Funcs {

	dependentPredicate := crtpredicate.Funcs{
		// We don't need to reconcile dependent resource creation events
		// because dependent resources are only ever created during
		// reconciliation. Another reconcile would be redundant.
		CreateFunc: func(e event.CreateEvent) bool {
			// o := e.Object.(*unstructured.Unstructured)
			// log.Info("CreateEvent. Filtering", "resource", o.GetName(), "namespace", o.GetNamespace(),
			//	"apiVersion", o.GroupVersionKind().GroupVersion(), "kind", o.GroupVersionKind().Kind)
			return false
		},

		// Reconcile when a dependent resource is deleted so that it can be
		// recreated.
		DeleteFunc: func(e event.DeleteEvent) bool {
			// o := e.Object.(*unstructured.Unstructured)
			// log.Info("DeleteEvent. Triggering", "resource", o.GetName(), "namespace", o.GetNamespace(),
			//	"apiVersion", o.GroupVersionKind().GroupVersion(), "kind", o.GroupVersionKind().Kind)
			return true
		},

		// Reconcile when a dependent resource is updated, so that it can
		// be patched back to the resource managed by the Argo workflow, if
		// necessary. Ignore updates that only change the status and
		// resourceVersion.
		UpdateFunc: func(e event.UpdateEvent) bool {
			u := e.ObjectOld.(*unstructured.Unstructured)
			v := e.ObjectNew.(*unstructured.Unstructured)

			// TODO(jeb): Note sure if we really want to do that
			// Filter on Kind Change
			if u.GetKind() == "ConfigMap" || u.GetKind() == "Secret" {
				return false
			}

			// Filter on Status change
			dep := &av1.KubernetesDependency{}
			changed, oldv, newv := dep.UnstructuredStatusChanged(u, v)
			if changed {
				log.Info("UpdateEvent. Status changed", "resource", u.GetName(), "namespace", u.GetNamespace(),
					"apiVersion", u.GroupVersionKind().GroupVersion(), "kind", u.GroupVersionKind().Kind,
					"old", oldv, "new", newv)
				return true
			}

			// Filter on Spec change
			old := u.DeepCopy()
			new := v.DeepCopy()

			delete(old.Object, "status")
			delete(new.Object, "status")
			old.SetResourceVersion("")
			new.SetResourceVersion("")

			if reflect.DeepEqual(old.Object, new.Object) {
				// log.Info("UpdateEvent. Spec unchanged", "resource", new.GetName(), "namespace", new.GetNamespace(),
				//	"apiVersion", new.GroupVersionKind().GroupVersion(), "kind", new.GroupVersionKind().Kind)
				return false
			} else {
				log.Info("UpdateEvent. Spec changed", "resource", new.GetName(), "namespace", new.GetNamespace(),
					"apiVersion", new.GroupVersionKind().GroupVersion(), "kind", new.GroupVersionKind().Kind)
				return true
			}
		},
	}

	return &dependentPredicate
}
