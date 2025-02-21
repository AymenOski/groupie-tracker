package functions

import (
	"net/http"

	"Student/tools"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderPageNotFound(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		RenderPageNotFound(w, http.StatusMethodNotAllowed)
		return
	}
	FeetchingDate()
	err := tools.Temp.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		RenderPageNotFound(w, http.StatusInternalServerError)
		return
	}
}
