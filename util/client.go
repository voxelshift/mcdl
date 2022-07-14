package util

import (
	"encoding/json"
	"net/http"
)

// Make a GET request to the given URL and parse its response into the provided interface
func GetJson(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return err
	}

	return nil
}
