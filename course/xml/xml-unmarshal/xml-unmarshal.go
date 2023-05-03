package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"
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
	var xmlBytes []byte
	var ship Boat
	fxml, err := os.Open("./test.xml")
	if err != nil {
		log.Fatal(err)
	}

	defer fxml.Close()

	scan := bufio.NewScanner(fxml)

	for scan.Scan() {
		xmlBytes = append(xmlBytes, scan.Bytes()...)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(fmt.Errorf("scanner %v", err))
	}
	if err := xml.Unmarshal(xmlBytes, &ship); err != nil {
		log.Fatal(fmt.Errorf("json.Unmarshal %v", err))
	}

	fmt.Println("xml-Crew", ship.Crew)
}
