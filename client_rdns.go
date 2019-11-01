package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/newsletter2go/hrobot-go/models"
)

func (c *Client) RDnsGetList() ([]models.Rdns, error) {
	url := c.baseURL + "/rdns"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var rdnsList []models.RdnsResponse
	err = json.Unmarshal(bytes, &rdnsList)
	if err != nil {
		return nil, err
	}

	var data []models.Rdns
	for _, rdns := range rdnsList {
		data = append(data, rdns.Rdns)
	}

	return data, nil
}

func (c *Client) RDnsGet(ip string) (*models.Rdns, error) {
	url := fmt.Sprintf(c.baseURL+"/rdns/%s", ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var rDnsResp models.RdnsResponse
	err = json.Unmarshal(bytes, &rDnsResp)
	if err != nil {
		return nil, err
	}

	return &rDnsResp.Rdns, nil
}
