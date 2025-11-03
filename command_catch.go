package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catch_pokemon(cfg *config, args ...string) error{
	if len(args) != 1{
		return errors.New("Error in file command_catch.go:Enter a valid Input\n")
	}

	name := args[0]
	pokemon,err := cfg.pokeapiClient.PokemonType(args[0])
	if err!= nil{
		return err
	}
	res := rand.Intn(pokemon.Base_experience)

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	if res>40{
		fmt.Printf("%s escaped!\n",name)
		return nil
	}else{
		cfg.caughtPokemon[name] = pokemon
		fmt.Printf("%s was caught!\n", name)
		return nil
	}
}