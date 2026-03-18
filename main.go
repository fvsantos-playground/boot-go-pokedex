package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
		err := cliCommand.callback(cliCommand.config)
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
