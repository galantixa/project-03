package main

import (
	"fmt"
	"net/http"

	authcontroller "github.com/galantixa/go-auth2/controllers"
)
func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/Login", authcontroller.Login)
	fmt.Println("Server running on http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}