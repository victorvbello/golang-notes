package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./test.csv")
	if err != nil {
		log.Fatal(fmt.Errorf("[ERROR] open file %v", err))
	}
	defer f.Close()

	var fileBuf bytes.Buffer
	fileTeeReader := io.TeeReader(f, &fileBuf) // remember this read and write at the same time

	// for smaller files
	csvReaderSmall := csv.NewReader(fileTeeReader)
	csvReaderSmall.Comment = '#'
	csvReaderSmall.Comma = ';' // default is ','

	recordsSmall, err := csvReaderSmall.ReadAll()
	if err != nil {
		log.Println(fmt.Errorf("[ERROR] csv readAll %v", err))
	}

	fmt.Println("CSV small - ", recordsSmall)

	// for larger files
	csvReader := csv.NewReader(&fileBuf)
	csvReader.Comment = '#'
	csvReader.Comma = ';' // default is ','

	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			if csvError, ok := err.(*csv.ParseError); ok {
				log.Printf("[ERROR] line: %v, err: %v \n", csvError.Line, csvError.Err)
				if csvError.Err == csv.ErrFieldCount {
					continue
				}
			}
			log.Fatal(fmt.Errorf("[ERROR] csv read %v", err))
		}
		fmt.Println("CSV - large", record)

	}

}
