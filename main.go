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
	fmt.Println(utils.MenuInstructions)
	utils.PrintMenu(utils.MenuMainMessages, true)
	getMenuSelection(utils.MenuMainMessages)
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

func getMenuSelection(menu map[string]string) (menuSelection string) {
	getInput := true
	for getInput == true {
		input := getUserInput()
		_, ok := menu[input]
		if ok == true {
			getInput = false
			menuSelection = menu[input]
		} else {
			fmt.Println(utils.MenuMessageWrongInput)
		}
	}
	fmt.Println(utils.SelectionUser, menuSelection)
	return menuSelection
}
