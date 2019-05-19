package domain

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/byronhallett/domain-go/types"
)

const residentialURL = "https://api.domain.com.au/v1/listings/residential/_search"

// ResidentialSearch searches for residential listings...
func (s *Session) ResidentialSearch(params types.SearchParameters) (*[]types.SearchResult, error) {

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", residentialURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.accessToken)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	results := new([]types.SearchResult)
	jsonData, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(jsonData, results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
