package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGetSampleFromAoC(t *testing.T) {
	dataB, err := os.ReadFile("testdata/sample.txt")
	if err != nil {
		t.Fatalf("Reading sample test data: %v", err)
	}

	data := parseSampleFromPage(dataB)

	fmt.Printf("%s\n", data)

}
