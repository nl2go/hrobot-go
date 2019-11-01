package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	neturl "net/url"
	"strconv"
	"strings"

	"gitlab.com/newsletter2go/hrobot-go/models"
)

func (c *Client) BootRescueGet(ip string) (*models.RescueOptions, error) {
	url := fmt.Sprintf(c.baseURL+"/boot/%s/rescue", ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var rescueResp models.RescueGetResponse
	err = json.Unmarshal(bytes, &rescueResp)
	if err != nil {
		return nil, err
	}

	return &rescueResp.Rescue, nil
}

func (c *Client) BootRescueSet(ip string, input *models.RescueSetInput) (*models.RescueValues, error) {
	url := fmt.Sprintf(c.baseURL+"/boot/%s/rescue", ip)

	formData := neturl.Values{}
	formData.Set("os", input.OS)
	if input.Arch > 0 {
		formData.Set("arch", strconv.Itoa(input.Arch))
	}
	if len(input.AuthorizedKey) > 0 {
		formData.Set("authorized_key", input.AuthorizedKey)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var rescueResp models.RescueSetResponse
	err = json.Unmarshal(bytes, &rescueResp)
	if err != nil {
		return nil, err
	}

	return &rescueResp.Rescue, nil
}
