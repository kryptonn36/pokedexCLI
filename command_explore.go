package main

import (
	// "bufio"
	// "os"
	"fmt"
	"errors"

	// "github.com/kryptonn36/pokedexcli/internal/pokeapi"
)

func explore_area(cfg *config, args ...string) error{
	if len(args) != 1{
		return errors.New("you need to provide a location name")
	}
	area_name := args[0]
	location_list,err := cfg.pokeapiClient.Pokemon_list(area_name)
	if err!= nil{
		return err
	}
	fmt.Printf("Exploring %s...\n", location_list.Name)
	fmt.Println("Found Pokemon: ")
	for _,enc := range location_list.Pokemon_encounters{
		fmt.Printf("%s\n",enc.Pokemon.Name)
	}
	return nil
}
