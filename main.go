package main

import (
	"fmt"
	"net/http"

	"github.com/TuTuRyYE/block-chain/pkg"
	"github.com/gorilla/mux"
)

func main() {
	blockChain := pkg.NewBlockChain()
	blockChain.AddBlock("First Bloc")
	blockChain.AddBlock("Second Bloc")

	r := mux.NewRouter()
	r.HandleFunc("/blockchain", func(w http.ResponseWriter, r *http.Request) {
		blockChain.Print(w)
		fmt.Fprintf(w, "End of the BlockChain")
	})
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
