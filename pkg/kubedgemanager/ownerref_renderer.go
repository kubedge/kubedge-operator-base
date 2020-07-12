// Copyright The Helm Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build v1

package basemanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"

	av1 "github.com/kubedge/kubedge-operator-base/pkg/apis/kubedgeoperators/v1alpha1"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	yaml "sigs.k8s.io/yaml"
)

var ()

// KubedgeResourceRenderer
type KubedgeResourceRenderer interface {
	RenderFile(name string, namespace string, fileName string) (*av1.SubResourceList, error)
}

type KubedgeBaseRenderer struct {
	Refs         []metav1.OwnerReference
	Suffix       string
	RenderFiles  []string
	RenderValues map[string]interface{}
}

// Adds the ownerrefs to all the documents in a YAML file
func (o *KubedgeBaseRenderer) RenderFile(name string, namespace string, fileName string) (*av1.SubResourceList, error) {

	yamlfmt, ferr := ioutil.ReadFile(fileName)
	if ferr != nil {
		log.Error(ferr, "Can not read file")
		return av1.NewSubResourceList(namespace, name), ferr
	}
	ownedRenderedFiles, err := o.FromYaml(name, namespace, string(yamlfmt))
	if err != nil {
		log.Info("Can not convert malformed yaml to unstructured", "filename", fileName)
		return ownedRenderedFiles, err
	}

	return ownedRenderedFiles, nil
}

func SortManifests(in map[string]string) []string {
	var keys []string
	for k := range in {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var manifests []string
	for _, k := range keys {
		manifests = append(manifests, in[k])
	}
	return manifests
}

var sep = regexp.MustCompile("(?:^|\\s*\n)---\\s*")

// SplitManifests takes a string of manifest and returns a map contains individual manifests
func (o *KubedgeBaseRenderer) SplitManifests(bigFile string) map[string]string {
	// Basically, we're quickly splitting a stream of YAML documents into an
	// array of YAML docs. In the current implementation, the file name is just
	// a place holder, and doesn't have any further meaning.
	tpl := "manifest-%d"
	res := map[string]string{}
	// Making sure that any extra whitespace in YAML stream doesn't interfere in splitting documents correctly.
	bigFileTmp := strings.TrimSpace(bigFile)
	docs := sep.Split(bigFileTmp, -1)
	var count int
	for _, d := range docs {

		if d == "" {
			continue
		}

		d = strings.TrimSpace(d)
		res[fmt.Sprintf(tpl, count)] = d
		count = count + 1
	}
	return res
}

// FromYaml converts a YAML document into a map[string]interface{}.
func (o *KubedgeBaseRenderer) Unmarshal(str string) map[string]interface{} {
	m := map[string]interface{}{}

	if err := yaml.Unmarshal([]byte(str), &m); err != nil {
		m["Error"] = err.Error()
	}
	return m
}

// Reads a yaml file and converts into an Unstructured object
func (o *KubedgeBaseRenderer) FromYaml(name string, namespace string, filecontent string) (*av1.SubResourceList, error) {

	ownedRenderedFiles := av1.NewSubResourceList(namespace, name)

	manifests := o.SplitManifests(filecontent)
	for _, manifest := range SortManifests(manifests) {
		manifestMap := o.Unmarshal(manifest)

		if _, ok := manifestMap["Error"]; ok {
			log.Error(nil, "error parsing rendered template to add ownerrefs")
			continue
		}

		// Check if the document is empty
		if len(manifestMap) == 0 {
			continue
		}

		unst, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&manifestMap)
		if err != nil {
			log.Error(err, "error converting to Unstructured")
			continue
		}

		u := &unstructured.Unstructured{Object: unst}
		u.SetOwnerReferences(o.Refs)

		// Init name and namespace
		if u.GetName() == "" {
			u.SetName(name + "-" + o.Suffix)
		}

		if u.GetNamespace() == "" {
			u.SetNamespace(namespace)
		}

		// Add OwnerReferences
		u.SetOwnerReferences(o.Refs)

		// Add labels
		// labels := map[string]string{
		// 	"app": name,
		// }
		// u.SetLabels(labels)

		ownedRenderedFiles.Items = append(ownedRenderedFiles.Items, *u)
	}

	return ownedRenderedFiles, nil
}

// MergePodTemplate takes a possibly nil container template and a
// list of steps, merging each of the steps with the container template, if
// it's not nil, and returning the resulting list.
// Deprecated
func (o *KubedgeBaseRenderer) MergePodTemplateSpec(template *corev1.PodTemplateSpec, k *av1.KubedgeSetSpec) error {
	if template == nil || k == nil {
		return nil
	}

	// We need JSON bytes to generate a patch to merge the step containers
	// onto the template container, so marshal the template.
	templateAsJSON, err := json.Marshal(template)
	if err != nil {
		return err
	}
	// We need to do a three-way merge to actually merge the template and
	// step containers, so we need an empty container as the "original"
	emptyAsJSON, err := json.Marshal(&corev1.PodTemplateSpec{})
	if err != nil {
		return err
	}

	// Marshal the step's to JSON
	stepAsJSON, err := json.Marshal(k.Template)
	if err != nil {
		return err
	}

	// Get the patch meta for Container, which is needed for generating and applying the merge patch.
	patchSchema, err := strategicpatch.NewPatchMetaFromStruct(template)

	if err != nil {
		return err
	}

	// Create a merge patch, with the empty JSON as the original, the step JSON as the modified, and the template
	// JSON as the current - this lets us do a deep merge of the template and step containers, with awareness of
	// the "patchMerge" tags.
	patch, err := strategicpatch.CreateThreeWayMergePatch(emptyAsJSON, stepAsJSON, templateAsJSON, patchSchema, true)
	if err != nil {
		return err
	}

	// Actually apply the merge patch to the template JSON.
	mergedAsJSON, err := strategicpatch.StrategicMergePatchUsingLookupPatchMeta(templateAsJSON, patch, patchSchema)
	if err != nil {
		return err
	}

	// Unmarshal the merged JSON to a Container pointer, and return it.
	err = json.Unmarshal(mergedAsJSON, template)
	if err != nil {
		return err
	}

	// If the container's args is nil, reset it to empty instead
	// if merged.Args == nil && s.Args != nil {
	//	merged.Args = []string{}
	// }

	return nil
}

// Update the Unstructured read in the file using the content of the Spec.
func (o *KubedgeBaseRenderer) UpdateStatefulSet(u *unstructured.Unstructured, k *av1.KubedgeSetSpec) {
	if k != nil {

		out := v1.StatefulSet{}
		err1 := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &out)
		if err1 != nil {
			log.Error(err1, "error converting from Unstructured")
		}

		if k.Replicas != nil {
			out.Spec.Replicas = k.Replicas
		}

		if k.Selector != nil {
			out.Spec.Selector = k.Selector.DeepCopy()
		}

		// out.Spec.Template = *(k.Template.DeepCopy())
		err2 := o.MergePodTemplateSpec(&out.Spec.Template, k)
		if err2 != nil {
			log.Error(err2, "error merging PodTemplateSpec")
		}

		unst, err3 := runtime.DefaultUnstructuredConverter.ToUnstructured(&out)
		if err3 != nil {
			log.Error(err3, "error converting to Unstructured")
		}

		u.SetUnstructuredContent(unst)
	}
}

// NewKubedgeBaseRenderer creates a new OwnerRef engine with a set of metav1.OwnerReferences to be added to assets
func NewKubedgeBaseRenderer(refs []metav1.OwnerReference, suffix string,
	renderFiles []string, renderValues map[string]interface{}) KubedgeResourceRenderer {
	return &KubedgeBaseRenderer{
		Refs:         refs,
		Suffix:       suffix,
		RenderFiles:  renderFiles,
		RenderValues: renderValues,
	}
}
