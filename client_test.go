package client_test

import (
	"testing"

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
