package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	urlEncoded := url.Values{}
	urlEncoded.Add("grant_type", "xxxxx")
	urlEncoded.Add("client_id", "xxxxx")
	urlEncoded.Add("client_secret", "xxxxx")
	res, err := http.PostForm("http://xxxx", urlEncoded)

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
