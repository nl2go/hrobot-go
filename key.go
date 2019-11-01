package client

import (
	"encoding/json"
	"net/http"

	"gitlab.com/newsletter2go/hrobot-go/models"
)

func (c *Client) KeyGetList() ([]models.Key, error) {
	url := c.baseURL + "/key"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var keys []models.KeyResponse
	err = json.Unmarshal(bytes, &keys)
	if err != nil {
		return nil, err
	}

	var data []models.Key
	for _, key := range keys {
		data = append(data, key.Key)
	}

	return data, nil
}
