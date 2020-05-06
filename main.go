package main

import (
	"bufio"
	"fmt"
	"friend-generator/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(utils.MessageWelcome)
	fmt.Println(utils.MenuInstructions)
	utils.PrintMessages(utils.MenuMainMessages, true)
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

func getMenuSelection(menu []string) string {
	input := getUserInput()
	inputInt, err := strconv.Atoi(input)
	if err != nil || inputInt > len(menu) {
		fmt.Println("Please select an option from the menu")
		getUserInput()
	}
	menuSelection := inputInt - 1
	fmt.Println("menu selection is", menuSelection) // debugging
	fmt.Println(menu[menuSelection])                // debugging
	return menu[menuSelection]
}

