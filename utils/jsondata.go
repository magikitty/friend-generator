package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// FriendResponse represents parent slice of API call results
type FriendResponse struct {
	Results []Friend `json:"results"`
}

// FriendsCollection represents a collection of friends written to a file
type FriendsCollection struct {
	FriendsSlice []Friend `json:"friends"`
}

// Friend represents selected json data for a friend kept in the FriendsCollection struct
type Friend struct {
	Gender string `json:"gender"`
	Name   struct {
		Title string `json:"title"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	DOB struct {
		Date string `json:"date"`
		Age  int    `json:"age"`
	} `json:"dob"`
	Location struct {
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"location"`
}

// GetJSONFriendData returns JSON data of type FriendResponse struct from API call
func GetJSONFriendData(url string) FriendResponse {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var friendResponse FriendResponse
	err = json.NewDecoder(response.Body).Decode(&friendResponse)
	if err != nil {
		log.Fatal(err)
	}

	return friendResponse
}
