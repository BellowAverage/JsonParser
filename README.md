### BellowAverage (Chris Wang) Json Parser

Module: github.com/BellowAverage/JsonParser
Credit: BellowAverage (Chris Wang) | 10.13.2024 | NU MSDS-431 Assignment 3
Description: This is a utility module that contains 2 commands (csvtojl and jsparse).

* csvtojl
Description:
csvtojl is a utility for converting CSV files to JSON lines.
This command provides 2 ways of parsing CSV's headers and their data types.
  1. Define them in the go script mannully, which is safe but time-consuming.
  2. (using -a option) Detect headers and data types automatically, which is fast but unsafe.

Usage:
  ./cmd/csvtojl/csvtojl [options] <input.csv> <output.jl>

Options:
  -a       Automatically detect header and data types.
  -help    Display available commands and options.

Example:
  ./cmd/csvtojl/csvtojl -a ./data/housesInput.csv housesOutput.jl

* jsparse
Description:
jsparse parses JSON file through pre-defined structs or using mapping (dictionary).
This command simply tests the 2 methods used to parse JSON in go.

Usage:
  ./cmd/jsparse/jsparse [options] ./data/exampleJson.json

Options:
  -m       Parse JSON using mapping.
  -s       Parse JSON using pre-set structs.
