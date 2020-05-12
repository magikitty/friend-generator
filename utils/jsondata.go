package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// FriendResponse is a struct for top level json response data from API call
type FriendResponse struct {
	Results []Friend `json:"results"`
}

// Friends is a struct containing a collection of friends
type Friends struct {
	FriendsCollection []Friend `json:"friends"`
}

// Friend is a struct containing selected json data for a friend
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
		Country string `json:"country"`
	} `json:"location"`
}

// GetJSONData returns JSON data of type FriendResponse struct from API call
func GetJSONData(url string) FriendResponse {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var friendResponse FriendResponse
	err = json.NewDecoder(response.Body).Decode(&friendResponse)
	if err != nil {
		log.Println(err)
	}

	return friendResponse
}
