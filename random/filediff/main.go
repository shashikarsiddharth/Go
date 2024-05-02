package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func loadFile(file string) (io.Reader, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	return f, nil
}

func parseFile(f io.Reader) map[string]int {
	index := make(map[string]int)
	fl := bufio.NewScanner(f)
	fl.Split(bufio.ScanLines)
	for fl.Scan() {
		index[fl.Text()]++
	}
	return index
}

func findMissingItem(inputFileIndex map[string]int, sourceFileIndex map[string]int) []string {
	var missingItems []string
	for k, _ := range sourceFileIndex {
		if _, ok := inputFileIndex[k]; !ok {
			missingItems = append(missingItems, k)
		}
	}
	return missingItems
}

func printOutput(result []string) {
	for i := range result {
		fmt.Println(result[i])
	}
	fmt.Printf("Total len: %v\n", len(result))
}

func main() {
	sourcePath := os.Getenv("SOURCE_PATH")
	testPath := os.Getenv("TEST_PATH")

	source, _ := loadFile(sourcePath)
	sourceFileIndex := parseFile(source)

	test, _ := loadFile(testPath)
	testFileIndex := parseFile(test)

	res := findMissingItem(testFileIndex, sourceFileIndex)
	printOutput(res)
}
