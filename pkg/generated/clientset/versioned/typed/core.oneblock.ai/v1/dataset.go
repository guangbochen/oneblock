/*
Copyright 2023 1block.ai.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/oneblock-ai/oneblock/pkg/apis/core.oneblock.ai/v1"
	scheme "github.com/oneblock-ai/oneblock/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DatasetsGetter has a method to return a DatasetInterface.
// A group's client should implement this interface.
type DatasetsGetter interface {
	Datasets(namespace string) DatasetInterface
}

// DatasetInterface has methods to work with Dataset resources.
type DatasetInterface interface {
	Create(ctx context.Context, dataset *v1.Dataset, opts metav1.CreateOptions) (*v1.Dataset, error)
	Update(ctx context.Context, dataset *v1.Dataset, opts metav1.UpdateOptions) (*v1.Dataset, error)
	UpdateStatus(ctx context.Context, dataset *v1.Dataset, opts metav1.UpdateOptions) (*v1.Dataset, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Dataset, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DatasetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Dataset, err error)
	DatasetExpansion
}

// datasets implements DatasetInterface
type datasets struct {
	client rest.Interface
	ns     string
}

// newDatasets returns a Datasets
func newDatasets(c *CoreV1Client, namespace string) *datasets {
	return &datasets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dataset, and returns the corresponding dataset object, and an error if there is any.
func (c *datasets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Dataset, err error) {
	result = &v1.Dataset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datasets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Datasets that match those selectors.
func (c *datasets) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DatasetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DatasetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested datasets.
func (c *datasets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("datasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dataset and creates it.  Returns the server's representation of the dataset, and an error, if there is any.
func (c *datasets) Create(ctx context.Context, dataset *v1.Dataset, opts metav1.CreateOptions) (result *v1.Dataset, err error) {
	result = &v1.Dataset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("datasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dataset).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dataset and updates it. Returns the server's representation of the dataset, and an error, if there is any.
func (c *datasets) Update(ctx context.Context, dataset *v1.Dataset, opts metav1.UpdateOptions) (result *v1.Dataset, err error) {
	result = &v1.Dataset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datasets").
		Name(dataset.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dataset).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *datasets) UpdateStatus(ctx context.Context, dataset *v1.Dataset, opts metav1.UpdateOptions) (result *v1.Dataset, err error) {
	result = &v1.Dataset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datasets").
		Name(dataset.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dataset).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dataset and deletes it. Returns an error if one occurs.
func (c *datasets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datasets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *datasets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datasets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dataset.
func (c *datasets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Dataset, err error) {
	result = &v1.Dataset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("datasets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
