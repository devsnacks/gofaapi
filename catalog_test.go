package gofaapi

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func TestAddressMarshalling(t *testing.T) {
	address := AddressList{
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
			Address:      "hustenstr. 13",
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
			FirstName:    "Walter",
			LastName:     "Peterson",
			Address:      "hustenstr. 13",
			AddressAdd:   "stock 3",
			Postcode:     "98234",
			City:         "frankfurt",
			County:       "hessen",
			Locale:       "de",
			Phone:        "09312423423"}}

	f, _ := ioutil.ReadFile("testdata/addresses.xml")

	actualXML, _ := xml.MarshalIndent(address, "", "   ")

	if string(actualXML) != string(f) {
		t.Error("XMLs are not equal!")
	}
}

func TestOptionMarshalling(t *testing.T) {
	options := make(Options)
	options[100] = 1
	options[300] = 40
	options[430] = 123

	f, _ := ioutil.ReadFile("testdata/options.xml")

	actualXML, _ := xml.MarshalIndent(options, "", "   ")

	if string(actualXML) != string(f) {
		t.Error("XMLs are not equal!")
	}
}

func TestUploadInfoMashalling(t *testing.T) {
	uploadinfo := UploadInfo{UploadType: "upload",
		Time:          "01.03.2019 00:10:11",
		Text:          "Upload via Api.",
		ReferenceText: "Upload Reference 1"}

	f, _ := ioutil.ReadFile("testdata/uploadinfo.xml")

	actualXML, _ := xml.MarshalIndent(uploadinfo, "", "   ")

	if string(actualXML) != string(f) {
		t.Error("XMLs are not equal!")
	}
}
