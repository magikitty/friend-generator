package main

import (
	"bufio"
	"fmt"
	"friend-generator/utils"
	"os"
	"strings"
)

func main() {
	fmt.Println(utils.MessageWelcome)
	menuMain()
}

func menuMain() {
	fmt.Println(utils.MenuInstructions)
	utils.PrintMenu(utils.MenuMainOptions, true)
	menuSelection := getMenuSelection(utils.MenuMainOptions)
	callMenuMainFunctions(menuSelection)
}

func menuGenerator() {
	fmt.Println(utils.MenuInstructions)
	utils.PrintMenu(utils.MenuGeneratorOptions, true)
	getMenuSelection(utils.MenuGeneratorOptions)
}

func callMenuMainFunctions(menuSelection string) {
	switch menuSelection {
	case utils.MenuMainOptions["1"]:
		menuGenerator()
	case utils.MenuMainOptions["2"]:
		fmt.Println("You quit the program. Bye bye!")
	default:
		fmt.Println("Default option, nothing else to do") // debugging
	}
}

func getMenuSelection(menu map[string]string) (menuSelection string) {
	getInput := true
	for getInput == true {
		input := getUserInput()
		_, ok := menu[input]
		if ok == true {
			getInput = false
			menuSelection = menu[input]
		} else {
			fmt.Println(utils.MessageMenuWrongInput)
		}
	}
	fmt.Println(utils.MessageSelectionUser, menuSelection)
	return menuSelection
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdout)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	inputTrimmed := strings.TrimSpace(userInput)
	return inputTrimmed
}
