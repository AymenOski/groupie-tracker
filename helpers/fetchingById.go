package helpers

import (
	"encoding/json"
	"net/http"
)

func Fetch_By_Id(url string, target interface{}) error {
	// get the data from   url
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// decode the data and set it into the target var
	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil
}
