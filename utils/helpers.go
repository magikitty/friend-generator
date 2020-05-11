package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// WriteDataToFile writes JSON data to a file
func WriteDataToFile(data FriendResponse) {
	file, err := os.OpenFile("./jsonresponses/test.json", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	jsondata, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(jsondata)
	if err != nil {
		fmt.Println(err)
	}
}

// ReadDataFromFile reads and returns data from a JSON file
func ReadDataFromFile() FriendResponse {
	var friendResponse FriendResponse
	jsondata, _ := ioutil.ReadFile("./jsonresponses/test.json")
	err := json.Unmarshal(jsondata, &friendResponse)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("We read and unmarshaled:", friendResponse)                                 // debugging
	fmt.Println("We read and unmarshaled gender result:", friendResponse.Results[0].Gender) // debugging

	return friendResponse
}
