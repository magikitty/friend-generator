package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// WriteDataToFile writes JSON data to a file
func WriteDataToFile(data FriendResponse, jsonFile string) {
	friendData := data.Results[0]
	var friends Friends
	dataSlice := friends.FriendsCollection
	dataSlice = append(dataSlice, friendData)
	friends = Friends{dataSlice}

	file, err := os.OpenFile(jsonFile, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(friends, "", "    ")
	if err != nil {
		log.Println(err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}

// ReadDataFromFile reads and returns data from a JSON file
func ReadDataFromFile(jsonFile string) FriendResponse {
	var friendResponse FriendResponse
	jsonData, _ := ioutil.ReadFile(jsonFile)
	err := json.Unmarshal(jsonData, &friendResponse)
	if err != nil {
		log.Println(err)
	}

	return friendResponse
}
