/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package internalversion

import (
	quota "github.com/openshift/origin/pkg/quota/apis/quota"
	scheme "github.com/openshift/origin/pkg/quota/generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterResourceQuotasGetter has a method to return a ClusterResourceQuotaInterface.
// A group's client should implement this interface.
type ClusterResourceQuotasGetter interface {
	ClusterResourceQuotas() ClusterResourceQuotaInterface
}

// ClusterResourceQuotaInterface has methods to work with ClusterResourceQuota resources.
type ClusterResourceQuotaInterface interface {
	Create(*quota.ClusterResourceQuota) (*quota.ClusterResourceQuota, error)
	Update(*quota.ClusterResourceQuota) (*quota.ClusterResourceQuota, error)
	UpdateStatus(*quota.ClusterResourceQuota) (*quota.ClusterResourceQuota, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*quota.ClusterResourceQuota, error)
	List(opts v1.ListOptions) (*quota.ClusterResourceQuotaList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *quota.ClusterResourceQuota, err error)
	ClusterResourceQuotaExpansion
}

// clusterResourceQuotas implements ClusterResourceQuotaInterface
type clusterResourceQuotas struct {
	client rest.Interface
}

// newClusterResourceQuotas returns a ClusterResourceQuotas
func newClusterResourceQuotas(c *QuotaClient) *clusterResourceQuotas {
	return &clusterResourceQuotas{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterResourceQuota, and returns the corresponding clusterResourceQuota object, and an error if there is any.
func (c *clusterResourceQuotas) Get(name string, options v1.GetOptions) (result *quota.ClusterResourceQuota, err error) {
	result = &quota.ClusterResourceQuota{}
	err = c.client.Get().
		Resource("clusterresourcequotas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterResourceQuotas that match those selectors.
func (c *clusterResourceQuotas) List(opts v1.ListOptions) (result *quota.ClusterResourceQuotaList, err error) {
	result = &quota.ClusterResourceQuotaList{}
	err = c.client.Get().
		Resource("clusterresourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterResourceQuotas.
func (c *clusterResourceQuotas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("clusterresourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a clusterResourceQuota and creates it.  Returns the server's representation of the clusterResourceQuota, and an error, if there is any.
func (c *clusterResourceQuotas) Create(clusterResourceQuota *quota.ClusterResourceQuota) (result *quota.ClusterResourceQuota, err error) {
	result = &quota.ClusterResourceQuota{}
	err = c.client.Post().
		Resource("clusterresourcequotas").
		Body(clusterResourceQuota).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterResourceQuota and updates it. Returns the server's representation of the clusterResourceQuota, and an error, if there is any.
func (c *clusterResourceQuotas) Update(clusterResourceQuota *quota.ClusterResourceQuota) (result *quota.ClusterResourceQuota, err error) {
	result = &quota.ClusterResourceQuota{}
	err = c.client.Put().
		Resource("clusterresourcequotas").
		Name(clusterResourceQuota.Name).
		Body(clusterResourceQuota).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterResourceQuotas) UpdateStatus(clusterResourceQuota *quota.ClusterResourceQuota) (result *quota.ClusterResourceQuota, err error) {
	result = &quota.ClusterResourceQuota{}
	err = c.client.Put().
		Resource("clusterresourcequotas").
		Name(clusterResourceQuota.Name).
		SubResource("status").
		Body(clusterResourceQuota).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterResourceQuota and deletes it. Returns an error if one occurs.
func (c *clusterResourceQuotas) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterresourcequotas").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterResourceQuotas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("clusterresourcequotas").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterResourceQuota.
func (c *clusterResourceQuotas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *quota.ClusterResourceQuota, err error) {
	result = &quota.ClusterResourceQuota{}
	err = c.client.Patch(pt).
		Resource("clusterresourcequotas").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
