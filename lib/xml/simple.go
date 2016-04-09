package xml

import (
	"encoding/xml"
	"fmt"
	"os"
	//"strconv"
)

func TestWrite() {
	var person = &Person{
		Id:        1,
		FirstName: "David",
		LastName:  "Mumladze",
		Age:       35}
	person.Comment = "Comment block"
	person.Addresses = append(person.Addresses, Address{
		Line1:   "1 Main St",
		City:    "Morristown",
		State:   "NJ",
		ZipCode: "07960",
		Country: "USA"})
	person.Addresses = append(person.Addresses, Address{
		Line1:   "1 Crappy Rd",
		City:    "New York",
		State:   "NY",
		ZipCode: "11245",
		Country: "USA"})
	var encoder = xml.NewEncoder(os.Stdout)
	encoder.Indent("", "   ")
	if err := encoder.Encode(person); err != nil {
		fmt.Println(err)
	}
	fmt.Println("")
	var arr = []*Person{person}
	var marshal, _ = xml.MarshalIndent(arr, "", "   ")
	fmt.Println(string(marshal))
}

var (
	count = 0
)

func counter() int {
	count = count + 1
	return count
}
