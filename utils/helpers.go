package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// AppendFriends appends new friends to friends collection
func AppendFriends(responseData FriendResponse, friendsFromFile []Friend) Friends {
	friendData := responseData.Results[0]
	friends := Friends{friendsFromFile}

	friendsCollection := friends.FriendsCollection
	friendsCollection = append(friendsCollection, friendData)
	friends = Friends{friendsCollection}

	return friends
}

// GetFriendsFromFile is
func GetFriendsFromFile(fileName string) []Friend {
	friendsFromFile := ReadDataFromFile(fileName)
	friendsCollection := friendsFromFile.FriendsCollection
	return friendsCollection
}

// ReadDataFromFile reads and returns Friends data from JSON file
func ReadDataFromFile(jsonFile string) Friends {
	var friends Friends
	jsonData, _ := ioutil.ReadFile(jsonFile)
	err := json.Unmarshal(jsonData, &friends)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("We read and unmarshaled:", friends)                                               // debugging
	fmt.Println("We read and unmarshaled gender:", friends.FriendsCollection[0].Gender)            // debugging
	fmt.Println("We read and unmarshaled country:", friends.FriendsCollection[0].Location.Country) // debugging

	return friends
}

// WriteDataToFile writes JSON data to a file
func WriteDataToFile(data Friends, jsonFile string) {

	file, err := os.OpenFile(jsonFile, os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Println(err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}
