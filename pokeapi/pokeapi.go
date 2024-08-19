package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
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

type Client struct {
	httpClient http.Client
}

type Config struct {
	Client   Client
	Previous *string
	Next     *string
}

const baseUrl = "https://pokeapi.co/api/v2"

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) ListLocations(pageUrl *string) (LocationsResponse, error) {
	endPoint := "/location/"
	url := baseUrl + endPoint
	if pageUrl != nil {
		url = *pageUrl
	}

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

	locations := LocationsResponse{}
	err = json.Unmarshal(body, &locations)

	if err != nil {
		log.Fatal(err)
	}

	return locations, nil
}
