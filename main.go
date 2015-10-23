package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {

	lines, err := readLines("./test.chordpro")
	if err != nil {
		panic(err)
	}

	for i, line := range lines {
		fmt.Printf("Line %v: %v \n", i+1, line)
	}
}
