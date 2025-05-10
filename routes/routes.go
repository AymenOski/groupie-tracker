package routes

import (
	"log"
	"net/http"
	"text/template"

	"groupie/handler"
	tools "groupie/tools"
)

func Route() {
	initTemplates()
	initRoutes()
}

func initTemplates() {
	var err error
	tools.Tp, err = template.ParseGlob("template/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
}

func initRoutes() {
	http.HandleFunc("/static/", handler.Style_Func)
	http.HandleFunc("/details/", handler.Detail_Card_Func)
	http.HandleFunc("/", handler.Groupie_Func)
}
