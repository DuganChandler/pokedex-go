package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonInfo(pokemonName string) (RespPokemonInfo, error) {
	url := baseURL + "/pokemon/" + pokemonName
	if pokemonName == "" {
		return RespPokemonInfo{}, fmt.Errorf("please provide a pokemon in order to get thier info")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonInfo{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	pokemonInfo := RespPokemonInfo{}
	err = json.Unmarshal(body, &pokemonInfo)
	if err != nil {
		return RespPokemonInfo{}, fmt.Errorf("unable to unmarshal pokemon info %v", err)
	}

	return pokemonInfo, nil
}
