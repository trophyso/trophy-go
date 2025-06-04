# Go Library for Trophy

The Trophy Go client library provides convenient access to the Trophy API from applications written
in Go.

Trophy provides APIs and tools for adding gamification to your application, keeping users engaged
through rewards, achievements, streaks, and personalized communication.

## Installation

Make sure your project is using Go Modules (it will have a `go.mod` file in its root if it already
is):

```bash
go mod init
```

Then, reference trophy-go in a Go program with import:

```go
import "github.com/trophy-so/trophy-go"
```

Alternatively, you can use `go get`:

```bash
go get github.com/trophy-so/trophy-go
```

## Usage

```go
package main

import (
	"github.com/trophyso/trophy-go"
	trophyclient "github.com/trophyso/trophy-go/client"
	"github.com/trophyso/trophy-go/option"

	"fmt"
	"context"
)

func main() {
	client := trophyclient.NewClient(
		option.WithApiKey("YOUR_API_KEY"),
	)

	email := "jk.rowling@harrypotter.com"

	response, err := client.Metrics.Event(
		context.TODO(),
		"words-written",
		&api.MetricsEventRequest{
			User: &api.EventRequestUser{
				Id: "user-id",
				Email: &email,
			},
			Value: 750,
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}
```

## Documentation

See the [Trophy API Docs](https://docs.trophy.so) for more
information on the accessible endpoints.

## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.
