package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	neturl "net/url"
	"strings"

	"gitlab.com/newsletter2go/hrobot-go/models"
)

func (c *Client) ResetGet(ip string) (*models.Reset, error) {
	url := fmt.Sprintf(c.baseURL+"/reset/%s", ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var resetResp models.ResetResponse
	err = json.Unmarshal(bytes, &resetResp)
	if err != nil {
		return nil, err
	}

	return &resetResp.Reset, nil
}

func (c *Client) ResetSet(ip string, input *models.ResetSetInput) (*models.ResetPost, error) {
	url := fmt.Sprintf(c.baseURL+"/reset/%s", ip)

	formData := neturl.Values{}
	formData.Set("type", input.Type)

	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	bytes, err := c.doRequest(req)
	var resetResp models.ResetPostResponse
	err = json.Unmarshal(bytes, &resetResp)
	if err != nil {
		return nil, err
	}

	return &resetResp.Reset, nil
}
