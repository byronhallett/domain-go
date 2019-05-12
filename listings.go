package domain

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const residentialURL = "https://api.domain.com.au/v1/listings/residential/_search"

// Search for residential listings
func ResidentialSearch(params SearchParameters) (*SearchResult, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	response, err := http.Post(
		residentialURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	results := new(SearchResult)
	jsonData, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(jsonData, results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
