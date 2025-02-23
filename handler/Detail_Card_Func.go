package handler

import (
	"net/http"
	"strconv"

	"groupie/helpers"
	"groupie/tools"
)

func Detail_Card_Func(w http.ResponseWriter, r *http.Request) {
	// check the method
	if r.Method != http.MethodGet {
		// execute the status page template
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}
	type fetchingData struct {
		Artist    *tools.Artists
		Locations *tools.Locations
		Dates     *tools.ConcertDates
		Relations *tools.Relations
	}
	// Get the id from the url
	id := r.URL.Query().Get("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorBadReq, http.StatusBadRequest)
		return
	}
	var artistFound *tools.Artists
	if Id > 0 && Id < 53 {
		er := helpers.Fetch_By_Id("https://groupietrackers.herokuapp.com/api/artists/{Id}", &artistFound)
		if er != nil {
			helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
			return
		}
		artistFound = &Artists[Id-1]
	}
	if artistFound == nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorNotFound, http.StatusNotFound)
		return
	}
	var locations *tools.Locations
	var dates *tools.ConcertDates
	var relations *tools.Relations
	// fetching the nested data
	errLoc := helpers.Fetch_By_Id(artistFound.Locations, &locations)
	errDate := helpers.Fetch_By_Id(artistFound.ConcertDates, &dates)
	errRel := helpers.Fetch_By_Id(artistFound.Relations, &relations)

	if errLoc != nil || errDate != nil || errRel != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	// set all the data  that we found into the fetching variable
	fetching_data := fetchingData{
		Artist:    artistFound,
		Locations: locations,
		Dates:     dates,
		Relations: relations,
	}
	helpers.RenderTemplates(w, "detailsCard.html", fetching_data, http.StatusOK)
}
