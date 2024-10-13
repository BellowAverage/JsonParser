## BellowAverage (Chris Wang) Json Parser

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
