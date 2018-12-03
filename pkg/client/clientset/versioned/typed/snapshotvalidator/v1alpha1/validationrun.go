/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/cisco-sso/snapshot-manager/pkg/apis/snapshotvalidator/v1alpha1"
	scheme "github.com/cisco-sso/snapshot-manager/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ValidationRunsGetter has a method to return a ValidationRunInterface.
// A group's client should implement this interface.
type ValidationRunsGetter interface {
	ValidationRuns(namespace string) ValidationRunInterface
}

// ValidationRunInterface has methods to work with ValidationRun resources.
type ValidationRunInterface interface {
	Create(*v1alpha1.ValidationRun) (*v1alpha1.ValidationRun, error)
	Update(*v1alpha1.ValidationRun) (*v1alpha1.ValidationRun, error)
	UpdateStatus(*v1alpha1.ValidationRun) (*v1alpha1.ValidationRun, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ValidationRun, error)
	List(opts v1.ListOptions) (*v1alpha1.ValidationRunList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ValidationRun, err error)
	ValidationRunExpansion
}

// validationRuns implements ValidationRunInterface
type validationRuns struct {
	client rest.Interface
	ns     string
}

// newValidationRuns returns a ValidationRuns
func newValidationRuns(c *SnapshotvalidatorV1alpha1Client, namespace string) *validationRuns {
	return &validationRuns{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the validationRun, and returns the corresponding validationRun object, and an error if there is any.
func (c *validationRuns) Get(name string, options v1.GetOptions) (result *v1alpha1.ValidationRun, err error) {
	result = &v1alpha1.ValidationRun{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("validationruns").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ValidationRuns that match those selectors.
func (c *validationRuns) List(opts v1.ListOptions) (result *v1alpha1.ValidationRunList, err error) {
	result = &v1alpha1.ValidationRunList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("validationruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested validationRuns.
func (c *validationRuns) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("validationruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a validationRun and creates it.  Returns the server's representation of the validationRun, and an error, if there is any.
func (c *validationRuns) Create(validationRun *v1alpha1.ValidationRun) (result *v1alpha1.ValidationRun, err error) {
	result = &v1alpha1.ValidationRun{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("validationruns").
		Body(validationRun).
		Do().
		Into(result)
	return
}

// Update takes the representation of a validationRun and updates it. Returns the server's representation of the validationRun, and an error, if there is any.
func (c *validationRuns) Update(validationRun *v1alpha1.ValidationRun) (result *v1alpha1.ValidationRun, err error) {
	result = &v1alpha1.ValidationRun{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("validationruns").
		Name(validationRun.Name).
		Body(validationRun).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *validationRuns) UpdateStatus(validationRun *v1alpha1.ValidationRun) (result *v1alpha1.ValidationRun, err error) {
	result = &v1alpha1.ValidationRun{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("validationruns").
		Name(validationRun.Name).
		SubResource("status").
		Body(validationRun).
		Do().
		Into(result)
	return
}

// Delete takes name of the validationRun and deletes it. Returns an error if one occurs.
func (c *validationRuns) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("validationruns").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *validationRuns) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("validationruns").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched validationRun.
func (c *validationRuns) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ValidationRun, err error) {
	result = &v1alpha1.ValidationRun{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("validationruns").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
