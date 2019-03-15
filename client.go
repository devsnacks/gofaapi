package gofaapi

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

// Envelope envelope
type envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	XMLNS   string   `xml:"xmlns:xsi,attr"`
	XMLNS2  string   `xml:"xmlns:ns2,attr"`
	Body    body
}

// Body body
type body struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Content interface{} `xml:",omitempty"`
}

// Client Soap
type Client struct {
	http *http.Client
	url  string
}

type loginRequest struct {
	XMLName  xml.Name `xml:"login"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
}

// ClientOptions to configure the client
type ClientOptions struct {
	URL      string
	Username string
	Password string
}

// NewClient return SOAP client
func NewClient(opt *ClientOptions) *Client {
	jar, _ := cookiejar.New(nil)

	tr := &http.Transport{}
	client := &http.Client{Jar: jar, Transport: tr}
	c := &Client{url: opt.URL, http: client}

	login(opt.Username, opt.Password, c)

	return c
}

func (c *Client) call(soapAction string, request interface{}) (response []byte, err error) {

	envelope := envelope{
		XMLNS:  "http://www.w3.org/2001/XMLSchema-instance",
		XMLNS2: "http://xml.apache.org/xml-soap",
		Body: body{
			Content: request,
		},
	}

	buffer := new(bytes.Buffer)
	buffer.WriteString(xml.Header)

	env, _ := xml.MarshalIndent(envelope, " ", " ")
	buffer.WriteString(string(env))

	req, err := http.NewRequest("POST", c.url, buffer)
	if err != nil {
		err = fmt.Errorf("failed to create POST request: %s", err.Error())
		return
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Set("SOAPAction", soapAction)
	req.Header.Set("Content-Length", string(buffer.Len()))
	req.Header.Set("User-Agent", "gofaapi client")

	req.Close = true

	res, err := c.http.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to send SOAP request: %s", err.Error())
		return
	}

	defer res.Body.Close()

	response, err = ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("failed to read SOAP body: %s", err.Error())
		return
	}

	return
}

func login(u string, p string, client *Client) {
	client.call("login", &loginRequest{
		Username: u, Password: p})
}
