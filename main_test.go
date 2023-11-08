package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"verifier/src/models"
)

func TestPrintSummary(t *testing.T) {
	// Create a map with test data
	dataCounts := map[string]models.RecordData{
		"testTable": {
			RowCount:       5,
			CollectionName: "testCollection",
			DocCount:       5,
		},
	}

	// Capture the standard output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	PrintSummary(dataCounts)

	// Stop capturing the standard output
	w.Close()
	os.Stdout = old

	// Read the captured standard output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Check that the output contains the expected text
	expected := []string{"testTable", "5", "testCollection", "5", "YES"}
	for _, e := range expected {
		if !strings.Contains(output, e) {
			t.Errorf("Expected output to contain %q, got %q", e, output)
		}
	}
}
