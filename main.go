package main

import (
	"errors"
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

func devide(w http.ResponseWriter, r *http.Request) {
	f, err := devideValue(100.2, 1.0)
	if err != nil {
		fmt.Fprintf(w, "cannot devided by zero")
	}

	fmt.Fprintf(w, fmt.Sprintf("%f devided by %f", 100.0, 10.0, f))

}

func devideValue(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot devide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	// fmt.Println("Hello, World!")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/devide", devide)

	fmt.Println("Server is running on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
