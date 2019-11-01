package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	neturl "net/url"
	"strings"

	"gitlab.com/newsletter2go/hrobot-go/models"
)

func (c *Client) ServerGetList() ([]models.Server, error) {
	url := c.baseURL + "/server"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var servers []models.ServerResponse
	err = json.Unmarshal(bytes, &servers)
	if err != nil {
		return nil, err
	}

	var data []models.Server
	for _, server := range servers {
		data = append(data, server.Server)
	}

	return data, nil
}

func (c *Client) ServerGet(ip string) (*models.Server, error) {
	url := fmt.Sprintf(c.baseURL+"/server/%s", ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var serverResp models.ServerResponse
	err = json.Unmarshal(bytes, &serverResp)
	if err != nil {
		return nil, err
	}

	return &serverResp.Server, nil
}

func (c *Client) ServerSetName(ip string, input *models.ServerSetNameInput) (*models.Server, error) {
	url := fmt.Sprintf(c.baseURL+"/server/%s", ip)

	formData := neturl.Values{}
	formData.Set("server_name", input.Name)

	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	bytes, err := c.doRequest(req)
	var serverResp models.ServerResponse
	err = json.Unmarshal(bytes, &serverResp)
	if err != nil {
		return nil, err
	}

	return &serverResp.Server, nil
}

func (c *Client) ServerReverse(ip string) (*models.Cancellation, error) {
	url := fmt.Sprintf(c.baseURL+"/server/%s/reversal", ip)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
	var cancelResp models.CancellationResponse
	err = json.Unmarshal(bytes, &cancelResp)
	if err != nil {
		return nil, err
	}

	return &cancelResp.Cancellation, nil
}
