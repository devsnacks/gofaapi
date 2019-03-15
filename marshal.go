package gofaapi

import (
	"encoding/xml"
	"strconv"
)

func (m AddressList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name = xml.Name{Space: "", Local: "addresses"}
	start.Attr = []xml.Attr{xml.Attr{Name: xml.Name{Local: "xsi:type"}, Value: "ns2:Map"}}
	e.EncodeToken(start)

	item := xml.StartElement{Name: xml.Name{Local: "item"}}
	value := xml.StartElement{Name: xml.Name{Local: "value"}}

	value.Attr = []xml.Attr{xml.Attr{Name: xml.Name{Local: "xsi:type"}, Value: "ns2:Map"}}

	e.EncodeToken(item)
	e.EncodeElement("sender", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeToken(value)

	MarshalAddress(e, m.SenderAddress)

	e.EncodeToken(value.End())
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("delivery", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeToken(value)

	MarshalAddress(e, m.DeliverAddress)

	e.EncodeToken(value.End())
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("invoice", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeToken(value)

	MarshalAddress(e, m.InvoiceAddress)

	e.EncodeToken(value.End())
	e.EncodeToken(item.End())

	return e.EncodeToken(start.End())
}

func MarshalAddress(e *xml.Encoder, address Address) {

	item := xml.StartElement{Name: xml.Name{Local: "item"}}

	e.EncodeToken(item)
	e.EncodeElement("customertype", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(address.Customertype, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("vatnumber", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(address.Vatnumber, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("taxnumber", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(address.Taxnumber, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("company", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(address.Company, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("gender", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(address.Gender, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("firstname", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(address.FirstName, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("lastname", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(address.LastName, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("address", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(address.Address, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("addressAdd", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(address.AddressAdd, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("postcode", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(address.Postcode, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("city", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(address.City, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())
}

func (m Options) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Attr = []xml.Attr{xml.Attr{Name: xml.Name{Local: "xsi:type"}, Value: "ns2:Map"}}
	e.EncodeToken(start)

	item := xml.StartElement{Name: xml.Name{Local: "item"}}
	for k, v := range m {
		e.EncodeToken(item)
		e.EncodeElement(strconv.Itoa(k), xml.StartElement{Name: xml.Name{Local: "key"}})
		e.EncodeElement(strconv.Itoa(v), xml.StartElement{Name: xml.Name{Local: "value"}})
		e.EncodeToken(item.End())
	}

	return e.EncodeToken(start.End())
}

func (m UploadInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name = xml.Name{Space: "", Local: "uploadData"}
	start.Attr = []xml.Attr{xml.Attr{Name: xml.Name{Local: "xsi:type"}, Value: "ns2:Map"}}

	e.EncodeToken(start)

	item := xml.StartElement{Name: xml.Name{Local: "item"}}

	e.EncodeToken(item)
	e.EncodeElement("dataTransferType", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(m.UploadType, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("dataTransferTime", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(m.Time, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("dataTransferText", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(m.Text, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())

	e.EncodeToken(item)
	e.EncodeElement("referenceText", xml.StartElement{Name: xml.Name{Local: "key"}})
	e.EncodeElement(m.Text, xml.StartElement{Name: xml.Name{Local: "value"}})
	e.EncodeToken(item.End())

	return e.EncodeToken(start.End())
}
