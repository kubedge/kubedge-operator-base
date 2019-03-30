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
	"io/ioutil"
	"sort"

	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/releaseutil"

	av1 "github.com/kubedge/kubedge-operator-base/pkg/apis/kubedgeoperators/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

var ()

type OwnerRefRenderer struct {
	refs         []metav1.OwnerReference
	suffix       string
	renderFiles  []string
	renderValues map[string]interface{}
}

// Adds the ownerrefs to all the documents in a YAML file
func (o *OwnerRefRenderer) RenderFile(name string, namespace string, fileName string) (*av1.SubResourceList, error) {

	yamlfmt, ferr := ioutil.ReadFile(fileName)
	if ferr != nil {
		log.Error(ferr, "Can not read file")
		return av1.NewSubResourceList(namespace, name), ferr
	}
	ownedRenderedFiles, err := o.fromYaml(name, namespace, string(yamlfmt))
	if err != nil {
		log.Info("Can not convert malformed yaml to unstructured", "filename", fileName)
		return ownedRenderedFiles, err
	}

	return ownedRenderedFiles, nil
}

func sortManifests(in map[string]string) []string {
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

// Reads a yaml file and converts into an Unstructured object
func (o *OwnerRefRenderer) fromYaml(name string, namespace string, filecontent string) (*av1.SubResourceList, error) {

	ownedRenderedFiles := av1.NewSubResourceList(namespace, name)

	manifests := releaseutil.SplitManifests(filecontent)
	for _, manifest := range sortManifests(manifests) {
		manifestMap := chartutil.FromYaml(manifest)

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
		u.SetOwnerReferences(o.refs)

		// Init name and namespace
		if u.GetName() == "" {
			u.SetName(name + "-" + o.suffix)
		}

		if u.GetNamespace() == "" {
			u.SetNamespace(namespace)
		}

		// Add OwnerReferences
		u.SetOwnerReferences(o.refs)

		// Add labels
		// labels := map[string]string{
		// 	"app": name,
		// }
		// u.SetLabels(labels)

		ownedRenderedFiles.Items = append(ownedRenderedFiles.Items, *u)
	}

	return ownedRenderedFiles, nil
}

// NewOwnerRefRenderer creates a new OwnerRef engine with a set of metav1.OwnerReferences to be added to assets
func NewOwnerRefRenderer(refs []metav1.OwnerReference, suffix string,
	renderFiles []string, renderValues map[string]interface{}) *OwnerRefRenderer {
	return &OwnerRefRenderer{
		refs:         refs,
		suffix:       suffix,
		renderFiles:  renderFiles,
		renderValues: renderValues,
	}
}
