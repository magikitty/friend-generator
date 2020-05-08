package menus

import (
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
		menuGenerator()
	case utils.MenuMainOptions["2"]:
		quit()
	default:
		fmt.Println("Default option, nothing else to do") // debugging
	}
}

func quit() {
	fmt.Println(utils.MessageQuit)
}
