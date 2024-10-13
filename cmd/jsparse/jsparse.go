package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func MappingJsonParser(fileContent string) map[string]interface{} {
	var data map[string]interface{}

	err_parse := json.Unmarshal([]byte(fileContent), &data)
	if err_parse != nil {
		fmt.Printf("Error parsing JSON string - %s", err_parse)
	}

	return data
}

type Nested struct {
	IsIt        bool   `json:"isit"`
	Description string `json:"description"`
}

type exampleJson struct {
	Name    string `json:"name"`
	Numbers []int  `json:"numbers"`
	Nested  Nested `json:"nested"`
}

func main() {

	fmt.Println()

	fmt.Println("Module: github.com/BellowAverage/JsonParser")
	fmt.Println("Credit: BellowAverage (Chris Wang) | 10.13.2024 | NU MSDS-431 Assignment 3")
	fmt.Println("Usage: ./cmd/jsparse/jsparse -[option] ./data/exampleJson.json")

	fmt.Println()

	// Define command line flags
	mappingOption := flag.Bool("m", false, "Use mapping to parse JSON")
	structOption := flag.Bool("s", false, "Use struct to parse JSON")
	flag.Parse()

	// Validate input
	if len(flag.Args()) < 1 {
		log.Fatalf("Usage: jlparse -[option] input.jl")
	}

	inputFile := flag.Arg(0)
	fileContent, err_io := os.ReadFile(inputFile)
	if err_io != nil {
		log.Fatalf("Error reading file - %s", err_io)
	}

	// Check which option is selected
	if *mappingOption {
		fmt.Println("Parsing JSON using Mapping: ")
		fmt.Println("----------------------------")
		data := MappingJsonParser(string(fileContent))

		// Print "name" value
		fmt.Printf("Name is %s\n", data["name"].(string))

		// Print "numbers" (handle as float64 array and print each element)
		numbers := data["numbers"].([]interface{})
		for i, num := range numbers {
			fmt.Printf("Number %d is %.0f\n", i, num.(float64))
		}

		// Print "nested" description value
		nestedDescription := data["nested"].(map[string]interface{})["description"].(string)
		fmt.Printf("Nested description is %s\n", nestedDescription)
	} else if *structOption {
		fmt.Println("Parsing JSON using Structs: ")
		fmt.Println("----------------------------")
		var exampleJson_data exampleJson

		err_exampleJson := json.Unmarshal(fileContent, &exampleJson_data)
		if err_exampleJson != nil {
			log.Fatalf("Error parsing JSON - %s", err_exampleJson)
		}

		// Access and print the fields
		fmt.Printf("Name is %s\n", exampleJson_data.Name)
		fmt.Printf("Numbers are %v\n", exampleJson_data.Numbers)
		fmt.Printf("Nested description is %s\n", exampleJson_data.Nested.Description)
	} else {
		log.Fatalf("Please specify an option: -m for mapping or -s for struct")
	}
}
