package main

import (
	"fmt"
	"net/http"
	"cowsay/cowsay"
	"cowsay/fortune"
)

var fortunes []string

func fortuneHandler(w http.ResponseWriter, r *http.Request) {
	f := cowsay.WrapText(fortune.RandomFortune(fortunes), 50)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, cowsay.Say(f))
}

func main() {
	var err error
	fortunes, err = fortune.LoadFortunes("fortunes.txt")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/fortune", fortuneHandler)
	fmt.Println("Server running at http://localhost:8080/fortune")
	http.ListenAndServe(":8080", nil)
}
