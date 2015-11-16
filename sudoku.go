package sudoku

import (
	"fmt"
	"strconv"
	"strings"
)

type Availables struct {
	Size    int
	Numbers map[int]bool
}

type Sudoku struct {
	Size       int
	SquareSize int
	Grid       [][]int
	Availables [][]Availables
	Groups     Groups
}
type Groups []Group

type Position struct {
	Y int
	X int
}

type Group struct {
	Size      int
	Positions []Position
}

func NewSudoku() Sudoku {
	return NewSudokuWithSize(3)
}

func NewAvailables(size int) Availables {
	availables := Availables{
		Size:    size,
		Numbers: make(map[int]bool, size),
	}
	for i := 1; i <= size; i++ {
		availables.Numbers[i] = true
	}
	return availables
}

func NewSudokuWithSize(sqsize int) Sudoku {
	size := sqsize * sqsize
	sudoku := Sudoku{
		SquareSize: sqsize,
		Size:       size,
		Grid:       make([][]int, size),
		Availables: make([][]Availables, size),
		Groups:     make([]Group, 0),
	}
	for i := 0; i < size; i++ {
		sudoku.Grid[i] = make([]int, size)
		sudoku.Availables[i] = make([]Availables, size)
		for j := 0; j < size; j++ {
			sudoku.Availables[i][j] = NewAvailables(size)
		}
	}

	// horizontal groups
	for y := 0; y < size; y++ {
		group := Group{}
		for x := 0; x < size; x++ {
			group.Positions = append(group.Positions, Position{y, x})
		}
		sudoku.Groups = append(sudoku.Groups, group)
	}

	// vertical groups
	for x := 0; x < size; x++ {
		group := Group{}
		for y := 0; y < size; y++ {
			group.Positions = append(group.Positions, Position{y, x})
		}
		sudoku.Groups = append(sudoku.Groups, group)
	}
	// zone groups
	for a := 0; a < sudoku.SquareSize; a++ {
		for b := 0; b < sudoku.SquareSize; b++ {
			group := Group{}
			for c := 0; c < sudoku.SquareSize; c++ {
				for d := 0; d < sudoku.SquareSize; d++ {
					y := a*sudoku.SquareSize + c
					x := b*sudoku.SquareSize + d
					group.Positions = append(group.Positions, Position{y, x})
				}
			}
			sudoku.Groups = append(sudoku.Groups, group)
		}
	}

	return sudoku
}

func (a *Availables) String() string {
	output := ""
	for i := 1; i <= a.Size; i++ {
		if a.Numbers[i] {
			output += strconv.Itoa(i)
		} else {
			output += "."
		}
	}
	return output
}

func (a *Availables) SetNumber(number int) {
	for i := 1; i <= a.Size; i++ {
		a.Numbers[i] = i == number
	}
}

func (a *Availables) RemoveNumber(number int) {
	a.Numbers[number] = false
}

func (g *Groups) MatchCoords(y, x int) []Group {
	groups := []Group{}
	for _, group := range *g {
		for _, pos := range group.Positions {
			if pos.Y == y && pos.X == x {
				groups = append(groups, group)
				break
			}
		}
	}
	return groups
}

func (s *Sudoku) SetNumber(y, x, number int) {
	s.Grid[y][x] = number
	s.Availables[y][x].SetNumber(number)
	for _, group := range s.Groups.MatchCoords(y, x) {
		for _, pos := range group.Positions {
			s.Availables[pos.Y][pos.X].RemoveNumber(number)
		}
	}
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
				s.SetNumber(y, x, colNb)
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

func (s *Sudoku) AvailablesString() string {
	lines := []string{}
	lines = append(lines, fmt.Sprintf("+%s+", strings.Repeat("-", s.Size*(s.Size+1)-1)))
	for _, availablesLine := range s.Availables {
		line := []string{}
		for _, availables := range availablesLine {
			line = append(line, availables.String())
		}
		lineStr := fmt.Sprintf("|%s|", strings.Join(line, " "))
		lines = append(lines, lineStr)
	}
	lines = append(lines, fmt.Sprintf("+%s+", strings.Repeat("-", s.Size*(s.Size+1)-1)))
	return strings.Join(lines, "\n")
}

func (s *Sudoku) Missings() int {
	missings := 0
	for _, line := range s.Grid {
		for _, col := range line {
			if col == 0 {
				missings++
			}
		}
	}
	return missings
}

func (s *Sudoku) Resolv() error {
	return nil
}
