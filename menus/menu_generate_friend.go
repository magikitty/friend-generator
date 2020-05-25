package menus

import (
	"errors"
	"fmt"
	"friend-generator/utils"
	"strings"

	"github.com/magikitty/menu"
)

func menuGenerateFriend() {
	fmt.Println(utils.MenuInstructions)
	menu.PrintMenu(utils.MenuGenerateFriendOptions, true)
	menuSelection := menu.GetMenuSelection(utils.MenuGenerateFriendOptions)
	callMenuGenerateFriendFunctions(menuSelection)
}

func callMenuGenerateFriendFunctions(menuSelection string) {
	switch menuSelection {
	case utils.MenuGenerateFriendOptions["1"]:
		generateRandomFriend()
	case utils.MenuGenerateFriendOptions["2"]:
		generateCustomFriend()
	case utils.MenuGenerateFriendOptions["3"]:
		MenuMain()
	default:
		fmt.Println(errors.New(utils.ErrorInvalidSelection))
	}
}

func generateRandomFriend() {
	fmt.Println(utils.MessageNewFriend)
	responseData := utils.GetJSONFriendData(utils.API)
	utils.CreateAndPopulateFile(utils.JSONDataFilePath, responseData)
	utils.PrintNewFriendInfo()
	menuGenerateFriend()
}

func generateCustomFriend() {
	generateFriendGender()
}

func generateFriendGender() {
	fmt.Println(utils.MessageSelectGender)
	menu.PrintMenu(utils.MenuCustomGenderOptions, true)
	menuSelection := menu.GetMenuSelection(utils.MenuCustomGenderOptions)
	apiGender := utils.APIGender + strings.ToLower(menuSelection)
	responseData := utils.GetJSONFriendData(apiGender)
	utils.CreateAndPopulateFile(utils.JSONDataFilePath, responseData)
	fmt.Println(utils.MessageNewFriend)
	utils.PrintNewFriendInfo()
	menuGenerateFriend()
}
