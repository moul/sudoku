package sapinapp

import "net/http"

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//sudo := sudoku.NewSudoku()
}
