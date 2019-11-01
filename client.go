package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"strings"

	"gitlab.com/newsletter2go/hrobot-go/models"
)

const baseURL string = "https://robot-ws.your-server.de"
const version = "1.0.0"
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
