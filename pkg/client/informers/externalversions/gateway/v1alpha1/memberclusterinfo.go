/*
Copyright 2021.

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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	gatewayv1alpha1 "github.com/luolanzone/tunnel-example/apis/gateway/v1alpha1"
	versioned "github.com/luolanzone/tunnel-example/pkg/client/clientset/versioned"
	internalinterfaces "github.com/luolanzone/tunnel-example/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/luolanzone/tunnel-example/pkg/client/listers/gateway/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MemberClusterInfoInformer provides access to a shared informer and lister for
// MemberClusterInfos.
type MemberClusterInfoInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MemberClusterInfoLister
}

type memberClusterInfoInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewMemberClusterInfoInformer constructs a new informer for MemberClusterInfo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMemberClusterInfoInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMemberClusterInfoInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredMemberClusterInfoInformer constructs a new informer for MemberClusterInfo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMemberClusterInfoInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GatewayV1alpha1().MemberClusterInfos().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GatewayV1alpha1().MemberClusterInfos().Watch(context.TODO(), options)
			},
		},
		&gatewayv1alpha1.MemberClusterInfo{},
		resyncPeriod,
		indexers,
	)
}

func (f *memberClusterInfoInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMemberClusterInfoInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *memberClusterInfoInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&gatewayv1alpha1.MemberClusterInfo{}, f.defaultInformer)
}

func (f *memberClusterInfoInformer) Lister() v1alpha1.MemberClusterInfoLister {
	return v1alpha1.NewMemberClusterInfoLister(f.Informer().GetIndexer())
}
