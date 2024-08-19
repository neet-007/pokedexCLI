package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/neet-007/pokecache"
)

type LocationsResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationRespone struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type Client struct {
	httpClient http.Client
	Cache      pokecache.Cache
}

type Config struct {
	Client   Client
	Previous *string
	Next     *string
}

const baseUrl = "https://pokeapi.co/api/v2"

func NewClient(interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		Cache: *pokecache.NewCache(interval),
	}
}

func (c *Client) ListLocations(pageUrl *string) (LocationsResponse, error) {
	endPoint := "/location-area/"
	url := baseUrl + endPoint
	if pageUrl != nil {
		url = *pageUrl
	}

	fmt.Println("cache hit")
	val, ok := c.Cache.Get(url)
	if ok {
		locations := LocationsResponse{}
		err := json.Unmarshal(val, &locations)

		if err != nil {
			log.Fatal(err)
		}

		return locations, nil
	}

	fmt.Println("cache miss")
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 399 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	c.Cache.Add(url, body)

	locations := LocationsResponse{}
	err = json.Unmarshal(body, &locations)

	if err != nil {
		log.Fatal(err)
	}

	return locations, nil
}

func (c *Client) ListLocation(locationName string) (LocationRespone, error) {
	endPoint := "/location-area/"
	url := baseUrl + endPoint + locationName

	fmt.Println("cache hit")
	val, ok := c.Cache.Get(url)
	if ok {
		location := LocationRespone{}
		err := json.Unmarshal(val, &location)

		if err != nil {
			log.Fatal(err)
		}

		return location, nil
	}

	fmt.Println("cache miss")
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 399 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	c.Cache.Add(url, body)

	location := LocationRespone{}
	err = json.Unmarshal(body, &location)

	if err != nil {
		log.Fatal(err)
	}

	return location, nil
}
