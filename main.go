package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Configuration struct {
	URL          string
	GrantType    string
	ClientId     string
	ClientSecret string
}

func main() {
	file, err := os.Open("tokenizer.json")
	if err != nil {
		fmt.Println("Error reading config file", err)
		return
	}

	decoder := json.NewDecoder(file)
	config := Configuration{}

	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error parsing config", err)
		return
	}

	urlEncoded := url.Values{}
	urlEncoded.Add("grant_type", config.GrantType)
	urlEncoded.Add("client_id", config.ClientId)
	urlEncoded.Add("client_secret", config.ClientSecret)
	res, err := http.PostForm(config.URL, urlEncoded)

	if err != nil {
		fmt.Println("Error posting", err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading body", err)
		return
	}

	fmt.Printf("%s", body)
}
