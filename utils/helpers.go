package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// CreateAndPopulateFile creates file if it does not exist
func CreateAndPopulateFile(filePath string, responseData FriendResponse) {
	_, err := os.Stat(filePath)
	if err != nil {

		_, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		friendsCollection := FriendsCollection{responseData.Results}
		WriteDataToFile(friendsCollection, filePath)

	} else {
		friendsFromFile := GetFriendsFromFile(filePath)
		friendsData := AppendFriends(responseData, friendsFromFile)
		WriteDataToFile(friendsData, filePath)
	}
}

// WriteDataToFile writes JSON data to a file
func WriteDataToFile(data FriendsCollection, filePath string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY, os.ModePerm)
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

// AppendFriends appends new friends to friends collection
func AppendFriends(responseData FriendResponse, friendsFromFile []Friend) FriendsCollection {
	friendData := responseData.Results[0]
	friends := FriendsCollection{friendsFromFile}

	friendsSlice := friends.FriendsSlice
	friendsSlice = append(friendsSlice, friendData)
	friends = FriendsCollection{friendsSlice}

	return friends
}

// GetFriendsFromFile returns a collection of FriendsCollection
func GetFriendsFromFile(filePath string) []Friend {
	friendsFromFile := ReadDataFromFile(filePath)
	friendsSlice := friendsFromFile.FriendsSlice
	return friendsSlice
}

// ReadDataFromFile reads and returns FriendsCollection data from JSON file
func ReadDataFromFile(filePath string) FriendsCollection {
	var friendsCollection FriendsCollection
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonData, &friendsCollection)
	if err != nil {
		log.Fatal(err)
	}

	return friendsCollection
}

// PrintNewFriendInfo prints out the information of the last friend in the file
func PrintNewFriendInfo() {
	friends := GetFriendsFromFile(JSONDataFilePath)
	friend := friends[len(friends)-1]
	fmt.Printf(Name, friend.Name.First, friend.Name.Last)
	fmt.Printf(Age, friend.DOB.Age)
	fmt.Printf(Gender, friend.Gender)
	fmt.Printf(Location, friend.Location.City, friend.Location.Country)
}
