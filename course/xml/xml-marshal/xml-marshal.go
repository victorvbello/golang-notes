package main

import (
	"encoding/xml"
	"fmt"
)

type Crew struct {
	ID    int      `xml:"id,omitempty"`
	Name  string   `xml:"name,attr"`
	Range string   `xml:"range,omitempty"`
	Tasks []string `xml:"task>item-task"`
}

type Boat struct {
	XMLName xml.Name `xml:"ship"`
	id      int
	Captain Crew   `xml:"captain"`
	Code    string `xml:"ship-details>serial-code"`
	Type    string `xml:"ship-details>ship-type"`
	Crew    []Crew `xml:"souls>soul"`
}

func main() {
	captainNemo := Crew{
		Name:  "Nemo",
		Range: "Captain",
		Tasks: []string{
			"Ship driver",
			"Define routes",
			"Binding Marriages",
		},
	}

	nautilus := Boat{
		id:      1,
		Captain: captainNemo,
		Code:    "F0R3V3R-H4PPY",
		Type:    "Oasis class",
		Crew: []Crew{
			captainNemo,
			{ID: 100, Name: "Victor"},
			{ID: 101, Name: "Rodrigo"},
			{ID: 102, Name: "Tere"},
			{Name: "homeless"},
		},
	}

	b, err := xml.MarshalIndent(nautilus, " ", "  ")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(xml.Header, string(b))

}
