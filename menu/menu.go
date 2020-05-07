package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PrintMenu prints messages with or without a number
func PrintMenu(menu map[string]string, showNumber bool) {
	for i := 1; i <= len(menu); i++ {
		if showNumber == true {
			fmt.Printf("%v. %s\n", i, menu[strconv.Itoa(i)])
		} else {
			fmt.Printf("%s\n", menu[strconv.Itoa(i)])
		}
	}
}

// GetMenuSelection returns the map value corresponding to option (key) user selected
func GetMenuSelection(menu map[string]string) (menuSelection string) {
	getInput := true
	for getInput == true {
		input := GetUserInput()
		_, ok := menu[input]
		if ok == true {
			getInput = false
			menuSelection = menu[input]
		} else {
			fmt.Println(messageMenuInvalidInput)
		}
	}
	fmt.Println(messageSelectionUser, menuSelection)
	return menuSelection
}

// GetUserInput returns sanitized user input
func GetUserInput() string {
	reader := bufio.NewReader(os.Stdout)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	inputTrimmed := strings.TrimSpace(userInput)
	return inputTrimmed
}
