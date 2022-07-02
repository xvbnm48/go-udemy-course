package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	// fmt.Println("Hello, World!")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Server is running on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
