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
			fmt.Printf("%v:%v: %v \n", filename, lineCount, line)
		}
		lineCount++
	}

	if scanErr := scanner.Err(); scanErr != nil {
		fmt.Printf("Error scanning file: %v \n", scanErr)
	}

}

func searchDirectory(path string, searchTerm string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error read directory %v\n", err)
		return
	}
	for _, entry := range entries {
		fullpath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			searchDirectory(fullpath, searchTerm)
		} else {
			searchFile(fullpath, searchTerm)
		}
	}
}

func search(path string, searchTerm string) {
	entries, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Error from os .stat%v\n", err)
		return
	}
	if entries.IsDir() {
		// if path is a directory
		searchDirectory(path, searchTerm)
	} else {
		// else it's a file
		searchFile(path, searchTerm)
	}
}

func main() {
	if len(os.Args) < 3 {
		return
	}
	path := os.Args[2:]
	searchTerm := os.Args[1]
	for _, p := range path {
		search(p, searchTerm)
	}
	fmt.Printf("Received input: %s\n", searchTerm)
}
