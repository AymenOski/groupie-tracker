package functions

import (
	"encoding/json"
	"net/http"

	"Student/tools"
)

var Artists []tools.Artists

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderErrPage(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		RenderErrPage(w, http.StatusMethodNotAllowed)
		return
	}
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"

	resp, err := http.Get(apiURL)
	if err != nil {
		RenderErrPage(w, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&Artists)
	if err != nil {
		RenderErrPage(w, http.StatusInternalServerError)
		return
	}
	err = tools.Temp.ExecuteTemplate(w, "index.html", Artists)
	if err != nil {
		RenderErrPage(w, http.StatusInternalServerError)
		return
	}
}
