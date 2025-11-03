package main

import "fmt"

func caught_pokemon(cfg *config, args ...string) error{
	fmt.Printf("Your Pokedex:\n")
	for _,pokemon := range cfg.caughtPokemon{
		fmt.Printf("  -%s \n",pokemon.Name)
	}
	return nil
}