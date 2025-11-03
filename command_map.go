package main

import "fmt"




func commandMapf(cfg *config, args ...string) error{
	locationResp, err := cfg.pokeapiClient.ListLocation(cfg.NextLocation)
	if err!= nil{
		return err
	}

	cfg.NextLocation = locationResp.Next
	cfg.PreviousLocation = locationResp.Previous

	for _, loc:= range locationResp.Results{
		fmt.Printf("%s\n",loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error{
	locationResp, err := cfg.pokeapiClient.ListLocation(cfg.PreviousLocation)
	if err != nil {
		return err
	}
	cfg.NextLocation = locationResp.Next
	cfg.PreviousLocation = locationResp.Previous

	for _,loc := range locationResp.Results{
		fmt.Printf("%s\n",loc.Name)
	}
	return nil
}