package csvtojl

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
)

func BellowAverageCSV2JL(inputFile string) ([]byte, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 1 {
		return nil, err
	}

	headers := records[0]
	var jsonData []byte
	for _, record := range records[1:] {
		jsonLine, err := convertToJSON(headers, record)
		if err != nil {
			return nil, err
		}
		jsonData = append(jsonData, jsonLine...)
	}
	return jsonData, nil
}

// Automatically detect the type and convert the CSV record to JSON
func convertToJSON(headers []string, record []string) ([]byte, error) {
	data := make(map[string]interface{})
	for i, header := range headers {
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
