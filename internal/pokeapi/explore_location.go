package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ExploreLocations -
func (c *Client) ExploreLocation(location string) (AreaAPIResponse, error) {
	url := baseURL + "/location-area/" + location
	if location == "" {
		return AreaAPIResponse{}, nil
	}
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaAPIResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
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