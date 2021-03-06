# apachelog [![Build Status](https://travis-ci.org/vinxi/apachelog.png)](https://travis-ci.org/vinxi/apachelog) [![GoDoc](https://godoc.org/github.com/vinxi/apachelog?status.svg)](https://godoc.org/github.com/vinxi/apachelog) [![Coverage Status](https://coveralls.io/repos/github/vinxi/apachelog/badge.svg?branch=master)](https://coveralls.io/github/vinxi/apachelog?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/vinxi/apachelog)](https://goreportcard.com/report/github.com/vinxi/apachelog)

Apache style logging middleware for your proxies.
Uses [apachelog](https://github.com/h2non/apachelog).

## Installation

```bash
go get -u gopkg.in/vinxi/apachelog.v0
```

## API

See [godoc](https://godoc.org/github.com/vinxi/apachelog) reference.

## Example

#### Default log to stdout

```go
package main

import (
  "fmt"
  "gopkg.in/vinxi/apachelog.v0"
  "gopkg.in/vinxi/vinxi.v0"
)

const port = 3100

func main() {
  // Create a new vinxi proxy
  vs := vinxi.NewServer(vinxi.ServerOptions{Port: port})
  
  // Attach the log middleware 
  vs.Use(log.Default)
  
  // Target server to forward
  vs.Forward("http://httpbin.org")

  fmt.Printf("Server listening on port: %d\n", port)
  err := vs.Listen()
  if err != nil {
    fmt.Errorf("Error: %s\n", err)
  }
}
```

#### Using a custom io.Writer

```go
package main

import (
  "fmt"
  "os"
  "gopkg.in/vinxi/apachelog.v0"
  "gopkg.in/vinxi/vinxi.v0"
)

const port = 3100

func main() {
  // Create a new vinxi proxy
  vs := vinxi.NewServer(vinxi.ServerOptions{Port: port})
  
  // Attach the log middleware 
  vs.Use(log.New(os.Stdout))
  
  // Target server to forward
  vs.Forward("http://httpbin.org")

  fmt.Printf("Server listening on port: %d\n", port)
  err := vs.Listen()
  if err != nil {
    fmt.Errorf("Error: %s\n", err)
  }
}
```

## License

MIT
