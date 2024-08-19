module github.com/neet-007/pokedexcli/pokedexclimain

replace github.com/neet-007/pokeapi v0.0.0 => ../pokeapi

replace github.com/neet-007/pokecache v0.0.0 => ../pokecache

require (
	github.com/neet-007/pokeapi v0.0.0
	github.com/neet-007/pokecache v0.0.0
)

go 1.22.6
