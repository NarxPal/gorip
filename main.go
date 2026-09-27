package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	searchTerm := os.Args[1]
	filename := os.Args[2]
	fmt.Printf("Received input: %s\n", searchTerm)

	fmt.Printf("Received input: %s\n", filename)

	file, opnErr := os.Open(filename)
	if opnErr != nil {
		fmt.Printf("Error opening file: %v\n", opnErr)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 1

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchTerm) {
			fmt.Printf("Line %v: %v \n", lineCount, line)
		}
		lineCount++
	}

	if scanErr := scanner.Err(); scanErr != nil {
		fmt.Printf("Error scanning file: %v \n", scanErr)
	}

}
