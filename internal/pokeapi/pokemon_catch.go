package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Pokemon struct{
	Id                int             `json:"id"`
	Name              string          `json:"name"`
	Base_experience   int             `json:"base_experience"`
	Height            int             `json:"height"`
	Is_default        bool            `json:"is_default"`
	Weight            int             `json:"weight"`
	Abilities [] struct{
		Is_hidden    bool             `json:"is_hidden"`
		Slot         int              `json:"slot"`
		Ability struct{
			Name     string           `json:"name"`
			URL      string           `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct{
		Name         string           `json:"name"`
		URL          string           `json:"url"`
	} `json:"forms"`
	Game_indices [] struct{
		Game_index   int              `json:"game_index"`
		Version struct{
			Name      string          `json:"name"`
			URL       string          `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Held_items []struct{
		Item struct{
			Name      string          `json:"name"`
			URL       string          `json:"url"`
		} `json:"item"`
		Version_details []struct{
			Version struct{
				Name  string          `json:"name"`
				URL   string          `json:"url"`
			} `json:"version"`
			Rarity    int             `json:"rarity"`
		} `json:"version_details"`
	} `json:"held_items"`
	Location_area_encounters  string  `json:"location_area_encounters"`
	Moves []struct{
		Move struct{
			Name     string           `json:"name"`
			URL      string           `json:"url"`
		} `json:"move"`
		Version_group_details []struct{
			Move_learn_method struct{
				Name    string        `json:"name"`
				URL     string        `json:"url"`
			} `json:"move_learn_method"`
			Version_group struct{
				Name    string         `json:"name"`
				URL     string         `json:"url"`
			} `json:"version_group"`
			Level_learnde_at  int       `json:"level_learned_at"`
			Order             int       `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Past_types []struct{
		Generation struct{
			Name             string        `json:"name"`
			URL              string        `json:"url"`
		} `json:"generation"`
		Types []struct{
			Slot             int            `json:"slot"`
			Type struct{
				Name         string          `json:"name"`
				URL          string          `json:"url"`
			} `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
	Past_abilities []struct{
		Generation struct{
			Name            string          `json:"name"`
			URL             string          `json:"url"`
		} `json:"generation"`
		Abilities [] struct{
			Is_hidden    bool             `json:"is_hidden"`
			Slot         int              `json:"slot"`
			Ability struct{
				Name     string           `json:"name"`
				URL      string           `json:"url"`
			} `json:"ability"`
		} `json:"abilities"`
	} `json:"past_abilities"`
	Sprites struct{
		Front_default           string          `json:"front_default"`
		Front_shiny             string          `json:"front_shiny"`
		Front_female            string          `json:"front_female"`
		Front_shiny_female      string          `json:"front_shiny_female"`
		Back_default            string          `json:"back_default"`
		Back_shiny              string          `json:"back_shiny"`
		Back_female             string          `json:"back_female"`
		Back_shiny_female       string          `json:"back_shiny_female"`
	} `json:"sprites"`
	Cries struct{
		Latest                 string           `json:"latest"`
		Legacy                 string           `json:"legacy"`
	} `json:"cries"`
	Species struct{
		Name                  string            `json:"name"`
		URL                   string            `json:"url"`
	} `json:"species"`
	Stats [] struct{
		Stat struct{
			Name              string            `json:"name"`
			URL               string            `json:"url"`
		} `json:"stat"`
		Effort                int               `json:"effort"`
		Base_stat             int               `json:"base_stat"`
	} `json:"stats"`
	Types []struct{
		Slot                  int               `json:"slot"`
		Type struct{
			Name              string            `json:"name"`
			URL               string            `json:"url"`
		} `json:"type"`
	} `json:"types"`
}



func (c *Client) PokemonType(pokemon_name string) (Pokemon,error){
	url := baseURL1 + "/pokemon/" + pokemon_name

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	req,err := http.NewRequest("GET", url, nil)
	if err!= nil{
		return Pokemon{}, err
	}

	resp,err := c.httpclient.Do(req)
	if err != nil{
		return Pokemon{}, err
	}

	pokemon_details := Pokemon{}

	result,err := io.ReadAll(resp.Body)
	if err!= nil{
		return pokemon_details,err
	}
	defer resp.Body.Close()
	
	err = json.Unmarshal(result,&pokemon_details)
	if err!=nil{
		return pokemon_details,err
	}
	return pokemon_details,nil
}