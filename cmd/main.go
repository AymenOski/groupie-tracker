package main

import (
	"fmt"
	"net/http"

	"groupie/routes"
)

func main() {
	// Register handlers
	routes.Route()
	//  run   the server
	fmt.Println("Server running at http://localhost:8050/")
	http.ListenAndServe(":8050", nil)
}
