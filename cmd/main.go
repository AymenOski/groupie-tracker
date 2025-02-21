package main

import (
	"fmt"
	"net/http"

	"Student/routes"
)

func init() {
	routes.Route()
}

func main() {
	fmt.Println("Server is running at http://localhost:8080 ")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
