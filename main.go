package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Line struct {
	Type string
	Text string
}

// readLines gets content of a file using its path.
// Using bufio Scanner to append each line into an array.
// It returns the array with lines and any Scanner error encountered.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// getLineType defines the type of the string following the ChordPro convention (http://tenbyten.com/software/songsgen/help/HtmlHelp/files_reference.htm)
// Using regexp to determine the type
// It returns the type of the string, its plain text without directives and any regexp compilation error encountered.
func getLineType(line string) (string, string, error) {

	return "", "", nil
}

func main() {

	lines, err := readLines("./test.chordpro")
	var lineArray []Line
	if err != nil {
		panic(err)
	}

	for i, line := range lines {
		lineArray = append(lineArray, Line{"line", line})
		fmt.Printf("Line %v: %v \n", i+1, line)
	}
}
