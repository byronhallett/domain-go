package domain

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"

	"github.com/byronhallett/domain-go/types"
)

const residentialURL = "https://api.domain.com.au/v1/listings/residential/_search"

// ResidentialSearch searches for residential listings...
func (s *Session) ResidentialSearch(params types.SearchParameters) (*[]types.SearchResult, error) {
	// Clamp the page size
	params.PageSize = 200
	results := new([]types.SearchResult)
	c := make(chan *[]types.SearchResult)
	e := make(chan error)
	for page := 1; page < 6; page++ {
		go s.searchPage(residentialURL, params, page, c, e)
	}
	for i := 0; i < 5; i++ {
		select {
		case nextResults := <-c:
			*results = append(*results, *nextResults...)
		case err := <-e:
			return nil, err
		}
	}
	return results, nil
}

func (s *Session) searchPage(url string, params types.SearchParameters, page int, c chan *[]types.SearchResult, e chan error) {
	params.Page = page
	results := new([]types.SearchResult)
	body, err := json.Marshal(params)
	if err != nil {
		e <- err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		e <- err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.accessToken)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		e <- err
	}
	defer response.Body.Close()
	jsonData, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(jsonData, results)
	if err != nil {
		spew.Dump(jsonData)
		e <- err
	}
	c <- results
}
