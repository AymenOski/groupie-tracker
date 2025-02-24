package handler

import (
	"encoding/json"
	"net/http"

	"groupie/helpers"
	tools "groupie/tools"
)


func Groupie_Func(w http.ResponseWriter, r *http.Request) {
	var Artists []tools.Artists
	// check the path
	if r.URL.Path != "/" {
		// execute the not found  template
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorNotFound, http.StatusNotFound)
		return
	}
	// check the methd
	if r.Method != http.MethodGet {
		// execute the not found  template
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}
	url := "https://groupietrackers.herokuapp.com/api/artists"
	// get the api data
	res, err := http.Get(url)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	// decode the jsone data
	err = json.NewDecoder(res.Body).Decode(&Artists)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	helpers.RenderTemplates(w, "index.html", Artists, http.StatusOK)
}
