package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// import "go/version"

type pokelist struct{
	Id  int              `json:"id"` 
	Name string          `json:"name"`
	Game_index int       `json:"game_index"`
	Encounter_method_rates []struct{
		Encounter_method struct{
			Name string      `json:"name"`
			URL string       `json:"url"`
		} `json:"encounter_method"`
		Version_details []struct{
			Rate int          `json:"rate"`
			Version struct{
				Name string    `json:"name"`
				URL string     `json:"url"`
			}`json:"version_details"`
		}`json:"encounter_method_rates"`
	}
	Location struct{
		Name string           `json:"name"`
		URL  string           `json:"url"`
	}`json:"location"`
	Names []struct{
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}`json:"language"`
		Name string     `json:"name"`
	}`json:"names"`
	Pokemon_encounters []struct{
		Pokemon struct{
			Name string        `json:"name"`
			URL  string        `json:"url"`
		}`json:"pokemon"`
		Version_details []struct{
			Version struct{
				Name string    `json:"name"`
				URL  string    `json:"url"`
			}`json:"version"`
			Max_chance int     `json:"max_chance"`
			Encounter_detals []struct{
				Min_level   int     `json:"min_level"`
				Max_level   int     `json:"max_level"`
				Conditon_values  struct{
					Name string      `json:"name"`
					URL  string      `json:"url"`
				}`json:"condition"`
				Chance int           `json:"chance"`
				Method struct{
					Name string      `json:"name"`
					URL  string      `json:"url"`
				}`json:"method"`
			}`json:"encounter_detals"`
		}`json:"version_details"`
	}`json:"pokemon_encounters"`

}

func (c * Client) Pokemon_list(area_name string) (pokelist,error){
	url := baseURL2 + "/"+ area_name
	
	if val, ok := c.cache.Get(url); ok {
		locationResp := pokelist{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return pokelist{}, err
		}
		return locationResp, nil
	}

	req,err := http.NewRequest("GET", url, nil)
	if err!= nil{
		return pokelist{},err
	}
	resp,err := c.httpclient.Do(req)
	if err!= nil{
		return pokelist{}, err
	}
	final_list := pokelist{}

	body, err := io.ReadAll(resp.Body)
	if err!= nil{
		return final_list,err
	}
	defer resp.Body.Close()

	err = json.Unmarshal(body, &final_list)
	if err!=nil{
		return final_list,err
	}
	return final_list,nil
}