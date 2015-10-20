package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hiya, pal")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8123", nil)
}
