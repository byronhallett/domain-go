package domain

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// curl -X POST -u '{client_id}:{client_secret}' -H "Content-Type: application/x-www-form-urlencoded" -d 'grant_type=client_credentials&scope=api_listings_read%20api_agencies_read'

const authURL = "https://auth.domain.com.au/v1/connect/token"

type accessResponse struct {
	AccessToken string `json:"access_token"`
}

// GetAccessToken gets the token which can be used in future api calls
func getAccessToken(clientID string, secret string) (string, error) {
	values := url.Values{
		"grant_type": []string{"client_credentials"},
		"scope":      []string{"api_listings_read", "api_agencies_read"},
	}
	req, err := http.NewRequest("POST", authURL, strings.NewReader(values.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, secret)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	responseObject := new(accessResponse)
	json.Unmarshal(b, responseObject)
	return responseObject.AccessToken, nil
}
