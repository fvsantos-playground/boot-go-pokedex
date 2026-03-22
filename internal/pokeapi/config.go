package pokeapi

const BaseURL = "https://pokeapi.co/api/v2"
const LocationURL = BaseURL + "/location-area"
const PokemonURL = BaseURL + "/pokemon"

type Config struct {
	Next     string
	Previous string
	Param    string
}
