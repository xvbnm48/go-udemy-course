package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	// n, err := fmt.Fprintf(w, "Hello, World!")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(fmt.Sprintf("number of bytes written:  %d", n))

	fmt.Fprintf(w, "Hello, World!")
}

// about page
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About")
}

func main() {
	// fmt.Println("Hello, World!")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Server is running on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
