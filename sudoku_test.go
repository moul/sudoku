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

func ExampleSudoku_Resolve_5x5() {
	sudoku := NewSudokuWithSize(5)
	sudoku.BruteLimit = 0
	sudoku.ParseString(`
+-------------------------------------------------+
|L U A     I   V     Y P   C   H N E F   T     G X|
|      W       Y T   R   Q N K     D X   O   C    |
|        D X       A       M       W   V       B U|
|E K N H     C P S   L T     U J     B Q   A F D  |
|T   X S P   U   N O H G D   J   L   C     I   W V|
|A Y       J           U   V X D   L     B W N Q O|
|  F   L     P   A   B       M U E     S     V    |
|D   J   W   O S C U   L         T     I G     M  |
|    B I M W L   V H S Q N   O         G   T R   K|
|  T   Q H   N G         E I           B   X S U  |
|R     V   O X K     I     A L   Q B         M T  |
|  M D J   L             H   B     A O E R S     C|
|    O         C H   U       P   V M         Y    |
|K     T L M W B     N   S             U   H O I  |
|  H C         I E   D O     F     K G N   V     L|
|  C M A   K           S R         I D   Y E   V  |
|Q   U F   V         W   L O G A S   M J D K I    |
|  V     I S     Q         D   C O H N   F   T   P|
|    P     N     W T M       A   K   E     O   L  |
|G O K X B     J   R F I   P           Y       N H|
|M P   B     K   X   G   A S V E J   T   I C L   Q|
|  G S O   C H     I P     K Q   F V R     J D X T|
|F Q       D   A       R       M       X E        |
|    W   X   V L     J B F   T   H C       U      |
|J L     C   Q N O M   D   H Y     G   W     B S R|
+-------------------------------------------------|

`)
	sudoku.Resolve()
	fmt.Println(sudoku.String())
	// Output:
	// +-------------------------------------------------+
	// |L U A M Q I D V K B Y P W C S H N E F O T R J G X|
	// |V B I W F H E Y T J R A Q N K G U D X M O L C P S|
	// |C J Y G D X R Q L A E F O M I T P W S V K N H B U|
	// |E K N H O G C P S W L T V X U J I R B Q M A F D Y|
	// |T R X S P F U M N O H G D B J K L Y C A Q I E W V|
	// |A Y G C S J F R I E T U P V X D M L K H B W N Q O|
	// |O F R L K Q P T A X B Y G W M U E N J S H D V C I|
	// |D E J N W B O S C U A L K R H V T X Q I G Y P M F|
	// |U X B I M W L D V H S Q N F O Y C P A G J T R E K|
	// |P T V Q H Y N G M K C J E I D O R F W B L X S U A|
	// |R S F V U O X K Y N I E J A L W Q B H C P G M T D|
	// |X M D J N L T F U P V W H G B I Y A O E R S Q K C|
	// |I W O E G A S C H D U K X Q P R V M L T N B Y F J|
	// |K A Q T L M W B G V N C S Y R F D J P U X H O I E|
	// |B H C P Y R J I E Q D O M T F S X K G N U V W A L|
	// |H C M A T K B O F L Q S R J N X G I D P Y E U V W|
	// |Q N U F E V Y X P C W H L O G A S T M J D K I R B|
	// |W V L Y I S A U Q G K X B D E C O H N R F M T J P|
	// |S D P R J N I H W T M V Y U A B K Q E F C O X L G|
	// |G O K X B E M J D R F I T P C L W U V Y S Q A N H|
	// |M P H B R U K W X F G N A S V E J O T D I C L Y Q|
	// |Y G S O A C H E B I P M U K Q N F V R L W J D X T|
	// |F Q T U V D G A J Y O R C L W M B S I X E P K H N|
	// |N I W D X P V L R S J B F E T Q H C Y K A U G O M|
	// |J L E K C T Q N O M X D I H Y P A G U W V F B S R|
	// +-------------------------------------------------+
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
