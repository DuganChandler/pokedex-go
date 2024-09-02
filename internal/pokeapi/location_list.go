package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespAreaLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespAreaLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespAreaLocations{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespAreaLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespAreaLocations{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespAreaLocations{}, nil
	}

	locationsResponse := RespAreaLocations{}
	err = json.Unmarshal(body, &locationsResponse)
	if err != nil {
		return RespAreaLocations{}, err
	}

	c.cache.Add(url, body)
	return locationsResponse, nil
}
