package client

import (
	"encoding/json"

	"github.com/nl2go/hrobot-go/models"
)

func (c *Client) KeyGetList() ([]models.Key, error) {
	url := c.baseURL + "/key"
	bytes, err := c.doGetRequest(url)
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
