package client_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	client "github.com/nl2go/hrobot-go"
	"github.com/nl2go/hrobot-go/models"
	. "gopkg.in/check.v1"
)

func (s *ClientSuite) TestServerGetListSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/server_list.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	servers, err := robotClient.ServerGetList()
	c.Assert(err, IsNil)
	c.Assert(len(servers), Equals, 2)
	c.Assert(servers[0].ServerIP, Equals, testServerIP)
	c.Assert(servers[1].ServerIP, Equals, testServerIP2)
}

func (s *ClientSuite) TestServerGetSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/server_get.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	server, err := robotClient.ServerGet(testServerIP)
	c.Assert(err, IsNil)
	c.Assert(server.ServerIP, Equals, testServerIP)
}

func (s *ClientSuite) TestServerGetNotFound(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/server_get_404.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	_, err := robotClient.ServerGet(testServerIP)
	c.Assert(err, NotNil)
}

func (s *ClientSuite) TestServerSetNameSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		c.Assert(reqContentType, Equals, "application/x-www-form-urlencoded")

		body, bodyErr := ioutil.ReadAll(r.Body)
		c.Assert(bodyErr, IsNil)
		c.Assert(string(body), Equals, "server_name=mongodb-prod-px62-nvme-hetzner-nbg1-dc1-123456")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/server_get.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	input := &models.ServerSetNameInput{
		Name: "mongodb-prod-px62-nvme-hetzner-nbg1-dc1-123456",
	}

	server, err := robotClient.ServerSetName(testServerIP, input)
	c.Assert(err, IsNil)
	c.Assert(server.ServerIP, Equals, testServerIP)
}

func (s *ClientSuite) TestServerReverseSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/server_reverse.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	cancellation, err := robotClient.ServerReverse(testServerIP)
	c.Assert(err, IsNil)
	c.Assert(cancellation.ServerIP, Equals, testServerIP)
	c.Assert(cancellation.CancellationDate, Equals, "2014-04-15")
}
