package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	err := http.ListenAndServe("0.0.0.0:14774", nil)

	if err != nil {
		log.Fatal(err)
	}
}
