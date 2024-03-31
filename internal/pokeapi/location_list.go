package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (LocationAPIResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAPIResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAPIResponse{}, err
	}
	
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAPIResponse{}, err
	}

	locationsResp := LocationAPIResponse{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationAPIResponse{}, err
	}

	return locationsResp, nil
}