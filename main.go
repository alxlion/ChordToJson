/* The MIT License (MIT)

Copyright (c) 2015 Alexandre Lion <contact@alexandrelion.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE. */

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
)

type Line struct {
	Type    map[string]string
	Content string
	Chords  map[int]string
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

// parseLine defines the type of the string following the ChordPro convention.
// Using regexp to determine the type.
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

// getChords gets position of every chords in the line.
// Using regexp to determine the positions.
// It returns the map of chords and their position.
func getChords(rawLine string) (map[int]string, error) {
	mappedChords := make(map[int]string)
	r, err := regexp.Compile("\\[[A-Za-z0-9#\\/]+\\]")
	matchChordIndex := r.FindAllStringIndex(rawLine, -1)
	matchChord := r.FindAllString(rawLine, -1)

	for i, position := range matchChordIndex {
		mappedChords[position[0]] = string(matchChord[i])
	}

	return mappedChords, err
}

// truncChords remove chord from line.
// It browses the line char by char to exclude chords.
// It returns the cleaned line.
func truncChords(rawLine string) string {
	var buf bytes.Buffer
	var insertMode bool = false

	for i := 0; i < len(rawLine); i++ {
		if string(rawLine[i]) == "[" {
			insertMode = true
		} else if string(rawLine[i]) == "]" {
			insertMode = false
		}

		if insertMode == false {
			if string(rawLine[i]) != "]" {
				buf.WriteString(string(rawLine[i]))
			}
		}
	}

	return buf.String()
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
			line := Line{rawLineType, rawLineContent, map[int]string{}}
			lineArray = append(lineArray, line)
			fmt.Println(getChords(rawLine))
		}
	}

	for _, line := range lineArray {
		fmt.Printf("%v: %v \n", line.Type, line.Content)
	}

}
