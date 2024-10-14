# BellowAverage (Chris Wang) Json Parser Module

**Introduction:** This repo includes a Json Parser Project that is used to practice handling Json files using Go. It also provides a publicly accessible API for reusing certain functions in this project.

## Json Parser Project

**Module:** [github.com/BellowAverage/JsonParser](https://github.com/BellowAverage/JsonParser)  
**Credit:** BellowAverage (Chris Wang) | 10.13.2024 | NU MSDS-431 Assignment 3  
**Description:** This is a utility module containing two commands: `csvtojl` and `jsparse`.

---

### csvtojl

**Description:**  
`csvtojl` is a utility for converting CSV files to JSON lines (JL). It provides two ways to parse CSV headers and their corresponding data types:

1. **Manual:** Define headers and data types manually in the Go script, which is safer but more time-consuming.
2. **Automatic (using the `-a` option):** Automatically detect headers and data types, which is faster but less reliable.

**Usage:**

```bash
./cmd/csvtojl/csvtojl [options] <input.csv> <output.jl>
```

**Options:**

- `-a` — Automatically detect headers and data types.  
- `-help` — Display available commands and options.

**Example:**

```bash
./cmd/csvtojl/csvtojl -a ./data/housesInput.csv housesOutputAuto.jl
./cmd/csvtojl/csvtojl -a ./data/random.csv randomAuto.jl

./cmd/csvtojl/csvtojl ./data/housesInput.csv housesOutput.jl
```

---

### jsparse

**Description:**  
`jsparse` parses a JSON file either using pre-defined structs or a mapping (dictionary). This command demonstrates the two methods available for parsing JSON in Go.

**Usage:**

```bash
./cmd/jsparse/jsparse [options] ./data/exampleJson.json
```

**Options:**

- `-m` — Parse JSON using mapping (dictionary).  
- `-s` — Parse JSON using pre-set structs.

**Example:**

```bash
./cmd/jsparse/jsparse -m ./data/exampleJson.json
./cmd/jsparse/jsparse -s ./data/exampleJson.json
```

---


## BellowAverageCSV2JL API

`csvtojl` is a Go module that converts CSV files into JSON Lines format by automatically detecting the CSV headers and data types. This module allows you to easily process CSV files and output each row as a JSON object, with headers as keys.

### Features
- Automatically detect CSV headers and corresponding data types (int, float, string).
- Convert CSV files to JSON Lines format (`.jl`).
- Easy-to-use API for integrating into other Go projects.

### Installation

To install the module, use `go get`:

```bash
go get github.com/BellowAverage/JsonParser/API
```

### Usage

#### 1. Example CSV File (`input.csv`)

Create a CSV file to use for conversion:

```csv
Product,Price,Quantity,Rating,Available
Laptop,999.99,10,4.5,true
Mouse,19.99,100,4.8,true
Keyboard,49.99,50,4.6,false
Monitor,149.99,25,4.3,true
Headphones,89.99,40,4.7,true
```

#### 2. Example Go Code

Here’s a basic example of how to use the `csvtojl.BellowAverageCSV2JL` function to convert a CSV file into a JSON Lines file:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BellowAverage/JsonParser/API" // Ensure correct import path
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
```

#### 3. Running the Program

1. Save the above Go code in a file called `main.go`.
2. Place your `input.csv` file in the same directory as `main.go`.
3. Run the program:

```bash
go run main.go
```

After running, the `output.jl` file will be created, containing the JSON Lines equivalent of the CSV data.

#### 4. Expected Output (`output.jl`)

For the example CSV provided, the output will look like this:

```json
{"Product":"Laptop","Price":999.99,"Quantity":10,"Rating":4.5,"Available":true}
{"Product":"Mouse","Price":19.99,"Quantity":100,"Rating":4.8,"Available":true}
{"Product":"Keyboard","Price":49.99,"Quantity":50,"Rating":4.6,"Available":false}
{"Product":"Monitor","Price":149.99,"Quantity":25,"Rating":4.3,"Available":true}
{"Product":"Headphones","Price":89.99,"Quantity":40,"Rating":4.7,"Available":true}
```

### Function Description

#### `csvtojl.BellowAverageCSV2JL`

This function takes a CSV file path as input and converts its contents into JSON Lines format. It automatically detects the headers and data types (int, float, string) for each field in the CSV.

##### Parameters:
- `inputFile (string)`: The path to the CSV file to be converted.

##### Returns:
- `([]byte, error)`: A byte slice containing the JSON data and an error if something goes wrong.

#### Example of Function Usage in a Go Program

```go
jsonData, err := csvtojl.BellowAverageCSV2JL("input.csv")
if err != nil {
    log.Fatalf("Failed to convert CSV to JSON: %v", err)
}

fmt.Println(string(jsonData))
```

### License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.