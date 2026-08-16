// Copyright 2020 The Kubedge Authors
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

package conditions_test

import (
	"testing"

	av1 "github.com/kubedge/kubedge-operator-base/pkg/apis/kubedgeoperators/v1alpha1"
	"github.com/kubedge/kubedge-operator-base/pkg/conditions"
)

// condHolder is a minimal Setter (and therefore Getter) for exercising the conditions
// package. It embeds *av1.Arpscan purely to satisfy the client.Object part of the Getter
// interface; only GetConditions/SetConditions are exercised by the package under test.
type condHolder struct {
	*av1.Arpscan
	conds av1.KubedgeConditions
}

func (c *condHolder) GetConditions() av1.KubedgeConditions   { return c.conds }
func (c *condHolder) SetConditions(cs av1.KubedgeConditions) { c.conds = cs }

func newHolder(conds ...av1.KubedgeCondition) *condHolder {
	return &condHolder{Arpscan: &av1.Arpscan{}, conds: conds}
}

func TestGetAndHas(t *testing.T) {
	h := newHolder(av1.KubedgeCondition{Type: av1.ConditionDeployed, Status: av1.ConditionStatusTrue})

	if got := conditions.Get(h, av1.ConditionDeployed); got == nil {
		t.Fatalf("Get(Deployed) = nil, want the condition")
	}
	if got := conditions.Get(h, av1.ConditionFailed); got != nil {
		t.Fatalf("Get(Failed) = %+v, want nil for an absent condition", got)
	}
	if !conditions.Has(h, av1.ConditionDeployed) {
		t.Errorf("Has(Deployed) = false, want true")
	}
	if conditions.Has(h, av1.ConditionFailed) {
		t.Errorf("Has(Failed) = true, want false")
	}
}

func TestIsTrueFalseUnknown(t *testing.T) {
	h := newHolder(
		av1.KubedgeCondition{Type: av1.ConditionDeployed, Status: av1.ConditionStatusTrue},
		av1.KubedgeCondition{Type: av1.ConditionFailed, Status: av1.ConditionStatusFalse},
	)

	if !conditions.IsTrue(h, av1.ConditionDeployed) {
		t.Errorf("IsTrue(Deployed) = false, want true")
	}
	if conditions.IsFalse(h, av1.ConditionDeployed) {
		t.Errorf("IsFalse(Deployed) = true, want false")
	}
	if !conditions.IsFalse(h, av1.ConditionFailed) {
		t.Errorf("IsFalse(Failed) = false, want true")
	}
	if conditions.IsTrue(h, av1.ConditionFailed) {
		t.Errorf("IsTrue(Failed) = true, want false")
	}

	// Absent condition: IsTrue/IsFalse are false, IsUnknown defaults to true.
	if conditions.IsTrue(h, av1.ConditionRunning) || conditions.IsFalse(h, av1.ConditionRunning) {
		t.Errorf("absent condition should be neither True nor False")
	}
	if !conditions.IsUnknown(h, av1.ConditionRunning) {
		t.Errorf("IsUnknown(absent) = false, want true")
	}
}

func TestSetAddsThenUpdates(t *testing.T) {
	h := newHolder()

	conditions.Set(h, conditions.TrueCondition(av1.ConditionDeployed))
	if !conditions.IsTrue(h, av1.ConditionDeployed) {
		t.Fatalf("after Set(TrueCondition(Deployed)), IsTrue = false")
	}

	conditions.Set(h, conditions.FalseCondition(av1.ConditionDeployed, av1.ReasonInstallError, av1.ConditionSeverityError, "boom %d", 7))
	if !conditions.IsFalse(h, av1.ConditionDeployed) {
		t.Errorf("after Set(FalseCondition(Deployed)), IsFalse = false")
	}
	if got := conditions.GetMessage(h, av1.ConditionDeployed); got != "boom 7" {
		t.Errorf("GetMessage = %q, want %q", got, "boom 7")
	}
	if got := conditions.GetReason(h, av1.ConditionDeployed); got != av1.ReasonInstallError {
		t.Errorf("GetReason = %q, want %q", got, av1.ReasonInstallError)
	}
	// Updating an existing condition must not duplicate it.
	if n := len(h.GetConditions()); n != 1 {
		t.Errorf("condition count = %d, want 1 after update", n)
	}
}
