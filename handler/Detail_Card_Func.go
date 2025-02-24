package handler

import (
	"fmt"
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
	// to int
	Id, err := strconv.Atoi(id)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorBadReq, http.StatusBadRequest)
		return
	}
	var artistFound *tools.Artists
	errFetch := helpers.Fetch_By_Id(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", Id), &artistFound)
	if artistFound.Id == 0 {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorNotFound, http.StatusNotFound)
		return
	}
	var locations *tools.Locations
	var dates *tools.ConcertDates
	var relations *tools.Relations
	// declarations of urls
	Locations_url := artistFound.Locations
	ConcertDates_url := artistFound.ConcertDates
	Relations_url := artistFound.Relations

	// fetch the nested data
	errLoc := helpers.Fetch_By_Id(Locations_url, &locations)
	errLates := helpers.Fetch_By_Id(ConcertDates_url, &dates)
	errRelation := helpers.Fetch_By_Id(Relations_url, &relations)

	if errLoc != nil || errLates != nil || errRelation != nil || errFetch != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	// set all the data that we fetched into one variable
	fetching_data := fetchingData{
		Artist:    artistFound,
		Locations: locations,
		Dates:     dates,
		Relations: relations,
	}
	helpers.RenderTemplates(w, "detailsCard.html", fetching_data, http.StatusOK)
}
