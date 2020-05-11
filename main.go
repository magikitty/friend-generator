package main

import (
	"encoding/json"
	"fmt"
	"friend-generator/utils"
	"log"
	"net/http"
)

func main() {
	fmt.Println(utils.MessageWelcome)
	// menus.MenuMain()
	getJSONData()
}

// FriendResult is a struct for json data from search
type FriendResult struct {
	Gender string `json:"gender"`
}

// FriendInfo is a struct for info
type FriendInfo struct {
	Seed        string `json:"seed"`
	ResultsPage int    `json:"results"`
	Page        int    `json:"page"`
	Version     string `json:"version"`
}

// FriendResponse is a struct for top level json response data
type FriendResponse struct {
	Results []FriendResult `json:"results"`
	Info    FriendInfo
}

func getJSONData() {
	response, err := http.Get("https://randomuser.me/api/?inc=gender")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var friendResponse FriendResponse
	err = json.NewDecoder(response.Body).Decode(&friendResponse)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}

	fmt.Println("The json friendResponse is", friendResponse) // debugging
}
