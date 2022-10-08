package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprintf(w, "Hello \n ")
	fmt.Fprintf(w, "Hello. I’m v2 \n")
}
func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
