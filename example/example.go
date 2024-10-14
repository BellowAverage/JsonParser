package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BellowAverage/JsonParser/public"
)

func main() {
	// Replace "input.csv" with the path to your CSV file
	jsonData, err := csvtojl.BellowAverageCSV2JL("input.csv")
	if err != nil {
		log.Fatalf("Failed to convert CSV to JSON: %v", err)
	}

	// Specify the output file path (JSON Lines format)
	outputFile := "output.jl"

	// Open the file for writing (create if it doesn't exist)
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.WriteString(string(jsonData))
	if err != nil {
		log.Fatalf("Failed to write to output file: %v", err)
	}

	fmt.Printf("Successfully written to %s\n", outputFile)
}
