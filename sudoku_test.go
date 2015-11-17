package sudoku

import "fmt"

func ExampleSudoku_Resolve() {
	sudoku := NewSudoku()
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
