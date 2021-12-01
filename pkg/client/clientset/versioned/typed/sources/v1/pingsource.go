/*
Copyright 2021 The Knative Authors

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

package v1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "knative.dev/eventing/pkg/apis/sources/v1"
	scheme "knative.dev/eventing/pkg/client/clientset/versioned/scheme"
)

// PingSourcesGetter has a method to return a PingSourceInterface.
// A group's client should implement this interface.
type PingSourcesGetter interface {
	PingSources(namespace string) PingSourceInterface
}

// PingSourceInterface has methods to work with PingSource resources.
type PingSourceInterface interface {
	Create(ctx context.Context, pingSource *v1.PingSource, opts metav1.CreateOptions) (*v1.PingSource, error)
	Update(ctx context.Context, pingSource *v1.PingSource, opts metav1.UpdateOptions) (*v1.PingSource, error)
	UpdateStatus(ctx context.Context, pingSource *v1.PingSource, opts metav1.UpdateOptions) (*v1.PingSource, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.PingSource, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PingSourceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PingSource, err error)
	PingSourceExpansion
}

// pingSources implements PingSourceInterface
type pingSources struct {
	client rest.Interface
	ns     string
}

// newPingSources returns a PingSources
func newPingSources(c *SourcesV1Client, namespace string) *pingSources {
	return &pingSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pingSource, and returns the corresponding pingSource object, and an error if there is any.
func (c *pingSources) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.PingSource, err error) {
	result = &v1.PingSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pingsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PingSources that match those selectors.
func (c *pingSources) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PingSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PingSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pingsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pingSources.
func (c *pingSources) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pingsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a pingSource and creates it.  Returns the server's representation of the pingSource, and an error, if there is any.
func (c *pingSources) Create(ctx context.Context, pingSource *v1.PingSource, opts metav1.CreateOptions) (result *v1.PingSource, err error) {
	result = &v1.PingSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pingsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pingSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a pingSource and updates it. Returns the server's representation of the pingSource, and an error, if there is any.
func (c *pingSources) Update(ctx context.Context, pingSource *v1.PingSource, opts metav1.UpdateOptions) (result *v1.PingSource, err error) {
	result = &v1.PingSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pingsources").
		Name(pingSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pingSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *pingSources) UpdateStatus(ctx context.Context, pingSource *v1.PingSource, opts metav1.UpdateOptions) (result *v1.PingSource, err error) {
	result = &v1.PingSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pingsources").
		Name(pingSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pingSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the pingSource and deletes it. Returns an error if one occurs.
func (c *pingSources) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pingsources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pingSources) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pingsources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched pingSource.
func (c *pingSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PingSource, err error) {
	result = &v1.PingSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pingsources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}