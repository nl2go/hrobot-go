package client_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	. "gopkg.in/check.v1"

	"github.com/nl2go/hrobot-go"
)

func (s *ClientSuite) TestIPListSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/ip_list.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	ips, err := robotClient.IPGetList()
	c.Assert(err, IsNil)
	c.Assert(len(ips), Equals, 2)
	c.Assert(ips[0].IP, Equals, testIP)
	c.Assert(ips[1].IP, Equals, testIP2)
	c.Assert(ips[0].ServerIP, Equals, testServerIP)
	c.Assert(ips[1].ServerIP, Equals, testServerIP)
}
