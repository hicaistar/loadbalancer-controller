/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/caicloud/clientset/pkg/apis/resource/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StorageServiceLister helps list StorageServices.
type StorageServiceLister interface {
	// List lists all StorageServices in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.StorageService, err error)
	// Get retrieves the StorageService from the index for a given name.
	Get(name string) (*v1beta1.StorageService, error)
	StorageServiceListerExpansion
}

// storageServiceLister implements the StorageServiceLister interface.
type storageServiceLister struct {
	indexer cache.Indexer
}

// NewStorageServiceLister returns a new StorageServiceLister.
func NewStorageServiceLister(indexer cache.Indexer) StorageServiceLister {
	return &storageServiceLister{indexer: indexer}
}

// List lists all StorageServices in the indexer.
func (s *storageServiceLister) List(selector labels.Selector) (ret []*v1beta1.StorageService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.StorageService))
	})
	return ret, err
}

// Get retrieves the StorageService from the index for a given name.
func (s *storageServiceLister) Get(name string) (*v1beta1.StorageService, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("storageservice"), name)
	}
	return obj.(*v1beta1.StorageService), nil
}
