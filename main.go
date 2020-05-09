package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := io.WriteString(w, "Hello, world!\n")
		if err != nil {
			http.Error(w, "Something bad happened", 500)
		}
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
