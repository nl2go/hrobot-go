package client

import (
	"encoding/json"
	"fmt"
	neturl "net/url"
	"strconv"

	"github.com/nl2go/hrobot-go/models"
)

func (c *Client) BootRescueGet(ip string) (*models.Rescue, error) {
	url := fmt.Sprintf(c.baseURL+"/boot/%s/rescue", ip)
	bytes, err := c.doGetRequest(url)
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

func (c *Client) BootRescueSet(ip string, input *models.RescueSetInput) (*models.Rescue, error) {
	url := fmt.Sprintf(c.baseURL+"/boot/%s/rescue", ip)

	formData := neturl.Values{}
	formData.Set("os", input.OS)
	if input.Arch > 0 {
		formData.Set("arch", strconv.Itoa(input.Arch))
	}
	if len(input.AuthorizedKey) > 0 {
		formData.Set("authorized_key", input.AuthorizedKey)
	}

	bytes, err := c.doPostFormRequest(url, formData)
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
