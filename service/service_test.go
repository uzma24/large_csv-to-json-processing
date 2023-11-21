package service

import (
	"testing"
)

func TestConvertCSVToJSON(t *testing.T) {
	// Test case: Valid CSV data
	csvData := [][]string{
		{"Name", "Age", "City"},
		{"John", "30", "New York"},
		{"Alice", "25", "London"},
		{"Bob", "35", "San Francisco"},
	}

	jsonBytes, err := ConvertCSVToJSON(csvData)

	// Assert the expected JSON output
	expectedJSON := `[
		{"Name":"John","Age":30,"City":"New York"},
		{"Name":"Alice","Age":25,"City":"London"},
		{"Name":"Bob","Age":35,"City":"San Francisco"}
	]`
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(jsonBytes) != expectedJSON {
		t.Errorf("Expected JSON:\n%s\nActual JSON:\n%s", expectedJSON, jsonBytes)
	}

	// Test case: Invalid CSV data (missing headers)
	invalidCSVData := [][]string{
		{"John", "30", "New York"},
		{"Alice", "25", "London"},
		{"Bob", "35", "San Francisco"},
	}

	_, err = ConvertCSVToJSON(invalidCSVData)

	// Assert that an error occurred due to missing headers
	if err == nil {
		t.Error("Expected an error for invalid CSV data (missing headers), but got none.")
	}
}

func TestGetMinimumCount(t *testing.T) {
	// Test case: Valid JSON data
	jsonData := []byte(`[
		{"Cid": 1, "UserID": 1, "AppID": 374, "CType": "LAPTOP"},
		{"Cid": 2, "UserID": 2, "AppID": 374, "CType": "DESKTOP"},
		{"Cid": 2, "UserID": 2, "AppID": 374, "CType": "DESKTOP"}
	]`)

	count := GetMinimumCount(jsonData)

	// Assert the expected count
	expectedCount := int64(2)
	if count != expectedCount {
		t.Errorf("Expected count %d, but got %d", expectedCount, count)
	}

	// Test case: Invalid JSON data
	invalidJSON := []byte(`invalid json`)

	count = GetMinimumCount(invalidJSON)

	// Assert that an error occurred during unmarshalling
	if count != 0 {
		t.Errorf("Expected count 0 for invalid JSON, but got %d", count)
	}
}
