# go-socket.io

This repo is a fork of the [repo](https://github.com/xuzuxing/go-socket.io/) which fully supports the integration of Volmex websocket endpoints. See more information [here](https://docs.volmex.finance/index-data-api#websockets).

go-socket.io is library an implementation of [Socket.IO](http://socket.io) in Golang, which is a realtime application framework.

Current this library supports 1.4 version of the Socket.IO client. It supports room, namespaces and broadcast at now.

## Badges

![Build Status](https://github.com/volmexfinance/go-socket.io/workflows/CI/badge.svg)
[![GoDoc](http://godoc.org/github.com/volmexfinance/go-socket.io?status.svg)](http://godoc.org/github.com/volmexfinance/go-socket.io)
[![License](https://img.shields.io/github/license/golangci/golangci-lint)](/LICENSE)
[![Release](https://img.shields.io/github/release/volmexfinance/go-socket.io.svg)](https://github.com/volmexfinance/go-socket.io/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/volmexfinance/go-socket.io)](https://goreportcard.com/report/github.com/volmexfinance/go-socket.io)

## Contents

- [Install](#install)
- [Example](#example)
- [License](#license)

## Install

Install the package with:

```bash
go get github.com/volmexfinance/go-socket.io
```

Import it with:

```go
import "github.com/volmexfinance/go-socket.io"
```

and use `socketio` as the package name inside the code.

## Example

Please check more examples into folder in project for details. [Examples](https://github.com/volmexfinance/go-socket.io/tree/master/_examples)

```
package main

import (
	"context"
	"fmt"

	socketio "github.com/volmexfinance/go-socket.io"
	"github.com/volmexfinance/go-socket.io/engineio"
)

type DataPoint struct {
	Symbol    string
	Price     float64
	Timestamp int64
}

func main() {

	ctx := context.Background()

	dataPointsChannel := make(chan *DataPoint)

	client, err := socketio.NewClient("ws://ws.volmex.finance", &engineio.Options{})
	if err != nil {
		return
	}

	client.OnConnect(func(c socketio.Conn) error {

		fmt.Println("Connected")

		client.Emit("fetch-indices-messages")

		client.OnEvent("indices-messages-stream", func(s socketio.Conn, msg DataPoint) {
			dataPointsChannel <- &msg

		})

		return nil
	})

	err = client.Connect()
	if err != nil {
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		case dataPoint := <-dataPointsChannel:
			fmt.Printf("Received data point %s - %f - %d\n", dataPoint.Symbol, dataPoint.Price, dataPoint.Timestamp)
		}
	}

}
```


## License

The 3-clause BSD License  - see [LICENSE](https://opensource.org/licenses/BSD-3-Clause) for more details
