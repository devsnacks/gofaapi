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

## Testing with Postman

To play around with the API using [postman](https://www.getpostman.com) you can import the folder [postman](./postman) to it.  
Currently it containts 2 requests and the environments for testing/live for each client.  
After importing you can give your credentials in the environment settings for global variables.  

You can then easily switch between live/testing environment for each client.