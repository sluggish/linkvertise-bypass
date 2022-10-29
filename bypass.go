package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type Response struct {
	Query string `json:"query"`
	Destination string `json:"destination"`

}

var url string

func bypass_lv() {
	b, err := http.Get("https://bypass.pm/bypass2?url="+url)

	if err != nil {
		fmt.Println(err)
	}

	defer b.Body.Close()

	body, _ := ioutil.ReadAll(b.Body)

	var response Response

	json.Unmarshal(body, &response)

	fmt.Printf("Linkvertise URL: %s", response.Query)
	fmt.Printf("\nResulting URL after bypass: %s", response.Destination)

}

func main() {
	fmt.Println("URL to bypass -> ")
	fmt.Scanln(&url)

	bypass_lv()
}