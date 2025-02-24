package helpers

import (
	"encoding/json"
	"net/http"
)

func Fetch_By_Id(url string, target interface{}) error {
	// get the data from the url
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// read the response body and decode it to the target variable
	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil
}
