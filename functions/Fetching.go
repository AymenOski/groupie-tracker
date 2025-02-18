package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"location"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func fetching(apiURL string) ([]Artist, error) {
	//get request
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()
	fmt.Println(resp)
	// turning the data into a slice of bytes
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	var artists []Artist
	//json to Go
	err = json.Unmarshal(body, &artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}
