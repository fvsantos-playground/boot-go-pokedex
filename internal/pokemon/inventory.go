package pokemon

import (
	"iter"
	"maps"
	"sync"

	"github.com/fvsantos-playground/boot-go-pokedex/internal/pokeapi"
)

var inventory = map[string]pokeapi.PokemonInfo{}
var mu = &sync.Mutex{}

func Add(name string, pokemon pokeapi.PokemonInfo) {
	mu.Lock()
	defer mu.Unlock()

	inventory[name] = pokemon
}

func Get(name string) (pokemon pokeapi.PokemonInfo, exists bool) {
	mu.Lock()
	defer mu.Unlock()

	val, exists := inventory[name]
	if !exists {
		return pokeapi.PokemonInfo{}, false
	}

	return val, true
}

func GetNames() iter.Seq[string] {
	return maps.Keys(inventory)
}
