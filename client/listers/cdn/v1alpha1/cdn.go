/*
Copyright AppsCode Inc. and Contributors

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/provider-digitalocean-api/apis/cdn/v1alpha1"
)

// CdnLister helps list Cdns.
// All objects returned here must be treated as read-only.
type CdnLister interface {
	// List lists all Cdns in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Cdn, err error)
	// Cdns returns an object that can list and get Cdns.
	Cdns(namespace string) CdnNamespaceLister
	CdnListerExpansion
}

// cdnLister implements the CdnLister interface.
type cdnLister struct {
	indexer cache.Indexer
}

// NewCdnLister returns a new CdnLister.
func NewCdnLister(indexer cache.Indexer) CdnLister {
	return &cdnLister{indexer: indexer}
}

// List lists all Cdns in the indexer.
func (s *cdnLister) List(selector labels.Selector) (ret []*v1alpha1.Cdn, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Cdn))
	})
	return ret, err
}

// Cdns returns an object that can list and get Cdns.
func (s *cdnLister) Cdns(namespace string) CdnNamespaceLister {
	return cdnNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CdnNamespaceLister helps list and get Cdns.
// All objects returned here must be treated as read-only.
type CdnNamespaceLister interface {
	// List lists all Cdns in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Cdn, err error)
	// Get retrieves the Cdn from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Cdn, error)
	CdnNamespaceListerExpansion
}

// cdnNamespaceLister implements the CdnNamespaceLister
// interface.
type cdnNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Cdns in the indexer for a given namespace.
func (s cdnNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Cdn, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Cdn))
	})
	return ret, err
}

// Get retrieves the Cdn from the indexer for a given namespace and name.
func (s cdnNamespaceLister) Get(name string) (*v1alpha1.Cdn, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cdn"), name)
	}
	return obj.(*v1alpha1.Cdn), nil
}
