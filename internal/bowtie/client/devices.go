package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DevicePayload struct {
	Devices map[string]Device `json:"devices"`
}

type Device struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	IPV6            string `json:"ipv6"`
	PublicKey       string `json:"public_key"`
	Serial          string `json:"serial"`
	State           string `json:"state"`
	ControllerID    string `json:"controller_id"`
	OwnedByOrg      string `json:"owned_by_org"`
	AssignedToUser  string `json:"assigned_to_user"`
	DeviceType      string `json:"device_type"`
	DeviceOS        string `json:"device_os"`
	LastSeen        string `json:"last_seen"`
	LastSeenVersion string `json:"last_seen_version"`
}

func (c *Client) DeleteDevice(id string) error {
	req, err := http.NewRequest(http.MethodDelete, c.getHostURL(fmt.Sprintf("/device/%s", id)), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

func (c *Client) ListDevices() (map[string]Device, error) {
	req, err := http.NewRequest(http.MethodGet, c.getHostURL("/device"), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var devices DevicePayload = DevicePayload{}
	err = json.Unmarshal(body, &devices)

	return devices.Devices, err
}
