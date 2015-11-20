package sudoku

import "fmt"

func ExampleSudoku_Resolve_2x2() {
	sudoku := NewSudokuWithSize(2)
	sudoku.BruteLimit = 4
	sudoku.ParseString(`
+-------+
|3      |
|       |
|    1  |
|4   3  |
+-------+
`)
	sudoku.Resolve()
	fmt.Println(sudoku.String())
	// Output:
	// +-------+
	// |3 2 4 1|
	// |1 4 2 3|
	// |2 3 1 4|
	// |4 1 3 2|
	// +-------+
}

func ExampleSudoku_Resolve_4x4() {
	sudoku := NewSudokuWithSize(4)
	sudoku.BruteLimit = 0
	sudoku.ParseString(`
+-------------------------------+
|5 9   2       F 3   4 7     G E|
|      4 E   2 3 B   8   D 9 F  |
|D B   E 7     G       9   A 3 5|
|          D 9   F A   1   4    |
|4   5     8 F   6   1 A 3     G|
|    A   B E   5     D 3 1   6  |
|  2 7 8             G   B   9  |
|  6 E B 9 G C 1 5   2   A      |
|  8 9 F C 5 3 4 7   A   2      |
|  3 D 7             6   9   5  |
|    4   D A   9     F B G   7  |
|B   2     1 7   E   9 5 F     4|
|          9 G   A 8   6   F    |
|8 E   5 4     D       2   B A 9|
|      G A   B E 4   7   5 8 D  |
|9 A   6       2 D   B F     C 1|
+-------------------------------|

`)
	sudoku.Resolve()
	fmt.Println(sudoku.String())
	// Output:
	// +-------------------------------+
	// |5 9 8 2 6 B A F 3 D 4 7 C 1 G E|
	// |A 1 6 4 E C 2 3 B 5 8 G D 9 F 7|
	// |D B F E 7 4 1 G 2 6 C 9 8 A 3 5|
	// |C 7 G 3 5 D 9 8 F A E 1 6 4 B 2|
	// |4 D 5 9 2 8 F 7 6 B 1 A 3 C E G|
	// |F G A C B E 4 5 9 7 D 3 1 2 6 8|
	// |1 2 7 8 3 6 D A C 4 G E B 5 9 F|
	// |3 6 E B 9 G C 1 5 F 2 8 A 7 4 D|
	// |E 8 9 F C 5 3 4 7 G A D 2 6 1 B|
	// |G 3 D 7 F 2 8 B 1 C 6 4 9 E 5 A|
	// |6 5 4 1 D A E 9 8 2 F B G 3 7 C|
	// |B C 2 A G 1 7 6 E 3 9 5 F D 8 4|
	// |7 4 B D 1 9 G C A 8 5 6 E F 2 3|
	// |8 E C 5 4 F 6 D G 1 3 2 7 B A 9|
	// |2 F 1 G A 3 B E 4 9 7 C 5 8 D 6|
	// |9 A 3 6 8 7 5 2 D E B F 4 G C 1|
	// +-------------------------------+
}

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
	// |8 1 2 7 5 3 6 4 9|
	// |9 4 3 6 8 2 1 7 5|
	// |6 7 5 4 9 1 2 8 3|
	// |1 5 4 2 3 7 8 9 6|
	// |3 6 9 8 4 5 7 2 1|
	// |2 8 7 1 6 9 5 3 4|
	// |5 2 1 9 7 4 3 6 8|
	// |4 3 8 5 2 6 9 1 7|
	// |7 9 6 3 1 8 4 5 2|
	// +-----------------+
}

func ExampleSudoku_Resolve_hardest3() {
	sudoku := NewSudoku()
	sudoku.BruteLimit = 2
	sudoku.ParseString(`
+-----------------+
|  8         1 5  |
|4   6 5   9   8  |
|          8      |
|                 |
|    2   4       3|
|3     8   1      |
|9       7        |
|6               4|
|1 5           9  |
+-----------------+
`)
	sudoku.Resolve()
	fmt.Println(sudoku.String())
	// Output:
	// +-----------------+
	// |2 8 3 4 6 7 1 5 9|
	// |4 7 6 5 1 9 3 8 2|
	// |5 9 1 2 3 8 6 4 7|
	// |7 4 5 6 9 3 8 2 1|
	// |8 1 2 7 4 5 9 6 3|
	// |3 6 9 8 2 1 4 7 5|
	// |9 2 4 1 7 6 5 3 8|
	// |6 3 8 9 5 2 7 1 4|
	// |1 5 7 3 8 4 2 9 6|
	// +-----------------+
}

func ExampleSudoku_Resolve_hardest4() {
	sudoku := NewSudoku()
	sudoku.BruteLimit = 2
	sudoku.ParseString(`
+-----------------+
|  8         1 5  |
|4   6 5   9   8  |
|          8      |
|                 |
|    2   4       3|
|3     8   1      |
|9       7        |
|6               4|
|1 5           9  |
+-----------------+
`)
	sudoku.Resolve()
	fmt.Println(sudoku.String())
	// Output:
	// +-----------------+
	// |2 8 3 4 6 7 1 5 9|
	// |4 7 6 5 1 9 3 8 2|
	// |5 9 1 2 3 8 6 4 7|
	// |7 4 5 6 9 3 8 2 1|
	// |8 1 2 7 4 5 9 6 3|
	// |3 6 9 8 2 1 4 7 5|
	// |9 2 4 1 7 6 5 3 8|
	// |6 3 8 9 5 2 7 1 4|
	// |1 5 7 3 8 4 2 9 6|
	// +-----------------+
}

func ExampleSudoku_Resolve_heart() {
	sudoku := NewSudoku()
	sudoku.BruteLimit = 2
	sudoku.ParseString(`
+-----------------+
|  5 8       4 1  |
|7     4   5     3|
|2       1       9|
|9       4       2|
|  7           3  |
|  6           5  |
|    1       8    |
|      2   7      |
|        5        |
+-----------------+
`)
	sudoku.Resolve()
	fmt.Println(sudoku.String())
	// Output:
	// +-----------------+
	// |3 5 8 6 2 9 4 1 7|
	// |7 1 9 4 8 5 6 2 3|
	// |2 4 6 7 1 3 5 8 9|
	// |9 8 3 5 4 1 7 6 2|
	// |5 7 2 8 9 6 1 3 4|
	// |1 6 4 3 7 2 9 5 8|
	// |6 2 1 9 3 4 8 7 5|
	// |8 9 5 2 6 7 3 4 1|
	// |4 3 7 1 5 8 2 9 6|
	// +-----------------+
}
