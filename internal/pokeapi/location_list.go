package pokeapi

import (
	"encoding/json"
	"io"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (LocationAPIResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	resp, err := c.httpClient.Get(url)
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