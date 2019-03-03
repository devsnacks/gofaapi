package gofaapi

import "encoding/xml"

type pidRequest struct {
	XMLName xml.Name `xml:"getProductGroupIds"`
	Space   string   `xml:"space"` // hack: we need a space between the opening and ending tag
}

// GetProductGroupIds gets all available productgroups
func (c *Client) GetProductGroupIds() {
	r, _ := c.call("getProductGroupIds", &pidRequest{Space: " "})

	println(string(r))
}
