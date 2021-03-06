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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/luolanzone/tunnel-example/apis/gateway/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTunnelEndpoints implements TunnelEndpointInterface
type FakeTunnelEndpoints struct {
	Fake *FakeGatewayV1alpha1
}

var tunnelendpointsResource = schema.GroupVersionResource{Group: "gateway.luolanzone.io", Version: "v1alpha1", Resource: "tunnelendpoints"}

var tunnelendpointsKind = schema.GroupVersionKind{Group: "gateway.luolanzone.io", Version: "v1alpha1", Kind: "TunnelEndpoint"}

// Get takes name of the tunnelEndpoint, and returns the corresponding tunnelEndpoint object, and an error if there is any.
func (c *FakeTunnelEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TunnelEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(tunnelendpointsResource, name), &v1alpha1.TunnelEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TunnelEndpoint), err
}

// List takes label and field selectors, and returns the list of TunnelEndpoints that match those selectors.
func (c *FakeTunnelEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TunnelEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(tunnelendpointsResource, tunnelendpointsKind, opts), &v1alpha1.TunnelEndpointList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TunnelEndpointList{ListMeta: obj.(*v1alpha1.TunnelEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.TunnelEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tunnelEndpoints.
func (c *FakeTunnelEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(tunnelendpointsResource, opts))
}

// Create takes the representation of a tunnelEndpoint and creates it.  Returns the server's representation of the tunnelEndpoint, and an error, if there is any.
func (c *FakeTunnelEndpoints) Create(ctx context.Context, tunnelEndpoint *v1alpha1.TunnelEndpoint, opts v1.CreateOptions) (result *v1alpha1.TunnelEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(tunnelendpointsResource, tunnelEndpoint), &v1alpha1.TunnelEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TunnelEndpoint), err
}

// Update takes the representation of a tunnelEndpoint and updates it. Returns the server's representation of the tunnelEndpoint, and an error, if there is any.
func (c *FakeTunnelEndpoints) Update(ctx context.Context, tunnelEndpoint *v1alpha1.TunnelEndpoint, opts v1.UpdateOptions) (result *v1alpha1.TunnelEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(tunnelendpointsResource, tunnelEndpoint), &v1alpha1.TunnelEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TunnelEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTunnelEndpoints) UpdateStatus(ctx context.Context, tunnelEndpoint *v1alpha1.TunnelEndpoint, opts v1.UpdateOptions) (*v1alpha1.TunnelEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(tunnelendpointsResource, "status", tunnelEndpoint), &v1alpha1.TunnelEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TunnelEndpoint), err
}

// Delete takes name of the tunnelEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeTunnelEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(tunnelendpointsResource, name), &v1alpha1.TunnelEndpoint{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTunnelEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(tunnelendpointsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TunnelEndpointList{})
	return err
}

// Patch applies the patch and returns the patched tunnelEndpoint.
func (c *FakeTunnelEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TunnelEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(tunnelendpointsResource, name, pt, data, subresources...), &v1alpha1.TunnelEndpoint{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TunnelEndpoint), err
}
