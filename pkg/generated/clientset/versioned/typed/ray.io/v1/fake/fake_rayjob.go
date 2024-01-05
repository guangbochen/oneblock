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

package fake

import (
	"context"

	v1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRayJobs implements RayJobInterface
type FakeRayJobs struct {
	Fake *FakeRayV1
}

var rayjobsResource = v1.SchemeGroupVersion.WithResource("rayjobs")

var rayjobsKind = v1.SchemeGroupVersion.WithKind("RayJob")

// Get takes name of the rayJob, and returns the corresponding rayJob object, and an error if there is any.
func (c *FakeRayJobs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RayJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(rayjobsResource, name), &v1.RayJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayJob), err
}

// List takes label and field selectors, and returns the list of RayJobs that match those selectors.
func (c *FakeRayJobs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RayJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(rayjobsResource, rayjobsKind, opts), &v1.RayJobList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RayJobList{ListMeta: obj.(*v1.RayJobList).ListMeta}
	for _, item := range obj.(*v1.RayJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rayJobs.
func (c *FakeRayJobs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(rayjobsResource, opts))
}

// Create takes the representation of a rayJob and creates it.  Returns the server's representation of the rayJob, and an error, if there is any.
func (c *FakeRayJobs) Create(ctx context.Context, rayJob *v1.RayJob, opts metav1.CreateOptions) (result *v1.RayJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(rayjobsResource, rayJob), &v1.RayJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayJob), err
}

// Update takes the representation of a rayJob and updates it. Returns the server's representation of the rayJob, and an error, if there is any.
func (c *FakeRayJobs) Update(ctx context.Context, rayJob *v1.RayJob, opts metav1.UpdateOptions) (result *v1.RayJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(rayjobsResource, rayJob), &v1.RayJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRayJobs) UpdateStatus(ctx context.Context, rayJob *v1.RayJob, opts metav1.UpdateOptions) (*v1.RayJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(rayjobsResource, "status", rayJob), &v1.RayJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayJob), err
}

// Delete takes name of the rayJob and deletes it. Returns an error if one occurs.
func (c *FakeRayJobs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(rayjobsResource, name, opts), &v1.RayJob{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRayJobs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(rayjobsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RayJobList{})
	return err
}

// Patch applies the patch and returns the patched rayJob.
func (c *FakeRayJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RayJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(rayjobsResource, name, pt, data, subresources...), &v1.RayJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayJob), err
}
