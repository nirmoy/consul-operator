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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/nirmoy/consul-operator/pkg/apis/consuloperator/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConsulClusterLister helps list ConsulClusters.
type ConsulClusterLister interface {
	// List lists all ConsulClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ConsulCluster, err error)
	// ConsulClusters returns an object that can list and get ConsulClusters.
	ConsulClusters(namespace string) ConsulClusterNamespaceLister
	ConsulClusterListerExpansion
}

// consulClusterLister implements the ConsulClusterLister interface.
type consulClusterLister struct {
	indexer cache.Indexer
}

// NewConsulClusterLister returns a new ConsulClusterLister.
func NewConsulClusterLister(indexer cache.Indexer) ConsulClusterLister {
	return &consulClusterLister{indexer: indexer}
}

// List lists all ConsulClusters in the indexer.
func (s *consulClusterLister) List(selector labels.Selector) (ret []*v1alpha1.ConsulCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConsulCluster))
	})
	return ret, err
}

// ConsulClusters returns an object that can list and get ConsulClusters.
func (s *consulClusterLister) ConsulClusters(namespace string) ConsulClusterNamespaceLister {
	return consulClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConsulClusterNamespaceLister helps list and get ConsulClusters.
type ConsulClusterNamespaceLister interface {
	// List lists all ConsulClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ConsulCluster, err error)
	// Get retrieves the ConsulCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ConsulCluster, error)
	ConsulClusterNamespaceListerExpansion
}

// consulClusterNamespaceLister implements the ConsulClusterNamespaceLister
// interface.
type consulClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConsulClusters in the indexer for a given namespace.
func (s consulClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConsulCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConsulCluster))
	})
	return ret, err
}

// Get retrieves the ConsulCluster from the indexer for a given namespace and name.
func (s consulClusterNamespaceLister) Get(name string) (*v1alpha1.ConsulCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("consulcluster"), name)
	}
	return obj.(*v1alpha1.ConsulCluster), nil
}
