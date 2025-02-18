package main

import (
	"fmt"
	"net/http"

	"Student/functions"
)

func main() {
	http.HandleFunc("/", functions.MainPage)
	fmt.Println("The server is runing at : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
