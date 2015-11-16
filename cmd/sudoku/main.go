package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/moul/sudoku"
)

func main() {
	sudo := sudoku.NewSudoku()

	if len(os.Args) < 2 {
		logrus.Fatalf("Usage: sudoku /path/to/map-file.txt")
	}

	filepath := os.Args[1]
	var input string
	if filepath == "-" {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			logrus.Fatalf("Failed to read from stdin: %v", err)
		}
		input = string(bytes)
	} else {
		buf, err := ioutil.ReadFile(filepath)
		if err != nil {
			logrus.Fatalf("Failed to read file %q: %v", filepath, err)
		}
		input = strings.TrimSpace(string(buf))
	}
	input = input[1 : len(input)-1]
	sudokus := strings.Split(input, "+\n+")
	for _, sudokuStr := range sudokus {
		sudokuStr = fmt.Sprintf("+%s+", sudokuStr)

		if err := sudo.ParseString(sudokuStr); err != nil {
			logrus.Fatalf("Failed to parse sudoku: %v", err)
		}

		if err := sudo.Resolve(); err != nil {
			logrus.Fatalf("Failed to resolve sudoku: %v", err)
		}

		fmt.Println(sudo.String())
		if sudo.Missings() > 0 {
			fmt.Println(sudo.AvailablesString())
			fmt.Printf("Missings: %d\n", sudo.Missings())
		}
	}
}
