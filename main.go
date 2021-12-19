package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

//main function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("\nStarting application on port, %s", port)

	_ = http.ListenAndServe(port, nil)
}
