// Copyright 2019 The OpenstackLcm Authors
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

package v1alpha1

import (
	// "encoding/json"
	yaml "gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Administractive state of the reconcilation of a CRD by the corresponding controller
type OpenstackLcmAdminState string

// Describe the Administrative State of the Chart
const (
	// StateUnknown indicates that a release/chart/chartgroup/manifest automatic reconcilation by the controller is enabled
	StateEnabled OpenstackLcmAdminState = "enabled"
	// StateUnknown indicates that a release/chart/chartgroup/manifest automatic reconcilation by the controller is disabled
	StateDisabled OpenstackLcmAdminState = "disabled"
)

// String converts a OpenstackLcmAdminState to a printable string
func (x OpenstackLcmAdminState) String() string { return string(x) }

// LcmResourceState is the status of a release/chart/chartgroup/manifest
type LcmResourceState string

type LcmResourceConditionType string

// LcmResourceConditionStatus represents the current status of a Condition
type LcmResourceConditionStatus string

type LcmResourceConditionReason string

// String converts a LcmResourceState to a printable string
func (x LcmResourceState) String() string { return string(x) }

// String converts a LcmResourceConditionType to a printable string
func (x LcmResourceConditionType) String() string { return string(x) }

// String converts a LcmResourceConditionState to a printable string
func (x LcmResourceConditionStatus) String() string { return string(x) }

// String converts a LcmResourceConditionReason to a printable string
func (x LcmResourceConditionReason) String() string { return string(x) }

// Describe the status of a release
const (
	// StateUninitialied indicates that a release/chart/chartgroup/manifest exists, but has not been acted upon
	StateUninitialied LcmResourceState = "uninitialized"
	// StateUnknown indicates that a release/chart/chartgroup/manifest is in an uncertain state.
	StateUnknown LcmResourceState = "unknown"
	// StateInitialized indicates that a release/chart/chartgroup/manifest is in an Kubernetes
	StateInitialized LcmResourceState = "initialized"
	// StateDeployed indicates that the release/chart/chartgroup/manifest has been downloaded from artifact repository
	StateDownloaded LcmResourceState = "downloaded"
	// StateDeployed indicates that the release/chart/chartgroup/manifest has been pushed to Kubernetes.
	StateDeployed LcmResourceState = "deployed"
	// StateUninstalled indicates that a release/chart/chartgroup/manifest has been uninstalled from Kubermetes.
	StateUninstalled LcmResourceState = "uninstalled"
	// StateSuperseded indicates that this release/chart/chartgroup/manifest object is outdated and a newer one exists.
	StateSuperseded LcmResourceState = "superseded"
	// StateFailed indicates that the release/chart/chartgroup/manifest was not successfully deployed.
	StateFailed LcmResourceState = "failed"
	// StateUninstalling indicates that a uninstall operation is underway.
	StateUninstalling LcmResourceState = "uninstalling"
	// StatePendingInstall indicates that an install operation is underway.
	StatePendingInstall LcmResourceState = "pending-install"
	// StatePendingUpgrade indicates that an upgrade operation is underway.
	StatePendingUpgrade LcmResourceState = "pending-upgrade"
	// StatePendingRollback indicates that an rollback operation is underway.
	StatePendingRollback LcmResourceState = "pending-rollback"
	// StatePendingBackup indicates that an data backup operation is underway.
	StatePendingBackup LcmResourceState = "pending-backup"
	// StatePendingRestore indicates that an data restore operation is underway.
	StatePendingRestore LcmResourceState = "pending-restore"
	// StatePendingInitialization indicates that an data initialization operation is underway.
	StatePendingInitialization LcmResourceState = "pending-initialization"
)

// These represent acceptable values for a LcmResourceConditionStatus
const (
	ConditionStatusTrue    LcmResourceConditionStatus = "True"
	ConditionStatusFalse                              = "False"
	ConditionStatusUnknown                            = "Unknown"
)

// These represent acceptable values for a LcmResourceConditionType
const (
	ConditionIrreconcilable LcmResourceConditionType = "Irreconcilable"
	ConditionFailed                                  = "Failed"
	ConditionInitialized                             = "Initializing"
	ConditionEnabled                                 = "Enabled"
	ConditionDownloaded                              = "Downloaded"
	ConditionDeployed                                = "Deployed"

	// JEB: Not sure we will ever be able to use those conditions
	ConditionBackedUp   LcmResourceConditionType = "BackedUp"
	ConditionRestored                            = "Restored"
	ConditionUpgraded                            = "Upgraded"
	ConditionRolledBack                          = "RolledBack"
)

// The following represent the more fine-grained reasons for a given condition
const (
	// Successful Conditions Reasons
	ReasonInstallSuccessful   LcmResourceConditionReason = "InstallSuccessful"
	ReasonDownloadSuccessful                             = "DownloadSuccessful"
	ReasonReconcileSuccessful                            = "ReconcileSuccessful"
	ReasonUninstallSuccessful                            = "UninstallSuccessful"
	ReasonUpdateSuccessful                               = "UpdateSuccessful"
	ReasonBackupSuccessful                               = "BackupSuccessful"
	ReasonRestoreSuccessful                              = "RestoreSuccessful"
	ReasonUpgradeSuccessful                              = "UpgradeSuccessful"
	ReasonRollbackSuccessful                             = "RollbackSuccessful"

	// Error Condition Reasons
	ReasonInstallError   LcmResourceConditionReason = "InstallError"
	ReasonDownloadError                             = "DownloadError"
	ReasonReconcileError                            = "ReconcileError"
	ReasonUninstallError                            = "UninstallError"
	ReasonUpdateError                               = "UpdateError"
	ReasonBackupError                               = "BackupError"
	ReasonRestoreError                              = "RestoreError"
	ReasonUpgradeError                              = "UpgradeError"
	ReasonRollbackError                             = "RollbackError"
)

// LcmResourceCondition represents one current condition of an Lcm resource
// A condition might not show up if it is not happening.
// For example, if a chart is not deploying, the Deploying condition would not show up.
// If a chart is deploying and encountered a problem that prevents the deployment,
// the Deploying condition's status will would be False and communicate the problem back.
type LcmResourceCondition struct {
	Type               LcmResourceConditionType   `json:"type"`
	Status             LcmResourceConditionStatus `json:"status"`
	Reason             LcmResourceConditionReason `json:"reason,omitempty"`
	Message            string                     `json:"message,omitempty"`
	ResourceName       string                     `json:"resourceName,omitempty"`
	ResourceVersion    int32                      `json:"resourceVersion,omitempty"`
	LastTransitionTime metav1.Time                `json:"lastTransitionTime,omitempty"`
}

type LcmResourceConditionListHelper struct {
	Items []LcmResourceCondition `json:"items"`
}

// OpenstackLcmStatus represents the common attributes shared amongst armada resources
type OpenstackLcmStatus struct {
	// Succeeded indicates if the release's ActualState satisfies its target state
	Succeeded bool `json:"satisfied"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"reason,omitempty"`
	// Actual state of the Lcm Custom Resources
	ActualState LcmResourceState `json:"actual_state"`
	// List of conditions and states related to the resource. JEB: Feature kind of overlap with event recorder
	Conditions []LcmResourceCondition `json:"conditions,omitempty"`
}

// PhaseStatus represents the common attributes shared amongst armada resources
type PhaseStatus struct {
	OpenstackLcmStatus `json:",inline"`

	// OpenstackVersion is the version of the backup openstack server.
	ActualOpenstackServiceVersion string `json:"actualOpenstackServiceVersion,omitempty"`
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *OpenstackLcmStatus) SetCondition(cond LcmResourceCondition, tgt LcmResourceState) {

	// Add the condition to the list
	chelper := LcmResourceConditionListHelper{Items: s.Conditions}
	s.Conditions = chelper.SetCondition(cond)

	// Recompute the state
	s.ComputeActualState(cond, tgt)
}

// RemoveCondition removes the condition with the passed condition type from
// the status object. If the condition is not already present, the returned
// status object is returned unchanged. RemoveCondition does not update the
// resource in the cluster.
func (s *OpenstackLcmStatus) RemoveCondition(conditionType LcmResourceConditionType) {
	for i, cond := range s.Conditions {
		if cond.Type == conditionType {
			s.Conditions = append(s.Conditions[:i], s.Conditions[i+1:]...)
			return
		}
	}
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *LcmResourceConditionListHelper) SetCondition(condition LcmResourceCondition) []LcmResourceCondition {

	// Initialize the Items array if needed
	if s.Items == nil {
		s.Items = make([]LcmResourceCondition, 0)
	}

	now := metav1.Now()
	for i := range s.Items {
		if s.Items[i].Type == condition.Type {
			if s.Items[i].Status != condition.Status {
				condition.LastTransitionTime = now
			} else {
				condition.LastTransitionTime = s.Items[i].LastTransitionTime
			}
			s.Items[i] = condition
			return s.Items
		}
	}

	// If the condition does not exist,
	// initialize the lastTransitionTime
	condition.LastTransitionTime = now
	s.Items = append(s.Items, condition)
	return s.Items
}

// RemoveCondition removes the condition with the passed condition type from
// the status object. If the condition is not already present, the returned
// status object is returned unchanged. RemoveCondition does not update the
// resource in the cluster.
func (s *LcmResourceConditionListHelper) RemoveCondition(conditionType LcmResourceConditionType) []LcmResourceCondition {

	// Initialize the Items array if needed
	if s.Items == nil {
		s.Items = make([]LcmResourceCondition, 0)
	}

	for i := range s.Items {
		if s.Items[i].Type == conditionType {
			s.Items = append(s.Items[:i], s.Items[i+1:]...)
			return s.Items
		}
	}
	return s.Items
}

// Initialize the LcmResourceCondition list
func (s *LcmResourceConditionListHelper) InitIfEmpty() []LcmResourceCondition {

	// Initialize the Items array if needed
	if s.Items == nil {
		s.Items = make([]LcmResourceCondition, 0)
	}

	return s.Items
}

// Utility function to print an LcmResourceCondition list
func (s *LcmResourceConditionListHelper) PrettyPrint() string {
	// res, _ := json.MarshalIndent(s.Items, "", "\t")
	res, _ := yaml.Marshal(s.Items)
	return string(res)
}

// Utility function to find an LcmResourceCondition within the List
func (s *LcmResourceConditionListHelper) FindCondition(conditionType LcmResourceConditionType, conditionStatus LcmResourceConditionStatus) *LcmResourceCondition {
	var found *LcmResourceCondition
	for _, condition := range s.Items {
		if condition.Type == conditionType && condition.Status == conditionStatus {
			found = &condition
			break
		}
	}
	return found
}

func (s *OpenstackLcmStatus) ComputeActualState(cond LcmResourceCondition, target LcmResourceState) {
	// TODO(Ian): finish this
	if cond.Status == ConditionStatusTrue {
		if cond.Type == ConditionInitialized {
			// Since that condition is set almost systematically
			// let's do not recompute the state.
			if (s.ActualState == "") || (s.ActualState == StateUnknown) {
				s.ActualState = StateInitialized
				s.Succeeded = (s.ActualState == target)
				s.Reason = ""
			}
		} else if cond.Type == ConditionDeployed {
			s.ActualState = StateDeployed
			s.Succeeded = (s.ActualState == target)
			s.Reason = ""
		} else if cond.Type == ConditionEnabled {
			if (s.ActualState == "") || (s.ActualState == StateUnknown) {
				s.ActualState = StatePendingInitialization
				s.Succeeded = (s.ActualState == target)
				s.Reason = ""
			}
		} else if cond.Type == ConditionIrreconcilable {
			s.ActualState = StateFailed
			s.Succeeded = false
			s.Reason = cond.Reason.String()
		} else {
			s.Succeeded = (s.ActualState == target)
			s.Reason = ""
		}
	} else {
		if cond.Type == ConditionDeployed {
			s.ActualState = StateUninstalled
			s.Succeeded = (s.ActualState == target)
			s.Reason = ""
		} else if cond.Type == ConditionEnabled {
			s.ActualState = StateUnknown
			s.Succeeded = true
			s.Reason = "Disabled Resource is always successful"
		} else {
			s.Succeeded = (s.ActualState == target)
			s.Reason = ""
		}
	}
}

// PhaseSource describe the location of the CR to create during a Phase of an
// Openstack Service Life Cycle.
type PhaseSource struct {
	// (optional) authentication method
	AuthMethod string `json:"authMethod,omitempty"`
	// ``url`` or ``path`` to the chart's parent directory
	Location string `json:"location"`
	// (optional) http proxy server
	ProxyServer string `json:"proxyServer,omitempty"`
	// (optional) branch, commit, or reference in the repo (``master`` if not specified)
	Reference string `json:"reference,omitempty"`
	// (optional) relative path to target chart from parent (``.`` if not specified)
	Subpath string `json:"subpath"`
	// source to build the chart: ``git``, ``local``, or ``tar``
	Type string `json:"type"`
}

// PhaseSpec defines the desired state of Phase
type PhaseSpec struct {
	// provide a path to a ``git repo``, ``local dir``, or ``tarball url`` chart
	Source *PhaseSource `json:"source"`
	// Openstack Service Name
	OpenstackServiceName string `json:"openstackServiceName"`
	// Openstack Service EndPoint
	OpenstackServiceEndPoint string `json:"openstackServiceEndPoint,omitempty"`

	// OpenstackServiceVersion is the version of the openstack service.
	TargetOpenstackServiceVersion string `json:"targetOpenstackServiceVersion,omitempty"`
	// Target state of the Lcm Custom Resources
	TargetState LcmResourceState `json:"targetState"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerRevision implements an immutable snapshot of state data. Clients
// are responsible for serializing and deserializing the objects that contain
// their internal state.
// Once a ControllerRevision has been successfully created, it can not be updated.
// The API Server will fail validation of all requests that attempt to mutate
// the Data field. ControllerRevisions may, however, be deleted. Note that, due to its use by both
// the DaemonSet and StatefulSet controllers for update and rollback, this object is beta. However,
// it may be subject to name and representation changes in future releases, and clients should not
// depend on its stability. It is primarily for internal use by controllers.
type ControllerRevision struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Data is the serialized representation of the state.
	Data runtime.RawExtension `json:"data,omitempty" protobuf:"bytes,2,opt,name=data"`

	// Revision indicates the revision of the state represented by Data.
	Revision int64 `json:"revision" protobuf:"varint,3,opt,name=revision"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerRevisionList is a resource containing a list of ControllerRevision objects.
type ControllerRevisionList struct {
	metav1.TypeMeta `json:",inline"`

	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of ControllerRevisions
	Items []ControllerRevision `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// SubResourceList represent the list of
type SubResourceList struct {
	Name      string
	Namespace string
	Notes     string
	Version   int32

	// Items is the list of Resources deployed in the K8s cluster
	Items [](unstructured.Unstructured) `json:"items"`
}

// Returns the Name for the SubResourceList
func (obj *SubResourceList) GetName() string {
	return obj.Name
}

// Returns the Namespace for this SubResourceList
func (obj *SubResourceList) GetNamespace() string {
	return obj.Namespace
}

// Returns the Notes for this SubResourceList
func (obj *SubResourceList) GetNotes() string {
	return obj.Notes
}

// Returns the Version for this SubResourceList
func (obj *SubResourceList) GetVersion() int32 {
	return obj.Version
}

// Returns the DependentResource for this SubResourceList
func (obj *SubResourceList) GetDependentResources() []unstructured.Unstructured {
	return obj.Items
}

// Returns a new SubResourceList
func NewSubResourceList(namespace string, name string) *SubResourceList {
	res := &SubResourceList{Namespace: namespace, Name: name}
	res.Items = make([]unstructured.Unstructured, 0)
	return res
}

// PhaseList represent the list of Phase for that Service
type PhaseList struct {
	Name      string
	Namespace string
	Notes     string
	Version   int32

	// Items is the list of Resources deployed in the K8s cluster
	Items [](unstructured.Unstructured) `json:"items"`
}

// Returns the Name for the PhaseList
func (obj *PhaseList) GetName() string {
	return obj.Name
}

// Returns the Namespace for this PhaseList
func (obj *PhaseList) GetNamespace() string {
	return obj.Namespace
}

// Returns the Notes for this PhaseList
func (obj *PhaseList) GetNotes() string {
	return obj.Notes
}

// Returns the Version for this PhaseList
func (obj *PhaseList) GetVersion() int32 {
	return obj.Version
}

// Returns the DependentResource for this PhaseList
func (obj *PhaseList) GetDependentResources() []unstructured.Unstructured {
	return obj.Items
}

// Returns a new PhaseList
func NewPhaseList(namespace string, name string) *PhaseList {
	res := &PhaseList{Namespace: namespace, Name: name}
	res.Items = make([]unstructured.Unstructured, 0)
	return res
}
