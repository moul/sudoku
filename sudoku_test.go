package sudoku

import "fmt"

func ExampleSudoku_Resolve() {
	sudoku := NewSudoku()
	sudoku.BruteLimit = 0
	sudoku.ParseString(`
+-----------------+
|    1 6          |
|  9   2         8|
|7 5              |
|5         3   6  |
|3     4   8     7|
|  4   7         9|
|              3 4|
|4         1   2  |
|          9 5    |
+-----------------+
`)
	sudoku.Resolve()
	fmt.Println(sudoku.String())
	// Output:
	// +-----------------+
	// |8 3 1 6 9 7 2 4 5|
	// |6 9 4 2 3 5 7 1 8|
	// |7 5 2 1 8 4 6 9 3|
	// |5 7 8 9 1 3 4 6 2|
	// |3 2 9 4 6 8 1 5 7|
	// |1 4 6 7 5 2 3 8 9|
	// |9 1 7 5 2 6 8 3 4|
	// |4 8 5 3 7 1 9 2 6|
	// |2 6 3 8 4 9 5 7 1|
	// +-----------------+
}

func ExampleSudoku_Resolve_brute() {
	sudoku := NewSudoku()
	sudoku.BruteLimit = 1
	sudoku.ParseString(`
+-----------------+
|          4   7 1|
|8     2 1        |
|    7   9   3    |
|            4 2 6|
|2               7|
|6 5 9            |
|    5   6   1    |
|        4 9     5|
|4 1   3          |
+-----------------+
`)
	sudoku.Resolve()
	fmt.Println(sudoku.String())
	// Output:
	// +-----------------+
	// |5 9 2 8 3 4 6 7 1|
	// |8 6 3 2 1 7 5 4 9|
	// |1 4 7 5 9 6 3 8 2|
	// |7 8 1 9 5 3 4 2 6|
	// |2 3 4 6 8 1 9 5 7|
	// |6 5 9 4 7 2 8 1 3|
	// |9 2 5 7 6 8 1 3 4|
	// |3 7 8 1 4 9 2 6 5|
	// |4 1 6 3 2 5 7 9 8|
	// +-----------------+
}

func ExampleSudoku_Resolve_hardest1() {
	sudoku := NewSudoku()
	sudoku.BruteLimit = 2
	sudoku.ParseString(`
+-----------------+
|1         7   9  |
|  3     2       8|
|    9 6     5    |
|    5 3     9    |
|  1     8       2|
|6         4      |
|3             1  |
|  4             7|
|    7       3    |
+-----------------+
`)
	sudoku.Resolve()
	fmt.Println(sudoku.String())
	// Output:
	// +-----------------+
	// |1 6 2 8 5 7 4 9 3|
	// |5 3 4 1 2 9 6 7 8|
	// |7 8 9 6 4 3 5 2 1|
	// |4 7 5 3 1 2 9 8 6|
	// |9 1 3 5 8 6 7 4 2|
	// |6 2 8 7 9 4 1 3 5|
	// |3 5 6 4 7 8 2 1 9|
	// |2 4 1 9 3 5 8 6 7|
	// |8 9 7 2 6 1 3 5 4|
	// +-----------------+
}

/*
func ExampleSudoku_Resolve_hardest2() {
	sudoku := NewSudoku()
	sudoku.BruteLimit = 3
	sudoku.ParseString(`
+-----------------+
|8                |
|    3 6          |
|  7     9   2    |
|  5       7      |
|        4 5 7    |
|      1       3  |
|    1         6 8|
|    8 5       1  |
|  9         4    |
+-----------------+
`)
	sudoku.Resolve()
	fmt.Println(sudoku.String())
	// Output:
	// +-----------------+
	// |1 6 2 8 5 7 4 9 3|
	// |5 3 4 1 2 9 6 7 8|
	// |7 8 9 6 4 3 5 2 1|
	// |4 7 5 3 1 2 9 8 6|
	// |9 1 3 5 8 6 7 4 2|
	// |6 2 8 7 9 4 1 3 5|
	// |3 5 6 4 7 8 2 1 9|
	// |2 4 1 9 3 5 8 6 7|
	// |8 9 7 2 6 1 3 5 4|
	// +-----------------+
}
*/
