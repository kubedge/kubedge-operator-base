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
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"

	av1 "github.com/kubedge/kubedge-operator-base/pkg/apis/kubedgeoperators/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
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
