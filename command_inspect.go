package main

import (
	"errors"
	"fmt"
	// "github.com/kryptonn36/pokedexcli/internal/pokeapi"
)

func inspect_pokemon(cfg *config, args ...string) error{
	if len(args) != 1{
		return errors.New("error in file command_inspect.go:Not a valid input")
	}
	name := args[0]
	res,ok := cfg.caughtPokemon[name]
	if ok{
		fmt.Printf("Name: %v \n",res.Name)
		fmt.Printf("Height: %v \n",res.Height)
		fmt.Printf("Weight: %v \n",res.Weight)

		fmt.Printf("Stats:\n")
		for _,stats := range res.Stats{
			fmt.Printf("  -%v: %v \n",stats.Stat.Name,stats.Base_stat)
		}

		fmt.Printf("Types:\n")
		for _,types := range res.Types{
			fmt.Printf("  -%v \n",types.Type.Name)
		}

		return nil
	}else{
		fmt.Printf("you have not caught that pokemon\n")
		return nil
	}
	
}