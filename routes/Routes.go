package routes

import (
	"fmt"
	"net/http"
	"text/template"

	"Student/functions"
	"Student/tools"
)

func Route() {
	var err error
	tools.Temp, err = template.ParseGlob("../templates/*.html")
	if err != nil {
		fmt.Println("Error parsing templates", err)
		return
	}
	http.HandleFunc("/static/", functions.StaticHandler)
	http.HandleFunc("/", functions.MainPage)
	http.HandleFunc("/more-info/", functions.MoreInfo)
}
