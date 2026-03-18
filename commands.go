package main

import (
	"fmt"
	"os"

	"github.com/fvsantos-playground/boot-go-pokedex/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
	config      *Config
}

type Config struct {
	Next     string
	Previous string
}

var locationConfig *Config = &Config{}

var cliCommandMap = map[string]cliCommand{
	"map": {
		name:        "map",
		description: "",
		callback:    commandMap,
		config:      locationConfig,
	},
	"mapb": {
		name:        "mapb",
		description: "",
		callback:    commandMapB,
		config:      locationConfig,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
		config:      &Config{},
	},
	"help": {
		name:        "help",
		description: "Show help message",
		callback:    commandHelp,
		config:      &Config{},
	},
}

func printLocationData(config *Config, data pokeapi.PaginationData) {
	config.Next, config.Previous = data.Next, data.Previous
	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
}

func commandMap(config *Config) error {
	if len(config.Previous) != 0 && len(config.Next) == 0 {
		fmt.Println("You have reached the last page")
		return nil
	}

	data, err := pokeapi.GetLocations(config.Next)
	if err != nil {
		return err
	}

	printLocationData(config, data)
	return nil
}

func commandMapB(config *Config) error {
	if len(config.Previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}

	data, err := pokeapi.GetLocations(config.Previous)
	if err != nil {
		return err
	}

	printLocationData(config, data)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Println(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
