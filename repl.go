package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kryptonn36/pokedexcli/internal/pokeapi"
)
type config struct{
	pokeapiClient pokeapi.Client 
	NextLocation *string
	PreviousLocation *string
	caughtPokemon    map[string]pokeapi.Pokemon
	// area_name string
}

func startRepl(cfg *config) {
	output := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		if !output.Scan(){
			break
		}

		text := output.Text()
		cleaned_text := cleanInput(text)
		// commandName := cleaned_text[0]
		if len(cleaned_text)==0{
			continue
		}

		args := []string{}
		if len(cleaned_text)>1{
			args = cleaned_text[1:]
		}

		if c,ok := commands[cleaned_text[0]]; ok{
			if err := c.callback(cfg, args...); err!=nil{
				fmt.Printf("Error %v",err)
			}
		}else{
			fmt.Printf("Command not found\n")
		}
	}

	if err := output.Err(); err!=nil{
		fmt.Fprintf(os.Stderr,"err reading input: %v\n",err)
	}
}


func cleanInput(text string) []string{
	var sliceText []string
	textLower := strings.ToLower(text)
	// textLower = strings.TrimSpace(textLower)
	sliceText = strings.Fields(textLower)
	return sliceText
}


type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

var commands = map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
	"help": {
		name:        "help",
		description: "Show the help log",
		callback:     help,
	},
	"map":{
		name:         "map",
		description:  "Show the locations on Map",
		callback:      commandMapf,
	},
	"mapb":{
		name:			"mapb",
		description:    "Show the Previous locations on Map",
		callback:       commandMapb,
	},
	"explore":{
		name           : "explore <location_name>",
		description    : "List of Pokemon at a Location Area",
		callback       : explore_area,
	},
	"catch":{
		name           : "catch <pokemon_name>",
		description    : "Catch the pokemon",
		callback       : catch_pokemon,
	},
	"inspect":{
		name           : "inspect <pokemon_name>",
		description    : "To inspect a Pokemon",
		callback       :  inspect_pokemon,
	},
	"pokedex":{
		name           : "pokedex",
		description    : "list of all pokedex caught",
		callback       : caught_pokemon,
	},
}