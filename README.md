# gofaapi

FLYERALARM Reseller API library for golang

## Install

```go
go get github.com/devsnacks/gofaapi
```

## Usage

```go
package main

import (
    "github.com/devsnacks/gofaapi"
)

func main() {
    opt := &gofaapi.ClientOptions{
        URL:      "...de/shop/soap/",
        Username: "",
        Password: ""}

    client := gofaapi.NewClient(opt)

    client.GetProductGroupIds()
})
```