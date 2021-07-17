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

package basemanager_test

import (
	"fmt"
	"testing"

	. "github.com/kubedge/kubedge-operator-base/pkg/kubedgemanager"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yaml "sigs.k8s.io/yaml"
)

func TestBaseRenderer(t *testing.T) {
	refs := []metav1.OwnerReference{}
	suffix := ""
	renderFiles := []string{}
	renderValues := map[string]interface{}{}
	renderer := KubedgeBaseRenderer{
		Refs:         refs,
		Suffix:       suffix,
		RenderFiles:  renderFiles,
		RenderValues: renderValues,
	}
	rendered, err := renderer.RenderFile("foo", "bar", "testdata/classic.yaml")
	if err != nil {
		t.Fatalf(`error %v`, err)
	}

	if len(rendered.Items) == 0 {
		t.Fatalf(`No items rendered`)
	}

	for _, toCreate := range rendered.Items {
		blob, _ := yaml.Marshal(toCreate)
		_ = fmt.Sprintf("%s", string(blob))
	}
}
