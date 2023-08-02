package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (c *Client) CreateDNS(name string, serverAddrs []Server, includeOnlySites []string, isCounted, isLog, isDropA, isDropAll bool, exlude []DNSExclude) (string, error) {
	id := uuid.NewString()
	return id, c.UpsertDNS(id, name, serverAddrs, includeOnlySites, isCounted, isLog, isDropA, isDropAll, exlude)
}

func (c *Client) UpsertDNS(id, name string, serverAddrs []Server, includeOnlySites []string, isCounted, isLog, isDropA, isDropAll bool, exlude []DNSExclude) error {
	var servers map[string]Server = map[string]Server{}
	for _, addr := range serverAddrs {
		servers[addr.ID] = addr
	}

	var dnsExclude map[string]DNSExclude = map[string]DNSExclude{}
	for _, record := range exlude {
		dnsExclude[record.ID] = record
	}
	var payload DNS = DNS{
		ID:               id,
		Name:             name,
		IsDNS64:          true,
		Servers:          servers,
		IncludeOnlySites: includeOnlySites,
		IsCounted:        isCounted,
		IsLog:            isLog,
		IsDropA:          isDropA,
		IsDropAll:        isDropAll,
		DNS64Exclude:     dnsExclude,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, c.getHostURL("/organization/dns/upsert"), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

func (c *Client) DeleteDNS(id string) error {
	req, err := http.NewRequest(http.MethodDelete, c.getHostURL(fmt.Sprintf("/organization/dns/%s", id)), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

func (c *Client) GetDNS(id string) (*DNS, error) {
	org, err := c.GetOrganization()
	if err != nil {
		return nil, err
	}

	result, ok := org.DNS[id]
	if !ok {
		return nil, fmt.Errorf("failed to locate the dns object")
	}

	return &result, nil
}
