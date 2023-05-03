package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	// small file
	csvDataSmall := [][]string{
		{"Name", "gender", "level", "hobbies"},
		{"Tere", "female", "20", "eat,sleep,sing"},
		{"Victor", "male", "10", "play,tv,sport"},
		{"Rodrigo", "male", "5", "sleep,run,jump"},
	}

	smallCSV, err := os.Create("./smallest.csv")
	if err != nil {
		log.Fatal(fmt.Errorf("create smallest file %v", err))
	}

	defer smallCSV.Close()

	csvWriterSmall := csv.NewWriter(smallCSV)
	csvWriterSmall.Comma = ';'

	csvWriterSmall.WriteAll(csvDataSmall)

	if err := csvWriterSmall.Error(); err != nil {
		log.Fatal(fmt.Errorf("writer smallest file %v", err))
	}

	// lager file
	csvDataLarge := [][]string{
		{"Name", "gender", "level", "hobbies"},
		{"Tere", "female", "20", "eat,sleep,sing"},
		{"Victor", "male", "10", "play,tv,sport"},
		{"Rodrigo", "male", "5", "sleep,run,jump"},
		{"Tere", "female", "20", "eat,sleep,sing"},
		{"Victor", "male", "10", "play,tv,sport"},
		{"Rodrigo", "male", "5", "sleep,run,jump"},
	}

	largeCSV, err := os.Create("./largest.csv")
	if err != nil {
		log.Fatal(fmt.Errorf("create largest file %v", err))
	}

	defer largeCSV.Close()

	csvWriterLarge := csv.NewWriter(largeCSV)
	csvWriterLarge.Comma = '\t'

	for _, row := range csvDataLarge {
		if err := csvWriterLarge.Write(row); err != nil {
			log.Fatal(fmt.Errorf("write largest file %v", err))
		}
	}

	csvWriterLarge.Flush()

	if err := csvWriterLarge.Error(); err != nil {
		log.Fatal(fmt.Errorf("writer largest file %v", err))
	}
}
