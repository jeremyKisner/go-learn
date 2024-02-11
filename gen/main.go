package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filePath := "./gen/RESODataDictionary-2.0 - Fields.csv"
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("had error reading file", err)
		return
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	rows, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	headers := rows[0]
	var structBuffer bytes.Buffer

	// Start writing to the buffer
	structBuffer.WriteString("package main\n\n")
	structBuffer.WriteString("type Property struct {\n")

	for _, row := range rows[1:] {
		rowData := make(map[string]string)

		for i, value := range row {
			rowData[headers[i]] = value
		}
		if rowData["ResourceName"] == "Property" {
			var dtype string
			switch rowData["SimpleDataType"] {
			case "String", "String List, Single":
				dtype = "string"
			case "String List, Multi":
				dtype = "[]string"
			case "Number":
				dtype = "int"
			case "Boolean":
				dtype = "bool"
			case "Date", "Timestamp":
				dtype = "string"
			case "Resource":
				dtype = rowData["StandardName"]
			case "Collection":
				dtype = "[]interface{}"
			default:
				fmt.Println("could not find data type to convert", rowData["StandardName"], rowData["SimpleDataType"])
				continue
			}

			// Write field definition to the buffer
			fieldName := rowData["StandardName"]
			structBuffer.WriteString(fmt.Sprintf("\t%s %s `json:\"%s,omitempty\"`\n", fieldName, dtype, strings.ToLower(fieldName)))
		}
	}

	// End struct definition
	structBuffer.WriteString("}\n")

	// Write the buffer content to a Go file
	outputFilePath := "./generated_struct.go"
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(structBuffer.String())
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	fmt.Printf("Generated struct written to %s\n", outputFilePath)
}
