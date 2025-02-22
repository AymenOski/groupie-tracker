package handler

import (
	"net/http"
	"os"
	"strings"

	"groupie/helpers"
	tools "groupie/tools"
)

func Style_Func(w http.ResponseWriter, r *http.Request) {
	filePath := strings.TrimPrefix(r.URL.Path, "/")
	File, err := os.Stat(filePath)
	if err != nil || File.IsDir() {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorNotFound, http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}
