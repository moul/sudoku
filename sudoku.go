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

	DoResolveNumbersThatAreOnlyInOnePosition    bool
	DoResolveNumbersThatCanOnlyBeInFewPositions bool
	DoResolveOnlyOne                            bool
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
	dest.DoResolveNumbersThatCanOnlyBeInFewPositions = s.DoResolveNumbersThatCanOnlyBeInFewPositions
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

		DoResolveNumbersThatAreOnlyInOnePosition:    true,
		DoResolveNumbersThatCanOnlyBeInFewPositions: false,
		DoResolveOnlyOne:                            true,
	}
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

func (s *Sudoku) RemoveNumbersThatCanOnlyBeInFewPositions() int {
	changes := 0
	for _, group := range s.Groups.AllGroups() {
		for idxA, posA := range group.Positions {
			identicalSlots := 0
			availableA := s.Availables[posA.Y][posA.X]
			for idxB, posB := range group.Positions {
				if idxA == idxB {
					continue
				}
				availableB := s.Availables[posB.Y][posB.X]
				if availableA.String() == availableB.String() {
					identicalSlots++
				}
			}
			if identicalSlots != availableA.Length-1 {
				continue
			}
			for idxB, posB := range group.Positions {
				if idxA == idxB {
					continue
				}
				availableB := s.Availables[posB.Y][posB.X]
				if availableA.String() != availableB.String() {
					for _, number := range availableA.Availables() {
						if s.Availables[posB.Y][posB.X].RemoveNumber(number) {
							changes++
						}
					}
				}
			}
		}
	}
	return changes
}

func (s *Sudoku) ResolveNumbersThatAreOnlyInOnePosition() int {
	changes := 0
	for _, group := range s.Groups.AllGroups() {
		for number := 1; number <= s.Size; number++ {
			count := 0
			for _, pos := range group.Positions {
				if s.Availables[pos.Y][pos.X].Numbers[number] {
					count++
				}
			}
			if count == 1 {
				for _, pos := range group.Positions {
					if s.Availables[pos.Y][pos.X].Numbers[number] {
						s.SetNumber(pos.Y, pos.X, number)
						changes++
					}
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

	if s.DoResolveNumbersThatCanOnlyBeInFewPositions && depth == 0 {
		kind = "resolve-numbers-that-can-only-be-in-few-positions"
		if changes = s.RemoveNumbersThatCanOnlyBeInFewPositions(); changes > 0 {
			goto start
		}
	}
}

func (s *Sudoku) ResolveRec(depth int) (*Sudoku, error) {
	s.ResolveNonBrute(depth)
	if s.Missings == 0 {
		return s, nil
	}

	if depth >= s.BruteLimit {
		return s, fmt.Errorf("Too deep")
	}

	for y := 0; y < s.Size; y++ {
		for x := 0; x < s.Size; x++ {
			if s.Availables[y][x].Length > 0 {
				clone := Sudoku{}
				s.Clone(&clone)
				clone.SetNumber(y, x, s.Availables[y][x].Availables()[0])
				newSudoku, err := clone.ResolveRec(depth + 1)
				if err != nil {
					continue
				}
				if newSudoku.Missings == 0 {
					return newSudoku, nil
				}
			}
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
