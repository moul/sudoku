package sudoku

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
)

type Availables struct {
	Size    int
	Numbers map[int]bool
	Length  int
}

type Sudoku struct {
	Debug      bool
	Size       int
	SquareSize int
	Grid       [][]int
	Availables [][]Availables
	Groups     Groups
	BruteLimit int
	Missings   int
	Chars      string
	CharsSlice []string

	DoResolveNumbersThatAreOnlyInOnePosition bool
	DoResolveOnlyOne                         bool
}

type GroupType int

const (
	HorizontalGroup GroupType = iota
	VerticalGroup
	RegionGroup
)

type Groups map[GroupType][]Group

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

func (s *Sudoku) Clone(dest *Sudoku) {
	dest.Debug = s.Debug
	dest.Size = s.Size
	dest.SquareSize = s.SquareSize
	dest.BruteLimit = s.BruteLimit
	dest.Missings = s.Size * s.Size

	dest.DoResolveNumbersThatAreOnlyInOnePosition = s.DoResolveNumbersThatAreOnlyInOnePosition
	dest.DoResolveOnlyOne = s.DoResolveOnlyOne

	dest.initFields()
	dest.Groups = s.Groups

	for y := 0; y < s.Size; y++ {
		for x := 0; x < s.Size; x++ {
			if s.Grid[y][x] > 0 {
				dest.SetNumber(y, x, s.Grid[y][x])
			}
		}
	}
}

func (s *Sudoku) initFields() {
	s.Grid = make([][]int, s.Size)
	s.Availables = make([][]Availables, s.Size)
	s.Groups = make(map[GroupType][]Group, 0)
	for i := 0; i < s.Size; i++ {
		s.Grid[i] = make([]int, s.Size)
		s.Availables[i] = make([]Availables, s.Size)
		for j := 0; j < s.Size; j++ {
			s.Availables[i][j] = NewAvailables(s.Size)
		}
	}
}

func NewSudokuWithSize(sqsize int) Sudoku {
	size := sqsize * sqsize
	sudoku := Sudoku{
		Debug:      os.Getenv("DEBUG") == "1",
		SquareSize: sqsize,
		Size:       size,
		BruteLimit: 2,
		Missings:   size * size,

		DoResolveNumbersThatAreOnlyInOnePosition: true,
		DoResolveOnlyOne:                         true,
	}
	if sqsize == 5 {
		sudoku.CharsSlice = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y"}
	} else {
		sudoku.CharsSlice = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G"}
	}
	sudoku.Chars = strings.Join(sudoku.CharsSlice, "")

	sudoku.initFields()
	sudoku.initGroups()
	return sudoku
}

func (s *Sudoku) initGroups() {
	// horizontal groups
	for y := 0; y < s.Size; y++ {
		group := Group{}
		for x := 0; x < s.Size; x++ {
			group.Positions = append(group.Positions, Position{y, x})
		}
		s.Groups[HorizontalGroup] = append(s.Groups[HorizontalGroup], group)
	}

	// vertical groups
	for x := 0; x < s.Size; x++ {
		group := Group{}
		for y := 0; y < s.Size; y++ {
			group.Positions = append(group.Positions, Position{y, x})
		}
		s.Groups[VerticalGroup] = append(s.Groups[VerticalGroup], group)
	}
	// zone groups
	for a := 0; a < s.SquareSize; a++ {
		for b := 0; b < s.SquareSize; b++ {
			group := Group{}
			for c := 0; c < s.SquareSize; c++ {
				for d := 0; d < s.SquareSize; d++ {
					y := a*s.SquareSize + c
					x := b*s.SquareSize + d
					group.Positions = append(group.Positions, Position{y, x})
				}
			}
			s.Groups[RegionGroup] = append(s.Groups[RegionGroup], group)
		}
	}
}

func NewAvailables(size int) Availables {
	availables := Availables{
		Size:    size,
		Numbers: make(map[int]bool, size),
		Length:  size,
	}
	for i := 1; i <= size; i++ {
		availables.Numbers[i] = true
	}
	return availables
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
	a.Length = 1
}

func (a *Availables) Availables() []int {
	availables := []int{}
	for i := 1; i <= a.Size; i++ {
		if a.Numbers[i] {
			availables = append(availables, i)
		}
	}
	return availables
}

func (a *Availables) RemoveNumber(number int) bool {
	if a.Numbers[number] {
		a.Length--
		a.Numbers[number] = false
		return true
	}
	return false
}

func (g *Groups) AllGroups() []Group {
	groups := []Group{}
	for _, typeGroups := range *g {
		groups = append(groups, typeGroups...)
	}
	return groups
}

func (g *Groups) MatchCoords(y, x int) []Group {
	groups := []Group{}
	for _, group := range g.AllGroups() {
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
	s.Missings--
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
			if col == " " {
				continue
			}
			idx := strings.Index(s.Chars, col)
			if idx == -1 {
				return fmt.Errorf("Invalid character: %q", col)
			}
			s.SetNumber(y, x, idx+1)
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
			if col == 0 {
				line = append(line, " ")
			} else {
				line = append(line, s.CharsSlice[col-1])
			}
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

func (s *Sudoku) ResolveOnlyOne() int {
	changed := 0
	for y := 0; y < s.Size; y++ {
		for x := 0; x < s.Size; x++ {
			if s.Grid[y][x] != 0 {
				continue
			}
			if s.Availables[y][x].Length == 1 {
				s.SetNumber(y, x, s.Availables[y][x].Availables()[0])
				changed++
			}
		}
	}
	return changed
}

func (s *Sudoku) ResolveNumbersThatAreOnlyInOnePosition() int {
	changes := 0
	for _, group := range s.Groups.AllGroups() {
		for number := 1; number <= s.Size; number++ {
			count := 0
			for _, pos := range group.Positions {
				if s.Availables[pos.Y][pos.X].Numbers[number] {
					if count++; count > 1 {
						break
					}
				}
			}
			if count != 1 {
				continue
			}
			for _, pos := range group.Positions {
				if s.Availables[pos.Y][pos.X].Numbers[number] {
					s.SetNumber(pos.Y, pos.X, number)
					changes++
				}
			}
		}
	}
	return changes
}

func (s *Sudoku) ResolveNonBrute(depth int) {
	changes := 0
	iteration := 0
	kind := "start"

start:
	if s.Missings == 0 {
		return
	}
	if s.Debug {
		logrus.Infof("#######  depth=%-2d iteration=%-3d changes=%-2d missings=%d kind=%s\n%s\n%s", depth, iteration, changes, s.Missings, kind, s.String(), s.AvailablesString())
		logrus.Infof(strings.Repeat("#", 42))
	}
	iteration++

	if s.DoResolveOnlyOne {
		kind = "resolve-only-one"
		if changes = s.ResolveOnlyOne(); changes > 0 {
			goto start
		}
	}

	if s.DoResolveNumbersThatAreOnlyInOnePosition {
		kind = "resolve-numbers-that-are-only-in-one-position"
		if changes = s.ResolveNumbersThatAreOnlyInOnePosition(); changes > 0 {
			goto start
		}
	}
}

type Clone struct {
	Sudoku      *Sudoku
	BrutePos    Position
	BruteNumber int
}

func (s *Sudoku) ResolveRec(depth int) (*Sudoku, error) {
	s.ResolveNonBrute(depth)
	if s.Missings == 0 {
		return s, nil
	}

	if depth >= s.BruteLimit {
		return s, fmt.Errorf("Too deep")
	}

	var clones []Clone
	for y := 0; y < s.Size; y++ {
		for x := 0; x < s.Size; x++ {
			if s.Availables[y][x].Length == 0 {
				continue
			}
			for i := 0; i < s.Size; i++ {
				if !s.Availables[y][x].Numbers[i] {
					continue
				}
				clones = append(clones, Clone{
					Sudoku:      &Sudoku{},
					BrutePos:    Position{Y: y, X: x},
					BruteNumber: i,
				})
			}
		}
	}

	for _, clone := range clones {
		s.Clone(clone.Sudoku)
		clone.Sudoku.SetNumber(clone.BrutePos.Y, clone.BrutePos.X, clone.BruteNumber)
		newSudoku, err := clone.Sudoku.ResolveRec(depth + 1)
		if err != nil {
			s.Availables[clone.BrutePos.Y][clone.BrutePos.X].RemoveNumber(clone.BruteNumber)
			continue
		}
		if newSudoku.Missings == 0 {
			return newSudoku, nil
		}
	}

	return s, nil
}

func (s *Sudoku) Resolve() error {
	newSudoku, err := s.ResolveRec(0)
	if err != nil {
		return err
	}

	if newSudoku != s {
		newSudoku.Clone(s)
	}

	return nil
}
