package sapinapp

import (
	"net/http"

	"github.com/moul/sudoku"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	sudo := sudoku.NewSudoku()
}
