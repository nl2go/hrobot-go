package client_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	. "gopkg.in/check.v1"

	client "gitlab.com/newsletter2go/hrobot-go"
	"gitlab.com/newsletter2go/hrobot-go/models"
)

func (s *ClientSuite) TestBootRescueGetInactiveSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/boot_rescue_get_inactive.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	rescue, err := robotClient.BootRescueGet(testServerIP)
	c.Assert(err, IsNil)
	c.Assert(rescue.ServerIP, Equals, testServerIP)
}

func (s *ClientSuite) TestBootRescueGetActiveSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/boot_rescue_get_active.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	rescue, err := robotClient.BootRescueGet(testServerIP)
	c.Assert(err, IsNil)
	c.Assert(rescue.ServerIP, Equals, testServerIP)
}

func (s *ClientSuite) TestBootRescueSetSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		c.Assert(reqContentType, Equals, "application/x-www-form-urlencoded")

		body, bodyErr := ioutil.ReadAll(r.Body)
		c.Assert(bodyErr, IsNil)
		c.Assert(string(body), Equals, "arch=64&os=linux")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/boot_rescue_set.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	input := &models.RescueSetInput{
		OS:   "linux",
		Arch: 64,
	}

	rescue, err := robotClient.BootRescueSet(testServerIP, input)
	c.Assert(err, IsNil)
	c.Assert(rescue.ServerIP, Equals, testServerIP)
}
