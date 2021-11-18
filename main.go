package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/elevenfofiv/crytocurrency/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmp1 := template.Must(template.ParseFiles("templates/home.gohtml"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmp1.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
