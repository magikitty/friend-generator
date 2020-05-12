package utils

import (
	"encoding/json"
	"fmt"
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
