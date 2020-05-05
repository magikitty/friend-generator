package main

import (
	"bufio"
	"fmt"
	"friend-generator/utils"
	"os"
)

func main() {
	fmt.Println(utils.MessageWelcome)
	fmt.Println(utils.MenuInstructions)
	utils.PrintMessages(utils.MenuMainMessages, true)
	getUserInput()
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdout)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return userInput
}
