/*
Copyright 2021 TriggerMesh Inc.

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
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/targets/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ElasticsearchTargetLister helps list ElasticsearchTargets.
// All objects returned here must be treated as read-only.
type ElasticsearchTargetLister interface {
	// List lists all ElasticsearchTargets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ElasticsearchTarget, err error)
	// ElasticsearchTargets returns an object that can list and get ElasticsearchTargets.
	ElasticsearchTargets(namespace string) ElasticsearchTargetNamespaceLister
	ElasticsearchTargetListerExpansion
}

// elasticsearchTargetLister implements the ElasticsearchTargetLister interface.
type elasticsearchTargetLister struct {
	indexer cache.Indexer
}

// NewElasticsearchTargetLister returns a new ElasticsearchTargetLister.
func NewElasticsearchTargetLister(indexer cache.Indexer) ElasticsearchTargetLister {
	return &elasticsearchTargetLister{indexer: indexer}
}

// List lists all ElasticsearchTargets in the indexer.
func (s *elasticsearchTargetLister) List(selector labels.Selector) (ret []*v1alpha1.ElasticsearchTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElasticsearchTarget))
	})
	return ret, err
}

// ElasticsearchTargets returns an object that can list and get ElasticsearchTargets.
func (s *elasticsearchTargetLister) ElasticsearchTargets(namespace string) ElasticsearchTargetNamespaceLister {
	return elasticsearchTargetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ElasticsearchTargetNamespaceLister helps list and get ElasticsearchTargets.
// All objects returned here must be treated as read-only.
type ElasticsearchTargetNamespaceLister interface {
	// List lists all ElasticsearchTargets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ElasticsearchTarget, err error)
	// Get retrieves the ElasticsearchTarget from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ElasticsearchTarget, error)
	ElasticsearchTargetNamespaceListerExpansion
}

// elasticsearchTargetNamespaceLister implements the ElasticsearchTargetNamespaceLister
// interface.
type elasticsearchTargetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ElasticsearchTargets in the indexer for a given namespace.
func (s elasticsearchTargetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ElasticsearchTarget, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElasticsearchTarget))
	})
	return ret, err
}

// Get retrieves the ElasticsearchTarget from the indexer for a given namespace and name.
func (s elasticsearchTargetNamespaceLister) Get(name string) (*v1alpha1.ElasticsearchTarget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("elasticsearchtarget"), name)
	}
	return obj.(*v1alpha1.ElasticsearchTarget), nil
}