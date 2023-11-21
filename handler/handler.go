// handler/handler.go
package handler

import (
	"fmt"

	"flexera.com/service"
)

// ProcessCSV reads the CSV file and converts it to JSON.
func ProcessCSV(csvFilePath string) error {
	csvData, err := service.ReadCSVFile(csvFilePath)
	if err != nil {
		return fmt.Errorf("error reading CSV file: %v", err)
	}

	jsonData, err := service.ConvertCSVToJSON(csvData)
	if err != nil {
		return fmt.Errorf("error converting CSV to JSON: %v", err)
	}

	count := service.GetMinimumCount(jsonData)
	fmt.Println("Minimum number of copies of application the company must purchase is:: ", count)

	return nil
}
