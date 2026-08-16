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

package v1alpha1

import "testing"

func TestComputeActualState(t *testing.T) {
	cases := []struct {
		name      string
		start     KubedgeResourceState
		cond      KubedgeCondition
		target    KubedgeResourceState
		wantState KubedgeResourceState
		wantSat   bool
	}{
		{
			name:      "pending true",
			cond:      KubedgeCondition{Type: ConditionPending, Status: ConditionStatusTrue},
			target:    StatePending,
			wantState: StatePending,
			wantSat:   true,
		},
		{
			name:      "running true is never satisfied",
			cond:      KubedgeCondition{Type: ConditionRunning, Status: ConditionStatusTrue},
			target:    StateRunning,
			wantState: StateRunning,
			wantSat:   false,
		},
		{
			name:      "deployed true matches target",
			cond:      KubedgeCondition{Type: ConditionDeployed, Status: ConditionStatusTrue},
			target:    StateDeployed,
			wantState: StateDeployed,
			wantSat:   true,
		},
		{
			name:      "failed true carries reason and is not satisfied",
			cond:      KubedgeCondition{Type: ConditionFailed, Status: ConditionStatusTrue, Reason: ReasonInstallError},
			target:    StateDeployed,
			wantState: StateFailed,
			wantSat:   false,
		},
		{
			name:      "error true",
			cond:      KubedgeCondition{Type: ConditionError, Status: ConditionStatusTrue, Reason: ReasonReconcileError},
			target:    StateDeployed,
			wantState: StateError,
			wantSat:   false,
		},
		{
			name:      "initialized true from empty state",
			cond:      KubedgeCondition{Type: ConditionInitialized, Status: ConditionStatusTrue},
			target:    StateInitialized,
			wantState: StateInitialized,
			wantSat:   true,
		},
		{
			name:      "deployed false means uninstalled",
			cond:      KubedgeCondition{Type: ConditionDeployed, Status: ConditionStatusFalse},
			target:    StateUninstalled,
			wantState: StateUninstalled,
			wantSat:   true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := &KubedgeStatus{ActualState: tc.start}
			s.ComputeActualState(tc.cond, tc.target)
			if s.ActualState != tc.wantState {
				t.Errorf("ActualState = %q, want %q", s.ActualState, tc.wantState)
			}
			if s.Satisfied != tc.wantSat {
				t.Errorf("Satisfied = %v, want %v", s.Satisfied, tc.wantSat)
			}
		})
	}
}

func TestComputeActualStateFailedReason(t *testing.T) {
	s := &KubedgeStatus{}
	s.ComputeActualState(KubedgeCondition{Type: ConditionFailed, Status: ConditionStatusTrue, Reason: ReasonInstallError}, StateDeployed)
	if s.Reason != ReasonInstallError.String() {
		t.Errorf("Reason = %q, want %q", s.Reason, ReasonInstallError.String())
	}
}

func TestSetConditionAndRemove(t *testing.T) {
	s := &KubedgeStatus{}

	s.SetCondition(KubedgeCondition{Type: ConditionDeployed, Status: ConditionStatusTrue}, StateDeployed)
	if len(s.Conditions) != 1 {
		t.Fatalf("condition count = %d, want 1", len(s.Conditions))
	}
	if s.ActualState != StateDeployed || !s.Satisfied {
		t.Errorf("after SetCondition(Deployed/True, target=Deployed): state=%q satisfied=%v, want deployed/true", s.ActualState, s.Satisfied)
	}

	// Re-setting the same type replaces rather than appends.
	s.SetCondition(KubedgeCondition{Type: ConditionDeployed, Status: ConditionStatusFalse}, StateUninstalled)
	if len(s.Conditions) != 1 {
		t.Errorf("condition count = %d, want 1 after replace", len(s.Conditions))
	}

	s.RemoveCondition(ConditionDeployed)
	if len(s.Conditions) != 0 {
		t.Errorf("condition count = %d, want 0 after remove", len(s.Conditions))
	}
	// Removing an absent condition is a no-op.
	s.RemoveCondition(ConditionDeployed)
}

func TestConditionListHelperFindCondition(t *testing.T) {
	h := &KubedgeConditionListHelper{Items: []KubedgeCondition{
		{Type: ConditionDeployed, Status: ConditionStatusTrue},
		{Type: ConditionFailed, Status: ConditionStatusFalse},
	}}

	if got := h.FindCondition(ConditionDeployed, ConditionStatusTrue); got == nil {
		t.Errorf("FindCondition(Deployed, True) = nil, want a match")
	}
	if got := h.FindCondition(ConditionDeployed, ConditionStatusFalse); got != nil {
		t.Errorf("FindCondition(Deployed, False) = %+v, want nil (status mismatch)", got)
	}
	if got := h.FindCondition(ConditionRunning, ConditionStatusTrue); got != nil {
		t.Errorf("FindCondition(Running, True) = %+v, want nil (absent)", got)
	}
}
