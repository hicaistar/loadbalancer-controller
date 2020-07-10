/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Datasets returns a DatasetInformer.
	Datasets() DatasetInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Datasets returns a DatasetInformer.
func (v *version) Datasets() DatasetInformer {
	return &datasetInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
