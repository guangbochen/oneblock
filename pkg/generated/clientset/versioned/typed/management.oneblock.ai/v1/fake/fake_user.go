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

	v1 "github.com/oneblock-ai/oneblock/pkg/apis/management.oneblock.ai/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUsers implements UserInterface
type FakeUsers struct {
	Fake *FakeManagementV1
}

var usersResource = v1.SchemeGroupVersion.WithResource("users")

var usersKind = v1.SchemeGroupVersion.WithKind("User")

// Get takes name of the user, and returns the corresponding user object, and an error if there is any.
func (c *FakeUsers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(usersResource, name), &v1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.User), err
}

// List takes label and field selectors, and returns the list of Users that match those selectors.
func (c *FakeUsers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.UserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(usersResource, usersKind, opts), &v1.UserList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.UserList{ListMeta: obj.(*v1.UserList).ListMeta}
	for _, item := range obj.(*v1.UserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested users.
func (c *FakeUsers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(usersResource, opts))
}

// Create takes the representation of a user and creates it.  Returns the server's representation of the user, and an error, if there is any.
func (c *FakeUsers) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) (result *v1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(usersResource, user), &v1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.User), err
}

// Update takes the representation of a user and updates it. Returns the server's representation of the user, and an error, if there is any.
func (c *FakeUsers) Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) (result *v1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(usersResource, user), &v1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.User), err
}

// Delete takes name of the user and deletes it. Returns an error if one occurs.
func (c *FakeUsers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(usersResource, name, opts), &v1.User{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUsers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(usersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.UserList{})
	return err
}

// Patch applies the patch and returns the patched user.
func (c *FakeUsers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(usersResource, name, pt, data, subresources...), &v1.User{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.User), err
}
