package main

import (
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
}

var commands = []cliCommand{
	{
		"help",
		"explore all commands",
	},
	{
		"map",
		"see next page of pokemon locations",
	},
	{
		"pmap",
		"see previous page of pokemon locations",
	},
	{
		"exit",
		"quit program",
	},
	{
		"explore [area]",
		"show pokemon encounters in a given area",
	},
}

func executeHelpCommand() error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage: ")
	for _, c := range commands {
		fmt.Printf("\t%s: %s\n", c.name, c.description)
	}
	return nil
}

func executeMapCommand(config *mapConfig) error {
	locations, err := fetchNextLocations(config)
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func executeExitCommand() error {
	os.Exit(0)
	return nil
}

func executePMapCommand(config *mapConfig) error {
	locations, err := fetchPrevLocations(config)
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func executeExploreCommand(area string) error {
	pokemons, err := fetchPokemonsFromArea(area)
	if err != nil {
		return err
	}
	for _, pokemon := range pokemons {
		fmt.Println(pokemon)
	}
	return nil
}

func executeCatchCommand(name string, usrPokedex map[string]pokemon) error {
	_, present := usrPokedex[name]
	if present {
		return fmt.Errorf("this pokemon is in your pokedex")
	}
	pokeInfo, err := fetchPokemonInfo(name)
	if err != nil {
		return err
	}
	fmt.Println("Throwing a Pokeball at " + name + "..")
	baseExperience := pokeInfo.BaseExperience
	chance := 0.0
	if baseExperience > 200 {
		chance = 0.1
	} else if baseExperience > 150 && baseExperience <= 200 {
		chance = 0.2
	} else if baseExperience > 100 && baseExperience <= 150 {
		chance = 0.3
	} else {
		chance = 0.4
	}
	if rand.Float64() < chance {
		fmt.Printf("%s was caught!\n", name)
		usrPokedex[name] = pokeInfo
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}

func executeInspectCommand(name string, usrPokedex map[string]pokemon) error {
	pokemonData, present := usrPokedex[name]
	if !present {
		return fmt.Errorf("you have not caught this pokemon")
	}
	fmt.Printf("Name: %s\n", pokemonData.Name)
	fmt.Printf("Height: %d\n", pokemonData.Height)
	fmt.Printf("Weight: %d\n", pokemonData.Weight)
	fmt.Println("Stats:")
	for _, statEntry := range pokemonData.Stats {
		fmt.Printf("  -%s: %d\n", statEntry.Stat.Name, statEntry.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeEntry := range pokemonData.Types {
		fmt.Printf("  - %s\n", typeEntry.Type.Name)
	}
	return nil
}

func executePokedexCommand(usrPokedex map[string]pokemon) error {
	fmt.Println("Your Pokedex:")
	for name := range usrPokedex {
		fmt.Printf("  - %s\n", name)
	}
	return nil
}
