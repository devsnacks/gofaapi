[![CircleCI](https://circleci.com/gh/devsnacks/gofaapi.svg?style=shield)](https://circleci.com/gh/devsnacks/gofaapi) [![Go Report Card](https://goreportcard.com/badge/github.com/devsnacks/gofaapi)](https://goreportcard.com/report/github.com/devsnacks/gofaapi)
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

    //Get all available productgroups
    client.GetProductGroups()

    //prepare order
    options := make(gofaapi.Options)
    options[127] = 1

    uploadinfo := gofaapi.UploadInfo{UploadType: "upload",
        Time:          "01.03.2019 00:10:11",
        Text:          "Upload via Api.",
        ReferenceText: "Upload Reference 1"}

    addressList := gofaapi.AddressList{
        SenderAddress: Address{
            Customertype: "Private",
            Vatnumber:    "12345",
            Taxnumber:    "44444",
            Company:      "Pizza Domain",
            Gender:       "Unknown",
            FirstName:    "Stefano",
            LastName:     "Priebsch",
            Address:      "Musterstrasse. 13",
            AddressAdd:   "3. Stock",
            Postcode:     "98234",
            City:         "Priebschhausen",
            County:       "Hessen",
            Locale:       "de",
            Phone:        "09333233323332"},
        DeliverAddress: Address{
            Customertype: "Company",
            Vatnumber:    "12345",
            Taxnumber:    "44444",
            Company:      "flyeralarm",
            Gender:       "male",
            FirstName:    "Gustavo",
            LastName:     "Gans",
            Address:      "Flyerweg 13",
            AddressAdd:   "stock 3",
            Postcode:     "98234",
            City:         "frankfurt",
            County:       "hessen",
            Locale:       "de",
            Phone:        "09312423423"},
        InvoiceAddress: Address{
            Customertype: "Company",
            Vatnumber:    "12345",
            Taxnumber:    "44444",
            Company:      "flyeralarm",
            Gender:       "male",
            FirstName:    "Peter",
            LastName:     "Peterson",
            Address:      "Antonstr. 13",
            AddressAdd:   "stock 3",
            Postcode:     "98234",
            City:         "frankfurt",
            County:       "hessen",
            Locale:       "de",
            Phone:        "09312423423"}}

    order := gofaapi.Order{Username: "",
        Password:        "",
        QuantityID:      8131138,
        ShippingTypeID:  1,
        Options:         options,
        AddressList:     addressList,
        ShippingID:      1,
        AddressHandling: 1,
        PaymentID:       1,
        UploadInfo:      uploadinfo}

    //send the order
   client.SendOrder(order)
})
```

## Testing with Postman

To play around with the API using [postman](https://www.getpostman.com) you can import the folder [postman](./postman) to it.  
Currently it containts 2 requests and the environments for testing/live for each client.  
After importing you can give your credentials in the environment settings for global variables.  

You can then easily switch between live/testing environment for each client.
