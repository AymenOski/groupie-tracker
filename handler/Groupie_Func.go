package handler

import (
	"encoding/json"
	"net/http"

	"groupie/helpers"
	tools "groupie/tools"
)

func Groupie_Func(w http.ResponseWriter, r *http.Request) {
	// Groupie_Func handles HTTP requests for the root path "/"
	// First, it checks if the path is valid, then checks the method,
	// and finally, it retrieves data from an external API to display it on a webpage.

	// Check if the requested path is not "/" (root path)
	if r.URL.Path != "/" {
		// If the path is invalid, render a 404 Not Found page
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorNotFound, http.StatusNotFound)
		return
	}

	// Check if the HTTP method is GET
	if r.Method != http.MethodGet {
		// If the method is not GET, render a 405 Method Not Allowed page
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorMethodnotAll, http.StatusMethodNotAllowed)
		return
	}

	// Declare a slice to hold artists data
	var Artists []tools.Artists

	// Define the URL to fetch data from the external API
	url := "https://groupietrackers.herokuapp.com/api/artists"

	// Send a GET request to the API
	res, err := http.Get(url)
	if err != nil {
		// If there's an error with the GET request (e.g., network failure), render an internal server error page
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}

	// Close the response body once we're done with it (to avoid resource leak)
	defer res.Body.Close()

	// Decode the JSON response body into the Artists slice
	err = json.NewDecoder(res.Body).Decode(&Artists)
	if err != nil {
		// If there's an error during JSON decoding (e.g., malformed JSON), render an internal server error page
		helpers.RenderTemplates(w, "statusPage.html", tools.ErrorInternalServerErr, http.StatusInternalServerError)
		return
	}

	// Render the index.html template, passing in the artists data and a 200 OK status
	helpers.RenderTemplates(w, "index.html", Artists, http.StatusOK)
}
