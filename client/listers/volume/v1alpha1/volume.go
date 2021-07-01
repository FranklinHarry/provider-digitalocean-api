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
	v1alpha1 "kubeform.dev/provider-digitalocean-api/apis/volume/v1alpha1"
)

// VolumeLister helps list Volumes.
// All objects returned here must be treated as read-only.
type VolumeLister interface {
	// List lists all Volumes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Volume, err error)
	// Volumes returns an object that can list and get Volumes.
	Volumes(namespace string) VolumeNamespaceLister
	VolumeListerExpansion
}

// volumeLister implements the VolumeLister interface.
type volumeLister struct {
	indexer cache.Indexer
}

// NewVolumeLister returns a new VolumeLister.
func NewVolumeLister(indexer cache.Indexer) VolumeLister {
	return &volumeLister{indexer: indexer}
}

// List lists all Volumes in the indexer.
func (s *volumeLister) List(selector labels.Selector) (ret []*v1alpha1.Volume, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Volume))
	})
	return ret, err
}

// Volumes returns an object that can list and get Volumes.
func (s *volumeLister) Volumes(namespace string) VolumeNamespaceLister {
	return volumeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VolumeNamespaceLister helps list and get Volumes.
// All objects returned here must be treated as read-only.
type VolumeNamespaceLister interface {
	// List lists all Volumes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Volume, err error)
	// Get retrieves the Volume from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Volume, error)
	VolumeNamespaceListerExpansion
}

// volumeNamespaceLister implements the VolumeNamespaceLister
// interface.
type volumeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Volumes in the indexer for a given namespace.
func (s volumeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Volume, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Volume))
	})
	return ret, err
}

// Get retrieves the Volume from the indexer for a given namespace and name.
func (s volumeNamespaceLister) Get(name string) (*v1alpha1.Volume, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("volume"), name)
	}
	return obj.(*v1alpha1.Volume), nil
}
