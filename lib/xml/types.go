package xml

import "encoding/xml"

type Address struct {
	Line1   string
	City    string
	State   string
	ZipCode string
	Country string
}
type Person struct {
	XMLName   xml.Name  `xml:"person"`
	Id        int       `xml:"id,attr"`
	FirstName string    `xml:"name>first"`
	LastName  string    `xml:"name>last"`
	Age       int       `xml:"age"`
	Height    int       `xml:"height"`
	Married   bool      `xml:"married,omitempty"`
	Addresses []Address `xml:"addresses>address"`
	Comment   string    `xml:",comment"`
}
