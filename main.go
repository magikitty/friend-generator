package main

import (
	"friend-generator/utils"
)

func main() {
	// fmt.Println(utils.MessageWelcome)
	// menus.MenuMain()

	data := utils.GetJSONData("https://randomuser.me/api/")
	friendsFromFile := utils.GetFriendsFromFile("./jsonresponses/test.json")
	friendsData := utils.AppendFriends(data, friendsFromFile)
	utils.WriteDataToFile(friendsData, "./jsonresponses/test.json")
}
