package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListAreaInfo(areaName string) (RespAreaInfo, error) {
	url := baseURL + "/location-area/" + areaName
	if areaName == "" {
		return RespAreaInfo{}, fmt.Errorf("please provide an area name in order to get area info")
	}

	if val, ok := c.cache.Get(url); ok {
		areaInfoResp := RespAreaInfo{}
		err := json.Unmarshal(val, &areaInfoResp)
		if err != nil {
			return RespAreaInfo{}, err
		}

		return areaInfoResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespAreaInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespAreaInfo{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespAreaInfo{}, err
	}

	areaInfoResp := RespAreaInfo{}
	err = json.Unmarshal(body, &areaInfoResp)
	if err != nil {
		return RespAreaInfo{}, err
	}

	c.cache.Add(url, body)
	return areaInfoResp, nil
}
