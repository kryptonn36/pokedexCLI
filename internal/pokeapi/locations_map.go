package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationsResp struct{
	Count int        `"json:count"`
	Next *string     `"json:next"`
	Previous *string `"json:previous"`
	Results []struct{
		Name string   `"json:name"`
		URL string    `"json:url"`
	} `"json:results"`
}

func (c *Client, ) ListLocation(pageURL *string) (LocationsResp, error){
	url := baseURL1 + "/location-area"
	if pageURL != nil{
		url = *pageURL
	}
	var body []byte
	var available bool
	body, available = c.cache.Get(url)

	req,err := http.NewRequest("GET", url, nil)
	if err!=nil{
		return LocationsResp{}, err
	}

	if !available {
		resp,err := c.httpclient.Do(req)
		if err!=nil{
			return LocationsResp{}, err
		}

		body,err = io.ReadAll(resp.Body)
		if err!= nil{
			return LocationsResp{}, err
		}
		defer resp.Body.Close()

		c.cache.Add(url, body)
	}
	locationslist := LocationsResp{}

	err = json.Unmarshal(body, &locationslist)
	if err != nil {
		return LocationsResp{}, err
	}
	return locationslist, nil
}