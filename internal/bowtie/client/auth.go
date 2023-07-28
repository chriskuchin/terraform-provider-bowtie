package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Me struct {
	User    User `json:"user"`
	Devices map[string]Device
}

type User struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	AuthZDevices      bool   `json:"authz_devices"`
	AuthZPolicies     bool   `json:"authz_policies"`
	AuthZControlPanel bool   `json:"authz_control_panel"`
	AuthZUsers        bool   `json:"authz_users"`
	Role              string `json:"role"`
}

type Device struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	IPV6           string `json:"ipv6"`
	PublicKey      string `json:"public_key"`
	Serial         string `json:"serial"`
	State          string `json:"state"`
	ControllerID   string `json:"controller_id"`
	OwnedByOrg     string `json:"owned_by_org"`
	AssignedToUser string `json:"assigned_to_user"`
	DeviceType     string `json:"device_type"`
	DeviceOS       string `json:"device_os"`
	LastSeen       string `json:"last_seen"`
}

func (c *Client) Login() error {
	payload, err := json.Marshal(c.auth)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.getHostURL("/user/login"), strings.NewReader(string(payload)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to login: %s", res.Status)
	}

	return nil
}

func (c *Client) WhoAmI() (*Me, error) {
	req, err := http.NewRequest("GET", c.getHostURL("/user/me"), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var me *Me = &Me{}
	json.Unmarshal(body, me)

	return me, nil
}
