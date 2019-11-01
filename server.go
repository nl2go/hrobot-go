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

func (c *Client) ServerSetName(ip, name string) error {
	url := fmt.Sprintf(c.baseURL+"/server/%s", ip)

	formData := neturl.Values{}
	formData.Set("server_name", name)

	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err = c.doRequest(req)
	return err
}

func (c *Client) ServerReverse(ip string) error {
	url := fmt.Sprintf(c.baseURL+"/server/%s/reversal", ip)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}
