package commands

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/fvsantos-playground/boot-go-pokedex/internal/pokeapi"
)

func Catch(config *pokeapi.Config) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", config.Param)
	data, err := pokeapi.GetPokemonInfo(config.Param)
	if err != nil {
		return err
	}

	src := rand.NewSource(time.Hour.Nanoseconds())
	rand := rand.New(src)

	baseCatchRate := math.Ceil(float64(data.BaseExperience / 50))

	if rand.Intn(int(baseCatchRate)) == 0 {
		fmt.Printf("%s was caught!\n", config.Param)
	} else {
		fmt.Printf("%s escaped!\n", config.Param)
	}

	return nil
}
