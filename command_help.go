package main
import "fmt"

func help(cfg *config, args ...string) error{
	msg := `Welcome to the Pokedex!
Usage:
	
help: Displays a help message
exit: Exit the Pokedex
`
	fmt.Printf("%s",msg)
	return nil
}