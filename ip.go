package client

import (
	"encoding/json"
	"net/http"

	"gitlab.com/newsletter2go/hrobot-go/models"
)

func (c *Client) IPGetList() ([]models.IP, error) {
	url := c.baseURL + "/ip"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var ips []models.IPResponse
	err = json.Unmarshal(bytes, &ips)
	if err != nil {
		return nil, err
	}

	var data []models.IP
	for _, ip := range ips {
		data = append(data, ip.IP)
	}

	return data, nil
}
