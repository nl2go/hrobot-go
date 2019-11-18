package client_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	client "github.com/nl2go/hrobot-go"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ClientSuite struct{}

var _ = Suite(&ClientSuite{})

const testServerIP = "123.123.123.123"
const testServerIP2 = "123.123.123.124"

const testIP = "123.123.123.123"
const testIP2 = "124.124.124.124"

func (s *ClientSuite) TestSetDefaultUserAgent(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqUserAgent := r.Header.Get("User-Agent")
		robotClient := client.NewBasicAuthClient("user", "pass")
		c.Assert(reqUserAgent, Equals, fmt.Sprintf("hrobot-client/%s", robotClient.GetVersion()))

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

	_, err := robotClient.ServerGetList()
	c.Assert(err, IsNil)
}

func (s *ClientSuite) TestSetCustomUserAgent(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqUserAgent := r.Header.Get("User-Agent")
		c.Assert(reqUserAgent, Equals, "hrobot-testsuite/0.0.1")

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
	robotClient.SetUserAgent("hrobot-testsuite/0.0.1")
	robotClient.SetBaseURL(ts.URL)

	_, err := robotClient.ServerGetList()
	c.Assert(err, IsNil)
}
