package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

// GetJSONData returns JSON data of type FriendResponse struct
func GetJSONData() FriendResponse {
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
