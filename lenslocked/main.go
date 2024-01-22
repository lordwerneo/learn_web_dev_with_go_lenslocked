package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)

	fmt.Println("Starting server at :8000")
	http.ListenAndServe(":8000", nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Just H1 tag</h1>")
}
