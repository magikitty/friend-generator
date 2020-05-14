package menus

import (
	"errors"
	"fmt"
	"friend-generator/utils"

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
		errors.New(utils.ErrorInvalidSelection)
	}
}

func generateCustomFriend() {
	fmt.Println(utils.MessageNewFriend)
}

func generateRandomFriend() {
	fmt.Println(utils.MessageNewFriend)
	responseData := utils.GetJSONFriendData(utils.API)
	utils.CreateAndPopulateFile(utils.JSONDataFilePath, responseData)
	utils.PrintNewFriendInfo()
}
