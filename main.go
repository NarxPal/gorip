package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func searchFile(filename string, searchTerm string, caseInsensitive *bool, lineNum *bool) {
	file, opnErr := os.Open(filename)
	if opnErr != nil {
		fmt.Printf("Error opening file: %v\n", opnErr)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lineCount := 1

	pattern := searchTerm
	if *caseInsensitive {
		pattern = "(?i)" + pattern // case-insensitive  mode on
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("invalid regex: %v\n", err)
		return
	}

	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			if *lineNum {
				fmt.Printf("%v:%v: %v \n", filename, lineCount, line)
			} else {
				fmt.Printf("%v: %v \n", filename, line)
			}
		}
		lineCount++
	}

	if scanErr := scanner.Err(); scanErr != nil {
		fmt.Printf("Error scanning file: %v \n", scanErr)
	}

}

func searchDirectory(path string, searchTerm string, caseInsensitive *bool, lineNum *bool) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error read directory %v\n", err)
		return
	}
	for _, entry := range entries {
		fullpath := filepath.Join(path, entry.Name())
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		if entry.IsDir() {
			searchDirectory(fullpath, searchTerm, caseInsensitive, lineNum)
		} else {
			searchFile(fullpath, searchTerm, caseInsensitive, lineNum)
		}
	}
}

func search(path string, searchTerm string, caseInsensitive *bool, lineNum *bool) {
	entries, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Error from os .stat %v\n", err)
		return
	}
	if entries.IsDir() {
		// if path is a directory
		searchDirectory(path, searchTerm, caseInsensitive, lineNum)
	} else {
		// else it's a file
		searchFile(path, searchTerm, caseInsensitive, lineNum)
	}
}

func main() {
	caseInsensitive := flag.Bool("i", false, "case-insensitive")
	lineNum := flag.Bool("n", false, "line-number")
	flag.Parse()
	args := flag.Args() // usin flags args and not os.args since we are using flag "-i" for case-insensitivity

	if len(args) < 2 {
		return
	}
	path := args[1:]
	searchTerm := args[0]
	for _, p := range path {
		search(p, searchTerm, caseInsensitive, lineNum)
	}
}
