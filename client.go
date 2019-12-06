package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const baseURL string = "https://robot-ws.your-server.de"
const version = "0.1.3"
const userAgent = "hrobot-client/" + version

type Client struct {
	Username  string
	Password  string
	baseURL   string
	userAgent string
}

func NewBasicAuthClient(username, password string) RobotClient {
	return &Client{
		Username:  username,
		Password:  password,
		baseURL:   baseURL,
		userAgent: userAgent,
	}
}

func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) GetVersion() string {
	return version
}

func (c *Client) doGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (c *Client) doPostFormRequest(url string, formData url.Values) ([]byte, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("User-Agent", c.userAgent)
	req.SetBasicAuth(c.Username, c.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
