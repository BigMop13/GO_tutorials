package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", printMultiplicationTable)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func printMultiplicationTable(w http.ResponseWriter, r *http.Request) {
	for x := 1; x <= 10; x++ {
		for y := 1; y <= 10; y++ {
			result := x * y
			io.WriteString(w, fmt.Sprintf("%4d", result))
		}
		io.WriteString(w, "\n")
	}
}
