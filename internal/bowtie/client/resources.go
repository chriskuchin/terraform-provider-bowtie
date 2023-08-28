package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type PoliciesEndpointResponse struct {
	Policies       map[string]BowtiePolicy        `json:"policies"`
	ResourceGroups map[string]BowtieResourceGroup `json:"resource_groups"`
	Resources      map[string]BowtieResource      `json:"resources"`
}

type BowtiePolicy struct {
	ID     string             `json:"id"`
	Source BowtiePolicySource `json:"source"`
	Dest   string             `json:"dest"`
	Action string             `json:"action"`
}

type BowtiePolicySource struct {
	ID        string `json:"id"`
	Predicate string `json:"predicate"`
}

type BowtieResourceGroup struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Inherited []string `json:"inherited"`
	Resources []string `json:"resources"`
}

type BowtieResource struct {
	ID       string                 `json:"id"`
	Name     string                 `json:"name"`
	Protocol string                 `json:"protocol"`
	Location BowtieResourceLocation `json:"location"`
	Ports    BowtieResourcePorts    `json:"ports"`
}

type BowtieResourceLocation struct {
	IP   string `json:"ip,omitempty"`
	CIDR string `json:"cidr,omitempty"`
	DNS  string `json:"dns,omitempty"`
}

type BowtieResourcePorts struct {
	Range      []int64                      `json:"range,omitempty"`
	Collection BowtieResourcePortCollection `json:"collection,omitempty"`
}

type BowtieResourcePortCollection struct {
	Ports []int64 `json:"ports,omitempty"`
}

func (c *Client) CreateResource(name, protocol string, ip, cidr, dns string, portRange, portCollection []int64) (string, BowtieResource, error) {
	id := uuid.NewString()
	resource, err := c.UpsertResource(id, name, protocol, ip, cidr, dns, portRange, portCollection)
	return id, resource, err
}

func (c *Client) UpsertResource(id, name, protocol, ip, cidr, dns string, portRange, portCollection []int64) (BowtieResource, error) {
	payload := BowtieResource{
		ID:       id,
		Name:     name,
		Protocol: protocol,
		Location: BowtieResourceLocation{
			IP:   ip,
			CIDR: cidr,
			DNS:  dns,
		},
		Ports: BowtieResourcePorts{
			Range: portRange,
			Collection: BowtieResourcePortCollection{
				Ports: portCollection,
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return BowtieResource{}, err
	}

	req, err := http.NewRequest(http.MethodPost, c.getHostURL("/policy/upsert_resource"), bytes.NewBuffer(body))
	if err != nil {
		return BowtieResource{}, err
	}

	responsePayload, err := c.doRequest(req)
	if err != nil {
		return BowtieResource{}, err
	}

	var resource BowtieResource = BowtieResource{}
	json.Unmarshal(responsePayload, &resource)

	return resource, nil
}

func (c *Client) GetPoliciesAndResources() (*PoliciesEndpointResponse, error) {
	req, err := http.NewRequest(http.MethodGet, c.getHostURL("/policy"), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var policy *PoliciesEndpointResponse = &PoliciesEndpointResponse{}
	return policy, json.Unmarshal(body, policy)
}

func (c *Client) GetPolicy(id string) (BowtiePolicy, error) {
	policyInfo, err := c.GetPoliciesAndResources()
	if err != nil {
		return BowtiePolicy{}, err
	}

	policy, ok := policyInfo.Policies[id]
	if !ok {
		return BowtiePolicy{}, fmt.Errorf("policy not found")
	}

	return policy, nil
}

func (c *Client) GetResourceGroup(id string) (BowtieResourceGroup, error) {

	return BowtieResourceGroup{}, nil
}

func (c *Client) GetResource(id string) (BowtieResource, error) {
	rp, err := c.GetPoliciesAndResources()
	if err != nil {
		return BowtieResource{}, err
	}

	for _, val := range rp.Resources {
		if val.ID == id {
			return val, nil
		}
	}
	return BowtieResource{}, fmt.Errorf("expected resource not found")
}

func (c *Client) DeletePolicy(id string) error {
	req, err := http.NewRequest(http.MethodDelete, c.getHostURL(fmt.Sprintf("/policy/%s", id)), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

func (c *Client) DeleteResource(id string) error {
	req, err := http.NewRequest(http.MethodDelete, c.getHostURL(fmt.Sprintf("/policy/resource/%s", id)), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

func (c *Client) CreateResourceGroup(name string, resources, resource_groups []string) (string, error) {
	id := uuid.NewString()
	return id, c.UpsertResourceGroup(id, name, resources, resource_groups)
}

func (c *Client) UpsertResourceGroup(id, name string, resources, resource_groups []string) error {
	payload := BowtieResourceGroup{
		ID:        id,
		Name:      name,
		Resources: resources,
		Inherited: resource_groups,
	}

	body, err := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, c.getHostURL("/policy/upsert_resource_group"), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

func (c *Client) DeleteResourceGroup(id string) error {
	req, err := http.NewRequest(http.MethodDelete, c.getHostURL(fmt.Sprintf("/policy/resource_group/%s", id)), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}
