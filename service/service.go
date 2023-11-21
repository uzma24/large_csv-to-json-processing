package service

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"

	"flexera.com/model"
)

var headers []string // global headers variable

const numWorkers = 4 // Number of goroutines to run concurrently

// ReadCSVFile reads a CSV file and returns a 2D slice of strings representing the CSV data.
func ReadCSVFile(csvFilePath string) ([][]string, error) {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// ConvertCSVToJSON converts CSV data (2D slice) to JSON.
func ConvertCSVToJSON(csvData [][]string) ([]byte, error) {
	headers := csvData[0]
	var records []map[string]interface{}

	for _, record := range csvData[1:] {
		data := make(map[string]interface{})
		for i, value := range record {
			if intValue, err := strconv.Atoi(value); err == nil {
				data[headers[i]] = intValue
			} else {
				data[headers[i]] = value
			}
		}
		records = append(records, data)
	}

	// Convert the JSON data to a byte slice
	jsonBytes, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func worker(wg *sync.WaitGroup, in <-chan []string, out chan<- map[string]interface{}, headers []string) {
	defer wg.Done()

	for record := range in {
		data := make(map[string]interface{})
		for i, value := range record {
			if intValue, err := strconv.Atoi(value); err == nil {
				data[headers[i]] = intValue
			} else {
				data[headers[i]] = value
			}
		}
		out <- data
	}
}

func ConvertToJSONParallel(records [][]string) ([]byte, error) {
	fmt.Println("within JSON parallel function")

	if len(records) == 0 {
		return nil, fmt.Errorf("empty records slice")
	}

	var jsonRecords []map[string]interface{}
	recordChannel := make(chan []string, numWorkers)
	resultChannel := make(chan map[string]interface{}, numWorkers)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		localHeaders := make([]string, len(headers))
		copy(localHeaders, headers)
		go worker(&wg, recordChannel, resultChannel, localHeaders)
	}

	// Start a goroutine to close the result channel when all workers are done
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	for _, record := range records {
		recordChannel <- record
	}

	close(recordChannel)

	// Collect results from workers
	for result := range resultChannel {
		jsonRecords = append(jsonRecords, result)
	}

	// Convert the JSON records to a byte slice
	jsonBytes, err := json.MarshalIndent(jsonRecords, "", "  ")
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

// getMinimum count
func GetMinimumCount(jsonData []byte) int64 {
	var rows []model.Row
	if err := json.Unmarshal(jsonData, &rows); err != nil {
		fmt.Println("Error unmarshalling jsonData:", err)
		return 0
	}

	dataSet := make(map[int]map[string]int)
	visited := make(map[int]bool)
	count := 0
	for _, row := range rows {
		if visited[row.Cid] {
			continue
		}
		if dataSet[row.UserID] == nil {
			dataSet[row.UserID] = map[string]int{
				"laptopCount":  0,
				"desktopCount": 0,
			}
		}

		if row.CType == "DESKTOP" || row.CType == "Desktop" || row.CType == "desktop" {
			dataSet[row.UserID]["desktopCount"]++
		} else {
			dataSet[row.UserID]["laptopCount"]++
		}

		visited[row.Cid] = true
	}

	for _, v := range dataSet {
		laptopCount := v["laptopCount"]
		desktopCount := v["desktopCount"]

		if laptopCount >= desktopCount {
			count += laptopCount
		} else {
			count += desktopCount
		}
	}

	return int64(count)
}
