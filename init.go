package main

import (
	"fmt"
	"strings"
)

var mConf = mapConfig{
	"https://pokeapi.co/api/v2/location?offset=0&limit=20",
	"",
}

var pokedex = map[string]pokemon{}

func parseInputItems(inputStr string) []string {
	inputItems := strings.Split(inputStr, " ")
	inputItems[len(inputItems)-1] = strings.ReplaceAll(inputItems[len(inputItems)-1], "\n", "")
	return inputItems
}

func executeCommand(inputItems []string) error {
	count := len(inputItems)
	if count == 0 {
		return fmt.Errorf("no input arguments")
	}
	if count == 1 {
		command := inputItems[0]
		return executeNoInputCommands(command)
	}
	if count > 1 {
		command := inputItems[0]
		args := inputItems[1:]
		return executeInputCommands(command, args)
	}
	return fmt.Errorf("no such command, try typing help to display list of available commands")
}

func executeInputCommands(command string, args []string) error {
	switch command {
	case "explore":
		{
			if len(args) > 1 {
				return fmt.Errorf("too much input arguments for explore command")
			}
			area := args[0]
			return executeExploreCommand(area)
		}
	case "catch":
		{
			if len(args) > 1 {
				return fmt.Errorf("too much input arguments for explore command")
			}
			name := args[0]
			return executeCatchCommand(name, pokedex)
		}
	case "inspect":
		{
			if len(args) > 1 {
				return fmt.Errorf("too much input arguments for explore command")
			}
			name := args[0]
			return executeInspectCommand(name, pokedex)
		}
	}
	return fmt.Errorf("no such command, try typing help to display list of available commands")
}

func executeNoInputCommands(command string) error {
	switch command {
	case "exit":
		return executeExitCommand()
	case "map":
		return executeMapCommand(&mConf)
	case "pmap":
		return executePMapCommand(&mConf)
	case "help":
		return executeHelpCommand()
	case "pokedex":
		return executePokedexCommand(pokedex)
	}
	return fmt.Errorf("no such command, try typing help to display list of available commands")
}
