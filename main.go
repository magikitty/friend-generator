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

func menuGenerator() {
	fmt.Println(utils.MenuInstructions)
	menu.PrintMenu(utils.MenuGeneratorOptions, true)
	menuSelection := menu.GetMenuSelection(utils.MenuGeneratorOptions)
	callMenuGeneratorFunctions(menuSelection)
}

func callMenuGeneratorFunctions(menuSelection string) {
	switch menuSelection {
	case utils.MenuGeneratorOptions["1"]:
		generateRandomFriend()
	case utils.MenuGeneratorOptions["2"]:
		generateCustomFriend()
	case utils.MenuGeneratorOptions["3"]:
		menuMain()
	default:
		fmt.Println("Default option, nothing else to do") // debugging
	}
}

func generateCustomFriend() {
	fmt.Println("Generating a custom friend for you...") // debugging
}

func generateRandomFriend() {
	fmt.Println("Generating a random friend for you...") // debugging
}
