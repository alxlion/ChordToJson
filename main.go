package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Line struct {
	Type    map[string]string
	Content string
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

// parseLine defines the type of the string following the ChordPro convention
// Using regexp to determine the type
// It returns the type of the string, its content without directives and any error encountered.
func parseLine(rawLine string) (map[string]string, string, error) {
	for key, value := range types {
		matchKey, err := regexp.MatchString(fmt.Sprint("^{(", key, ")"), rawLine)
		matchValue, err := regexp.MatchString(fmt.Sprint("^{(", value, ")"), rawLine)

		if matchKey {
			return map[string]string{key: value}, rawLine[1+len(key)+1 : len(rawLine)-1], err
		} else if matchValue {
			return map[string]string{key: value}, rawLine[1+len(value)+1 : len(rawLine)-1], err
		}
	}
	return map[string]string{"v": "verse"}, rawLine, nil // Simple line: verse
}

func main() {

	rawLines, err := readLines("./test.chordpro")
	var lineArray []Line
	if err != nil {
		panic(err)
	}

	for _, rawLine := range rawLines {
		if rawLine != "" {
			rawLineType, rawLineContent, _ := parseLine(rawLine)
			line := Line{rawLineType, rawLineContent}
			lineArray = append(lineArray, line)
		}
	}

	for _, line := range lineArray {
		fmt.Printf("%v: %v \n", line.Type, line.Content)
	}
}
