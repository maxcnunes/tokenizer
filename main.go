package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
)

// OAuth2Service configuration
type OAuth2Service struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	GrantType    string `json:"grantType"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// Configuration contains a list of OAuth2Services
type Configuration struct {
	OAuth2Services []OAuth2Service `json:"oAuth2Services"`
}

func loadConfiguration() (config Configuration, err error) {
	config = Configuration{}

	usr, err := user.Current()
	if err != nil {
		return config, err
	}

	file, err := os.Open(usr.HomeDir + "/.tokenizer.json")
	if err != nil {
		return config, err
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func selectOAuth2ServiceName(config Configuration) (name string, err error) {
	fmt.Println("Select the oauth2 service:")
	for _, oAuth2Service := range config.OAuth2Services {
		fmt.Printf("\t> %s\n", oAuth2Service.Name)
	}
	fmt.Print(">>> ")

	_, err = fmt.Scanln(&name)
	return name, err
}

func selectOAuth2Service(config Configuration, name *string) *OAuth2Service {
	for _, oAuth2Service := range config.OAuth2Services {
		if *name == oAuth2Service.Name {
			return &oAuth2Service
		}
	}
	return nil
}

func request(oAuth2Service *OAuth2Service) (res *http.Response, err error) {
	urlEncoded := url.Values{}
	urlEncoded.Add("grant_type", oAuth2Service.GrantType)
	urlEncoded.Add("client_id", oAuth2Service.ClientID)
	urlEncoded.Add("client_secret", oAuth2Service.ClientSecret)
	res, err = http.PostForm(oAuth2Service.URL, urlEncoded)
	return res, err
}

func printResult(body io.Reader) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("Error reading body", err)
		return
	}

	var token map[string]interface{}
	if err := json.Unmarshal(data, &token); err != nil {
		panic(err)
	}

	result, _ := json.MarshalIndent(token, "", "\t")
	os.Stdout.Write(result)
}

func main() {
	oAuth2ServiceName := flag.String("name", "", "oauth2 service name")
	flag.Parse()

	http.DefaultTransport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	config, err := loadConfiguration()
	if err != nil {
		log.Fatal("Error reading config file", err)
	}

	if *oAuth2ServiceName == "" {
		if *oAuth2ServiceName, err = selectOAuth2ServiceName(config); err != nil {
			log.Fatal("Wrong option", err)
		}
	}

	oAuth2Service := selectOAuth2Service(config, oAuth2ServiceName)
	if oAuth2Service == nil {
		log.Fatal("Service configuration not found")
	}

	res, err := request(oAuth2Service)
	if err != nil {
		log.Fatal("Error posting", err)
	}

	defer res.Body.Close()
	printResult(res.Body)
}
