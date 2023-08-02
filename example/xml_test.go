package example

import (
	"encoding/xml"
	"fmt"
	"github.com/beevik/etree"
	xmlUtil "github.com/wegoteam/wepkg/io/xml"
	"log"
	"os"
	"strings"
	"testing"
)

type Animal int
type Size int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)
const (
	Unrecognized Size = iota
	Small
	Large
)

func (a *Animal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}

	return nil
}

func (a Animal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var s string
	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}
	return e.EncodeElement(s, start)
}

func (s *Size) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	default:
		*s = Unrecognized
	case "small":
		*s = Small
	case "large":
		*s = Large
	}
	return nil
}

func (s Size) MarshalText() ([]byte, error) {
	var name string
	switch s {
	default:
		name = "unrecognized"
	case Small:
		name = "small"
	case Large:
		name = "large"
	}
	return []byte(name), nil
}

func TestName(t *testing.T) {
	blob := `
	<animals>
		<animal>gopher</animal>
		<animal>armadillo</animal>
		<animal>zebra</animal>
		<animal>unknown</animal>
		<animal>gopher</animal>
		<animal>bee</animal>
		<animal>gopher</animal>
		<animal>zebra</animal>
	</animals>`
	var zoo struct {
		Animals []Animal `xml:"animal"`
	}
	if err := xml.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
	}

	census := make(map[Animal]int)
	for _, animal := range zoo.Animals {
		census[animal] += 1
	}

	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras:  %d\n* Unknown: %d\n",
		census[Gopher], census[Zebra], census[Unknown])

	blob2 := `
	<sizes>
		<size>small</size>
		<size>regular</size>
		<size>large</size>
		<size>unrecognized</size>
		<size>small</size>
		<size>normal</size>
		<size>small</size>
		<size>large</size>
	</sizes>`
	var inventory struct {
		Sizes []Size `xml:"size"`
	}
	if err := xml.Unmarshal([]byte(blob2), &inventory); err != nil {
		log.Fatal(err)
	}

	counts := make(map[Size]int)
	for _, size := range inventory.Sizes {
		counts[size] += 1
	}

	fmt.Printf("Inventory Counts:\n* Small:        %d\n* Large:        %d\n* Unrecognized: %d\n",
		counts[Small], counts[Large], counts[Unrecognized])

}

func TestXML(t *testing.T) {
	type User struct {
		Name string `json:"name" xml:"name"`
	}
	var user User
	user.Name = "test"
	marshal, err := xmlUtil.Marshal(user)
	if err != nil {
		fmt.Errorf("xml.Marshal err: %v", err)
	}
	fmt.Println(marshal)

	err = xmlUtil.Unmarshal(marshal, &user)
	if err != nil {
		fmt.Errorf("xml.Unmarshal err: %v", err)
	}
	fmt.Println(user)
}
func TestDocument(t *testing.T) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	doc.CreateProcInst("xml-stylesheet", `type="text/xsl" href="style.xsl"`)

	people := doc.CreateElement("People")
	people.CreateComment("These are all known people")

	jon := people.CreateElement("Person")
	jon.CreateAttr("name", "Jon")

	sally := people.CreateElement("Person")
	sally.CreateAttr("name", "Sally")

	doc.Indent(2)
	doc.WriteTo(os.Stdout)
}
