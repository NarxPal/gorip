package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func searchFile(filename string, searchTerm string) {
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

func searchPath(path string, searchTerm string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error read directory %v\n", err)
		return
	}
	for _, entry := range entries {
		fullpath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			searchPath(fullpath, searchTerm)
		} else {
			searchFile(fullpath, searchTerm)
		}
	}
}

func listDirectory(path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error read directory %v\n", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println(entry.Name(), "DIRECTORY")
		} else {
			fmt.Println(entry.Name(), "FILE")
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		return
	}
	path := os.Args[2]
	// listDirectory(path)
	searchTerm := os.Args[1]
	filenames := os.Args[2:]
	searchPath(path, searchTerm)
	fmt.Printf("Received input: %s\n", searchTerm)

	fmt.Printf("Received input: %s\n", filenames)

	for _, filename := range filenames {
		searchFile(filename, searchTerm)

	}

}
