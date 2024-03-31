package pokeapi

import (
	"encoding/json"
	"io"
)

// ExploreLocations -
func (c *Client) ExploreLocation(location string) (AreaAPIResponse, error) {
	url := baseURL + "/location-area/" + location
	if location == "" {
		return AreaAPIResponse{}, nil
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return AreaAPIResponse{}, err
	}

	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaAPIResponse{}, err
	}

	areaResp := AreaAPIResponse{}
	err = json.Unmarshal(dat, &areaResp)
	if err != nil {
		return AreaAPIResponse{}, err
	}

	return areaResp, nil
}