/*
Copyright 2024 1block.ai.

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

	v1 "github.com/oneblock-ai/oneblock/pkg/apis/ml.oneblock.ai/v1"
	scheme "github.com/oneblock-ai/oneblock/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ModelTemplatesGetter has a method to return a ModelTemplateInterface.
// A group's client should implement this interface.
type ModelTemplatesGetter interface {
	ModelTemplates(namespace string) ModelTemplateInterface
}

// ModelTemplateInterface has methods to work with ModelTemplate resources.
type ModelTemplateInterface interface {
	Create(ctx context.Context, modelTemplate *v1.ModelTemplate, opts metav1.CreateOptions) (*v1.ModelTemplate, error)
	Update(ctx context.Context, modelTemplate *v1.ModelTemplate, opts metav1.UpdateOptions) (*v1.ModelTemplate, error)
	UpdateStatus(ctx context.Context, modelTemplate *v1.ModelTemplate, opts metav1.UpdateOptions) (*v1.ModelTemplate, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ModelTemplate, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ModelTemplateList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ModelTemplate, err error)
	ModelTemplateExpansion
}

// modelTemplates implements ModelTemplateInterface
type modelTemplates struct {
	client rest.Interface
	ns     string
}

// newModelTemplates returns a ModelTemplates
func newModelTemplates(c *MlV1Client, namespace string) *modelTemplates {
	return &modelTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the modelTemplate, and returns the corresponding modelTemplate object, and an error if there is any.
func (c *modelTemplates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ModelTemplate, err error) {
	result = &v1.ModelTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("modeltemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ModelTemplates that match those selectors.
func (c *modelTemplates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ModelTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ModelTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("modeltemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested modelTemplates.
func (c *modelTemplates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("modeltemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a modelTemplate and creates it.  Returns the server's representation of the modelTemplate, and an error, if there is any.
func (c *modelTemplates) Create(ctx context.Context, modelTemplate *v1.ModelTemplate, opts metav1.CreateOptions) (result *v1.ModelTemplate, err error) {
	result = &v1.ModelTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("modeltemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(modelTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a modelTemplate and updates it. Returns the server's representation of the modelTemplate, and an error, if there is any.
func (c *modelTemplates) Update(ctx context.Context, modelTemplate *v1.ModelTemplate, opts metav1.UpdateOptions) (result *v1.ModelTemplate, err error) {
	result = &v1.ModelTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("modeltemplates").
		Name(modelTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(modelTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *modelTemplates) UpdateStatus(ctx context.Context, modelTemplate *v1.ModelTemplate, opts metav1.UpdateOptions) (result *v1.ModelTemplate, err error) {
	result = &v1.ModelTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("modeltemplates").
		Name(modelTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(modelTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the modelTemplate and deletes it. Returns an error if one occurs.
func (c *modelTemplates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("modeltemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *modelTemplates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("modeltemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched modelTemplate.
func (c *modelTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ModelTemplate, err error) {
	result = &v1.ModelTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("modeltemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
