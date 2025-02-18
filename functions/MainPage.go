package functions

import (
	"fmt"
	"net/http"
	"os"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"

	artists, err := fetching(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Artists :\n", artists)
}
