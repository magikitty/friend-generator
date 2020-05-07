package main

import (
	"fmt"
	"friend-generator/utils"

	"github.com/magikitty/menu"
)

func main() {
	fmt.Println(utils.MessageWelcome)
	menuMain()
}

func menuMain() {
	fmt.Println(utils.MenuInstructions)
	menu.PrintMenu(utils.MenuMainOptions, true)
	menuSelection := menu.GetMenuSelection(utils.MenuMainOptions)
	callMenuMainFunctions(menuSelection)
}

func menuGenerator() {
	fmt.Println(utils.MenuInstructions)
	menu.PrintMenu(utils.MenuGeneratorOptions, true)
	menu.GetMenuSelection(utils.MenuGeneratorOptions)
}

func callMenuMainFunctions(menuSelection string) {
	switch menuSelection {
	case utils.MenuMainOptions["1"]:
		menuGenerator()
	case utils.MenuMainOptions["2"]:
		fmt.Println("You quit the program. Bye bye!") // debugging
	default:
		fmt.Println("Default option, nothing else to do") // debugging
	}
}
