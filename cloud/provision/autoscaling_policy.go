package provision

import (
	"fmt"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/cloud/rest"
)

// AutoscalingPoliciesGetter is the getter for autoscaling policies
type AutoscalingPoliciesGetter interface {
	AutoscalingPolicies(organizationID, clusterID string) AutoscalingPolicyInterface
}

// AutoscalingPolicyInterface is the interface for autoscaling policies
type AutoscalingPolicyInterface interface {
	Create(*types.AutoscalingPolicy) (*types.AutoscalingPolicy, error)
	Get(id string) (*types.AutoscalingPolicy, error)
	Delete(id string) error
	List() ([]types.AutoscalingPolicy, error)
}

// autoscalingPolicies implements AutoscalingPolicyInterface
type autoscalingPolicies struct {
	client         rest.Interface
	organizationID string
	clusterID      string
}

func newAutoscalingPolicies(c *Client, organizationID, clusterID string) *autoscalingPolicies {
	return &autoscalingPolicies{
		client:         c.RESTClient(),
		organizationID: organizationID,
		clusterID:      clusterID,
	}
}

// Create creates an autoscaling policy
func (c *autoscalingPolicies) Create(asp *types.AutoscalingPolicy) (*types.AutoscalingPolicy, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/autoscaling-policies",
		c.organizationID, c.clusterID)
	var out types.AutoscalingPolicy
	return &out, c.client.Post(path, asp, &out)
}

// Get gets an autoscaling policy
func (c *autoscalingPolicies) Get(id string) (*types.AutoscalingPolicy, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/autoscaling-policies/%s",
		c.organizationID, c.clusterID, id)
	var out types.AutoscalingPolicy
	return &out, c.client.Get(path, &out)
}

// Delete deletes an autoscaling policy
func (c *autoscalingPolicies) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/autoscaling-policies/%s",
		c.organizationID, c.clusterID, id)
	return c.client.Delete(path)
}

// List lists all autoscaling policies
func (c *autoscalingPolicies) List() ([]types.AutoscalingPolicy, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/autoscaling-policies",
		c.organizationID, c.clusterID)
	out := make([]types.AutoscalingPolicy, 0)
	return out, c.client.Get(path, &out)
}
