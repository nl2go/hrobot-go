# hrobot-go: A Go library for the Hetzner Robot Webservice

![Build](https://gitlab.com/newsletter2go/hrobot-go/badges/master/pipeline.svg) ![Coverage](https://gitlab.com/newsletter2go/hrobot-go/badges/master/coverage.svg)

Package hrobot-go is a library for the Hetzner Robot Webservice.

The library’s documentation is available at [GoDoc](https://godoc.org/github.com/nl2go/hrobot-go),
the public API documentation is available at [robot.your-server.de](https://robot.your-server.de/doc/webservice/en.html).

## Example

```go
package main

import (
    "fmt"
    "log"

    client "github.com/nl2go/hrobot-go"
)

func main() {
    robotClient := client.NewBasicAuthClient("user", "pass")

    servers, err := robotClient.ServerGetList()
    if err != nil {
        log.Fatalf("error while retrieving server list: %s\n", err)
    }

    fmt.Println(servers)
}
```
