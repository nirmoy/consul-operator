/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/nirmoy/consul-operator/pkg/apis/consuloperator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConsulClusters implements ConsulClusterInterface
type FakeConsulClusters struct {
	Fake *FakeConsuloperatorV1alpha1
	ns   string
}

var consulclustersResource = schema.GroupVersionResource{Group: "consuloperator.k8s.io", Version: "v1alpha1", Resource: "consulclusters"}

var consulclustersKind = schema.GroupVersionKind{Group: "consuloperator.k8s.io", Version: "v1alpha1", Kind: "ConsulCluster"}

// Get takes name of the consulCluster, and returns the corresponding consulCluster object, and an error if there is any.
func (c *FakeConsulClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.ConsulCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(consulclustersResource, c.ns, name), &v1alpha1.ConsulCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsulCluster), err
}

// List takes label and field selectors, and returns the list of ConsulClusters that match those selectors.
func (c *FakeConsulClusters) List(opts v1.ListOptions) (result *v1alpha1.ConsulClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(consulclustersResource, consulclustersKind, c.ns, opts), &v1alpha1.ConsulClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConsulClusterList{ListMeta: obj.(*v1alpha1.ConsulClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConsulClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested consulClusters.
func (c *FakeConsulClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(consulclustersResource, c.ns, opts))

}

// Create takes the representation of a consulCluster and creates it.  Returns the server's representation of the consulCluster, and an error, if there is any.
func (c *FakeConsulClusters) Create(consulCluster *v1alpha1.ConsulCluster) (result *v1alpha1.ConsulCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(consulclustersResource, c.ns, consulCluster), &v1alpha1.ConsulCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsulCluster), err
}

// Update takes the representation of a consulCluster and updates it. Returns the server's representation of the consulCluster, and an error, if there is any.
func (c *FakeConsulClusters) Update(consulCluster *v1alpha1.ConsulCluster) (result *v1alpha1.ConsulCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(consulclustersResource, c.ns, consulCluster), &v1alpha1.ConsulCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsulCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConsulClusters) UpdateStatus(consulCluster *v1alpha1.ConsulCluster) (*v1alpha1.ConsulCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(consulclustersResource, "status", c.ns, consulCluster), &v1alpha1.ConsulCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsulCluster), err
}

// Delete takes name of the consulCluster and deletes it. Returns an error if one occurs.
func (c *FakeConsulClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(consulclustersResource, c.ns, name), &v1alpha1.ConsulCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConsulClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(consulclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConsulClusterList{})
	return err
}

// Patch applies the patch and returns the patched consulCluster.
func (c *FakeConsulClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConsulCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(consulclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConsulCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsulCluster), err
}
