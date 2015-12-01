package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Usage:\t%s file [options] \n\n", os.Args[0])
	fmt.Printf("Options:\n%s\tOutput pretty JSON \n", "--pretty")
}

func hasArg(args []string, arg string) bool {
	for _, b := range args {
		if arg == b {
			return true
		}
	}
	return false
}
