package menus

import (
	"errors"
	"fmt"
	"friend-generator/utils"

	"github.com/magikitty/menu"
)

// MenuMain calls the functions for displaying and getting input for main menu
func MenuMain() {
	fmt.Println(utils.MenuInstructions)
	menu.PrintMenu(utils.MenuMainOptions, true)
	menuSelection := menu.GetMenuSelection(utils.MenuMainOptions)
	callMenuMainFunctions(menuSelection)
}

func callMenuMainFunctions(menuSelection string) {
	switch menuSelection {
	case utils.MenuMainOptions["1"]:
		menuGenerateFriend()
	case utils.MenuMainOptions["2"]:
		quit()
	default:
		fmt.Println(errors.New(utils.ErrorInvalidSelection))
	}
}

func quit() {
	fmt.Println(utils.MessageQuit)
}
