package main

import (
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/moul/sudoku"
)

func main() {
	sudo := sudoku.NewSudoku()
	input := `+-----------------+
|  3 5 6       7 8|
|8     9     6    |
|      4 7 8      |
|  5 4 3         1|
|1 6           5 7|
|3         1 2 9  |
|      8 4 5      |
|    3     6     9|
|  4       9 5 8  |
+-----------------+`
	if err := sudo.ParseString(input); err != nil {
		logrus.Fatalf("Failed to parse sudoku: %v", err)
	}
	fmt.Println(sudo.String())
	fmt.Println(sudo.AvailablesString())
	fmt.Printf("Missings: %d\n", sudo.Missings())

	fmt.Println(strings.Repeat("-", 40))

	if err := sudo.Resolv(); err != nil {
		logrus.Fatalf("Failed to resolve sudoku: %v", err)
	}
	fmt.Println(sudo.String())
	fmt.Println(sudo.AvailablesString())
	fmt.Printf("Missings: %d\n", sudo.Missings())
}
