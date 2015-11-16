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
	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		logrus.Fatalf("Failed to read file %q: %v", os.Args[1], err)
	}
	input := strings.TrimSpace(string(buf))
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
