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

	scheme "github.com/oneblock-ai/oneblock/pkg/generated/clientset/versioned/scheme"
	v1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RayJobsGetter has a method to return a RayJobInterface.
// A group's client should implement this interface.
type RayJobsGetter interface {
	RayJobs() RayJobInterface
}

// RayJobInterface has methods to work with RayJob resources.
type RayJobInterface interface {
	Create(ctx context.Context, rayJob *v1.RayJob, opts metav1.CreateOptions) (*v1.RayJob, error)
	Update(ctx context.Context, rayJob *v1.RayJob, opts metav1.UpdateOptions) (*v1.RayJob, error)
	UpdateStatus(ctx context.Context, rayJob *v1.RayJob, opts metav1.UpdateOptions) (*v1.RayJob, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.RayJob, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.RayJobList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RayJob, err error)
	RayJobExpansion
}

// rayJobs implements RayJobInterface
type rayJobs struct {
	client rest.Interface
}

// newRayJobs returns a RayJobs
func newRayJobs(c *RayV1Client) *rayJobs {
	return &rayJobs{
		client: c.RESTClient(),
	}
}

// Get takes name of the rayJob, and returns the corresponding rayJob object, and an error if there is any.
func (c *rayJobs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RayJob, err error) {
	result = &v1.RayJob{}
	err = c.client.Get().
		Resource("rayjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RayJobs that match those selectors.
func (c *rayJobs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RayJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RayJobList{}
	err = c.client.Get().
		Resource("rayjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested rayJobs.
func (c *rayJobs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("rayjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a rayJob and creates it.  Returns the server's representation of the rayJob, and an error, if there is any.
func (c *rayJobs) Create(ctx context.Context, rayJob *v1.RayJob, opts metav1.CreateOptions) (result *v1.RayJob, err error) {
	result = &v1.RayJob{}
	err = c.client.Post().
		Resource("rayjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(rayJob).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a rayJob and updates it. Returns the server's representation of the rayJob, and an error, if there is any.
func (c *rayJobs) Update(ctx context.Context, rayJob *v1.RayJob, opts metav1.UpdateOptions) (result *v1.RayJob, err error) {
	result = &v1.RayJob{}
	err = c.client.Put().
		Resource("rayjobs").
		Name(rayJob.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(rayJob).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *rayJobs) UpdateStatus(ctx context.Context, rayJob *v1.RayJob, opts metav1.UpdateOptions) (result *v1.RayJob, err error) {
	result = &v1.RayJob{}
	err = c.client.Put().
		Resource("rayjobs").
		Name(rayJob.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(rayJob).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the rayJob and deletes it. Returns an error if one occurs.
func (c *rayJobs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("rayjobs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rayJobs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("rayjobs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched rayJob.
func (c *rayJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RayJob, err error) {
	result = &v1.RayJob{}
	err = c.client.Patch(pt).
		Resource("rayjobs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
