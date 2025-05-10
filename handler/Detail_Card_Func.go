package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"groupie/helpers"
	"groupie/tools"
)

func Detail_Card_Func(w http.ResponseWriter, r *http.Request) {
	// check the method
	if r.Method != http.MethodGet {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}

	// parsing id ...
	id := r.URL.Query().Get("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorBadReq, http.StatusBadRequest)
		return
	}

	// fetch main artist data synchronously
	var artistFound *tools.Artists
	if err := helpers.Fetch_By_Id(
		fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", Id),
		&artistFound,
	); err != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	if artistFound.Id == 0 {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorNotFound, http.StatusNotFound)
		return
	}

	// prepare containers for nested data
	var locations *tools.Locations
	var dates *tools.ConcertDates
	var relations *tools.Relations
	locationsURL := artistFound.Locations
	datesURL := artistFound.ConcertDates
	relationsURL := artistFound.Relations

	// use WaitGroup to fetch nested APIs concurrently
	var wg sync.WaitGroup
	var fetchErr error
	var mu sync.Mutex
	// set counter = 3 tasks
	wg.Add(3)

	go func() {
		defer wg.Done()
		if err := helpers.Fetch_By_Id(locationsURL, &locations); err != nil {
			mu.Lock()
			fetchErr = err
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		if err := helpers.Fetch_By_Id(datesURL, &dates); err != nil {
			mu.Lock()
			fetchErr = err
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		if err := helpers.Fetch_By_Id(relationsURL, &relations); err != nil {
			mu.Lock()
			fetchErr = err
			mu.Unlock()
		}
	}()

	// wait for all fetches to complete
	wg.Wait()
	if fetchErr != nil {
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}
	type fetchingData struct {
		Artist    *tools.Artists
		Locations *tools.Locations
		Dates     *tools.ConcertDates
		Relations *tools.Relations
	}
	// aggregate all into one struct
	fetching_data := &fetchingData{
		Artist:    artistFound,
		Locations: locations,
		Dates:     dates,
		Relations: relations,
	}

	helpers.RenderTemplates(w, "detailsCard.html", fetching_data, http.StatusOK)
}
