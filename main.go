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

// Chord format reference
var types = map[string]string{
	"t":   "title",
	"st":  "subtitle",
	"soc": "start_of_chorus",
	"eoc": "end_of_chorus",
	"c":   "comment",
	"sot": "start_of_tab",
	"eot": "end_of_tab",
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

// getLineType defines the type of the string following the ChordPro convention
// Using regexp to determine the type
// It returns the type of the string, its plain text without directives and any error encountered.
func getLineType(line string) (string, string, error) {
	for key, value := range types {
		match, _ := regexp.MatchString(fmt.Sprint("^{(", value, "|", key, ")"), line)
		if match {
			return value, "", nil
		}
	}

	return "verse", "", nil
}

func main() {

	lines, err := readLines("./test.chordpro")
	var lineArray []Line
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		if line != "" {
			lineType, _, _ := getLineType(line)
			lineArray = append(lineArray, Line{"line", line})
			fmt.Printf("%v: %v \n", lineType, line)
		}
	}
}
