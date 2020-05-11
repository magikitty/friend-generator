package main

import (
	"encoding/json"
	"fmt"
	"friend-generator/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println(utils.MessageWelcome)
	// menus.MenuMain()
	data := getJSONData()
	writeDataToFile(data)
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
	Info    FriendInfo     `json:"info"`
}

func getJSONData() FriendResponse {
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

	fmt.Println("The json friendResponse is", friendResponse)    // debugging
	fmt.Println("The json gender is", friendResponse.Results[0]) // debugging
	fmt.Println("The json page is", friendResponse.Info.Page)    // debugging
	return friendResponse
}

func writeDataToFile(data FriendResponse) {
	file, _ := os.OpenFile("./jsonresponses/test.json", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(data)
}
