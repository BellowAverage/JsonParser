package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type HouseInput struct {
	Value      float32 `json:"value"`
	Income     float32 `json:"income"`
	Age        int32   `json:"age"`
	Rooms      int32   `json:"rooms"`
	Bedrooms   int32   `json:"bedrooms"`
	Population int32   `json:"pop"`
	Households int32   `json:"hh"`
}

func main() {

	fmt.Println()

	fmt.Println("Module: github.com/BellowAverage/JsonParser")
	fmt.Println("Credit: BellowAverage (Chris Wang) | 10.13.2024 | NU MSDS-431 Assignment 3")
	fmt.Println("Usage: ./csvtojl -[option] <input.csv> <output.jl>")

	fmt.Println()

	// -a: detect header's csv's data types automatically
	// -help: display available commands and options

	auto := flag.Bool("a", false, "Automatically detect header and data types")
	help := flag.Bool("help", false, "Display available commands and options")
	flag.Parse()

	args := flag.Args()

	if *help {
		printHelp()
		return
	}

	if len(args) != 2 {
		fmt.Println("Error: 2 arguments are expected.")
		fmt.Println("Run 'csvtojl -help' for more information.")
		fmt.Println()
		os.Exit(1)
	}

	inputFile := args[0]
	outputFile := args[1]

	// Opening CSV file for reading

	log.Println("Opening CSV file for reading...")

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	// Parsing CSV file

	log.Println("Parsing CSV file...")

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse CSV file: %v", err)
	}

	if len(records) < 1 {
		log.Fatal("No records found in the CSV file.")
	}

	// Opening output JSON lines file for writing

	log.Println("Opening output JSON lines file for writing...")

	outFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer outFile.Close()

	headers := records[0]

	log.Println("Headers detected in CSV file:", headers)

	// Parse CSV records and convert to JSON lines
	if *auto {
		log.Println("Automatic data type detection enabled. Converting records...")
		for _, record := range records[1:] {
			if len(record) != len(headers) {
				log.Fatalf("Invalid CSV row: expected %d fields, got %d", len(headers), len(record))
			}

			jsonLine, err := convertToJSON(headers, record)
			if err != nil {
				log.Fatalf("Failed to parse CSV row: %v", err)
			}

			_, err = outFile.Write(append(jsonLine, '\n'))
			if err != nil {
				log.Fatalf("Failed to write to output file: %v", err)
			}
		}
	} else {
		log.Println("Using predefined data type struct (HouseInput). Converting records...")
		for _, record := range records[1:] {
			if len(record) != 7 {
				log.Printf("Warning: Invalid CSV row: expected 7 fields, got %d. Skipping row.", len(record))
				continue
			}

			var house HouseInput
			err := parseCSVRow(record, &house)
			if err != nil {
				log.Printf("Warning: Failed to parse CSV row: %v. Skipping row.", err)
				continue
			}

			jsonLine, err := json.Marshal(house)
			if err != nil {
				log.Fatalf("Failed to serialize to JSON: %v", err)
			}

			_, err = outFile.Write(append(jsonLine, '\n'))
			if err != nil {
				log.Fatalf("Failed to write to output file: %v", err)
			}
		}
	}

	log.Println("Conversion completed successfully.")
}

func printHelp() {
	fmt.Println("csvtojl - A utility for converting CSV files to JSON lines.")
	fmt.Println("\nUsage:")
	fmt.Println("  csvtojl [options] <input_csv> <output_json_lines>")
	fmt.Println("\nOptions:")
	fmt.Println("  -a       Automatically detect header and data types.")
	fmt.Println("  -help    Display available commands and options.")
	fmt.Println("\nExample:")
	fmt.Println("  csvtojl -a housesInput.csv housesOutput.jl")
}

func convertToJSON(headers []string, record []string) ([]byte, error) {
	data := make(map[string]interface{})
	for i, header := range headers {
		// Attempt to parse value to best guess type
		if intVal, err := strconv.Atoi(record[i]); err == nil {
			data[header] = intVal
		} else if floatVal, err := strconv.ParseFloat(record[i], 64); err == nil {
			data[header] = floatVal
		} else {
			data[header] = record[i]
		}
	}
	return json.Marshal(data)
}

func parseCSVRow(record []string, house *HouseInput) error {
	var err error

	valueFloat64, err := strconv.ParseFloat(record[0], 32)
	if err != nil {
		return err
	}
	house.Value = float32(valueFloat64)

	incomeFloat64, err := strconv.ParseFloat(record[1], 32)
	if err != nil {
		return err
	}
	house.Income = float32(incomeFloat64)

	ageInt, err := strconv.Atoi(record[2])
	if err != nil {
		return err
	}
	house.Age = int32(ageInt)

	roomsInt, err := strconv.Atoi(record[3])
	if err != nil {
		return err
	}
	house.Rooms = int32(roomsInt)

	bedroomsInt, err := strconv.Atoi(record[4])
	if err != nil {
		return err
	}
	house.Bedrooms = int32(bedroomsInt)

	populationInt, err := strconv.Atoi(record[5])
	if err != nil {
		return err
	}
	house.Population = int32(populationInt)

	householdsInt, err := strconv.Atoi(record[6])
	if err != nil {
		return err
	}
	house.Households = int32(householdsInt)

	return nil
}
