package gofaapi

import (
	"encoding/xml"
)

type Order struct {
	XMLName         xml.Name `xml:"sendFullOrder"`
	Username        string   `xml:"username"`
	Password        string   `xml:"password"`
	QuantityID      int      `xml:"quantityId"`
	ShippingTypeID  int      `xml:"shippingTypeId"`
	Options         Options
	AddressList     AddressList
	ShippingID      int `xml:"shippingId"`
	AddressHandling int `xml:"isStandard"`
	UpgradeID       int `xml:"upgradeId,omitempty"`
	PaymentID       int `xml:"paymentId"`
	UploadInfo      UploadInfo
	ResellerPrice   string `xml:"resellerGross,omitempty"`
	Width           string `xml:"width,omitempty"`
	Height          string `xml:"height,omitempty"`
}

type Options map[int]int

type UploadInfo struct {
	UploadType    string
	Time          string
	Text          string
	ReferenceText string
}

type AddressList struct {
	DeliverAddress Address
	SenderAddress  Address
	InvoiceAddress Address
}
type Address struct {
	Customertype string
	Vatnumber    string
	Taxnumber    string
	Company      string
	Gender       string
	FirstName    string
	LastName     string
	Address      string
	AddressAdd   string
	Postcode     string
	City         string
	County       string
	Locale       string
	Phone        string
}

type requestGetPGroups struct {
	XMLName xml.Name `xml:"getProductGroupIds"`
	Space   string   `xml:"space"` // hack: we need a space between the opening and ending tag
}
type requestGetAttr struct {
	XMLName  xml.Name `xml:"getProductAttributesByGroupId"`
	PGroupID int      `xml:"PGroupID"`
}
type requestGetQuantityID struct {
	XMLName xml.Name `xml:"findProductByQuantityId"`
	QID     int      `xml:"QID"`
}
type requestAddProductToCart struct {
	XMLName xml.Name `xml:"addProductToCart"`
	QID     int      `xml:"QID"`
}

// GetProductGroups gets all available productgroups
func (c *Client) GetProductGroups() []byte {
	r, _ := c.call("getProductGroupIds", &requestGetPGroups{Space: " "})

	return r
}

// GetProductAttributes get all attributes of a productgroup
func (c *Client) GetProductAttributes(PGroupID int) []byte {
	r, _ := c.call("getProductAttributesByGroupId", &requestGetAttr{})

	return r
}

// GetProduct get a Product by QuantityID
func (c *Client) GetProduct(PGroupID int) []byte {
	r, _ := c.call("findProductByQuantityId", &requestGetQuantityID{})

	return r
}

// AddProductToCart adds a product to the cart
func (c *Client) AddProductToCart(PGroupID int) []byte {
	r, _ := c.call("addProductToCart", &requestAddProductToCart{})

	return r
}

// SendOrder sends a fully configured order
func (c *Client) SendOrder(order Order) []byte {
	r, _ := c.call("sendFullOrder", &order)

	return r
}
