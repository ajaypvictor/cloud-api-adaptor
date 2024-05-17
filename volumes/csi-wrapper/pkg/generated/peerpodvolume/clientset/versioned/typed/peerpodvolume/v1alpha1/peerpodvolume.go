// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/confidential-containers/cloud-api-adaptor/volumes/csi-wrapper/pkg/apis/peerpodvolume/v1alpha1"
	scheme "github.com/confidential-containers/cloud-api-adaptor/volumes/csi-wrapper/pkg/generated/peerpodvolume/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PeerpodVolumesGetter has a method to return a PeerpodVolumeInterface.
// A group's client should implement this interface.
type PeerpodVolumesGetter interface {
	PeerpodVolumes(namespace string) PeerpodVolumeInterface
}

// PeerpodVolumeInterface has methods to work with PeerpodVolume resources.
type PeerpodVolumeInterface interface {
	Create(ctx context.Context, peerpodVolume *v1alpha1.PeerpodVolume, opts v1.CreateOptions) (*v1alpha1.PeerpodVolume, error)
	Update(ctx context.Context, peerpodVolume *v1alpha1.PeerpodVolume, opts v1.UpdateOptions) (*v1alpha1.PeerpodVolume, error)
	UpdateStatus(ctx context.Context, peerpodVolume *v1alpha1.PeerpodVolume, opts v1.UpdateOptions) (*v1alpha1.PeerpodVolume, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PeerpodVolume, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PeerpodVolumeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PeerpodVolume, err error)
	PeerpodVolumeExpansion
}

// peerpodVolumes implements PeerpodVolumeInterface
type peerpodVolumes struct {
	client rest.Interface
	ns     string
}

// newPeerpodVolumes returns a PeerpodVolumes
func newPeerpodVolumes(c *ConfidentialcontainersV1alpha1Client, namespace string) *peerpodVolumes {
	return &peerpodVolumes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the peerpodVolume, and returns the corresponding peerpodVolume object, and an error if there is any.
func (c *peerpodVolumes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PeerpodVolume, err error) {
	result = &v1alpha1.PeerpodVolume{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("peerpodvolumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PeerpodVolumes that match those selectors.
func (c *peerpodVolumes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PeerpodVolumeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PeerpodVolumeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("peerpodvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested peerpodVolumes.
func (c *peerpodVolumes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("peerpodvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a peerpodVolume and creates it.  Returns the server's representation of the peerpodVolume, and an error, if there is any.
func (c *peerpodVolumes) Create(ctx context.Context, peerpodVolume *v1alpha1.PeerpodVolume, opts v1.CreateOptions) (result *v1alpha1.PeerpodVolume, err error) {
	result = &v1alpha1.PeerpodVolume{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("peerpodvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(peerpodVolume).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a peerpodVolume and updates it. Returns the server's representation of the peerpodVolume, and an error, if there is any.
func (c *peerpodVolumes) Update(ctx context.Context, peerpodVolume *v1alpha1.PeerpodVolume, opts v1.UpdateOptions) (result *v1alpha1.PeerpodVolume, err error) {
	result = &v1alpha1.PeerpodVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("peerpodvolumes").
		Name(peerpodVolume.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(peerpodVolume).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *peerpodVolumes) UpdateStatus(ctx context.Context, peerpodVolume *v1alpha1.PeerpodVolume, opts v1.UpdateOptions) (result *v1alpha1.PeerpodVolume, err error) {
	result = &v1alpha1.PeerpodVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("peerpodvolumes").
		Name(peerpodVolume.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(peerpodVolume).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the peerpodVolume and deletes it. Returns an error if one occurs.
func (c *peerpodVolumes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("peerpodvolumes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *peerpodVolumes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("peerpodvolumes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched peerpodVolume.
func (c *peerpodVolumes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PeerpodVolume, err error) {
	result = &v1alpha1.PeerpodVolume{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("peerpodvolumes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}