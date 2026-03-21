package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/fvsantos-playground/boot-go-pokedex/internal/pokeapi"
	"github.com/fvsantos-playground/boot-go-pokedex/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(time.Second * 10)
	pokeapi.SetCache(cache)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		command := words[0]
		cliCommand, ok := cliCommandMap[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if len(words) > 1 {
			cliCommand.config.Param = words[1]
		}

		err := cliCommand.callback(cliCommand.config)
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
