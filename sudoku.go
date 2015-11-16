package sudoku

import (
	"fmt"
	"strconv"
	"strings"
)

type Sudoku struct {
	Size int
	Grid [][]int
}

func NewSudoku() *Sudoku {
	return NewSudokuWithSize(9)
}

func NewSudokuWithSize(size int) *Sudoku {
	sudoku := Sudoku{
		Size: size,
		Grid: make([][]int, size),
	}
	for i := 0; i < size; i++ {
		sudoku.Grid[i] = make([]int, size)
	}
	return &sudoku
}

func (s *Sudoku) ParseString(input string) error {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	for y, line := range lines[1 : s.Size+1] {
		for x := 0; x < s.Size; x++ {
			col := line[1+x*2 : 1+x*2+1]
			if col != " " {
				colNb, err := strconv.Atoi(col)
				if err != nil {
					return err
				}
				s.Grid[y][x] = colNb
			}
		}
	}
	return nil
}

func (s *Sudoku) String() string {
	lines := []string{}
	lines = append(lines, fmt.Sprintf("+%s+", strings.Repeat("-", s.Size*2-1)))

	for _, gridLine := range s.Grid {
		line := []string{}
		for _, col := range gridLine {
			line = append(line, strconv.Itoa(col))
		}
		lineStr := fmt.Sprintf("|%s|", strings.Join(line, " "))
		lineStr = strings.Replace(lineStr, "0", " ", -1)
		lines = append(lines, lineStr)
	}
	lines = append(lines, fmt.Sprintf("+%s+", strings.Repeat("-", s.Size*2-1)))
	return strings.Join(lines, "\n")
}
