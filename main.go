package main

import (
	"bufio"
	"fmt"
	"os"
)

const menuInstructions = "What do you want to do? Press the number of your choice."
const messageWelcome = "Welcome to the Friend Generator!\n\n"

func main() {
	fmt.Print(messageWelcome)
	menuMain()
}

func menuMain() {
	fmt.Println(menuInstructions)
	fmt.Println("1. Generate friend")
	fmt.Println("2. Quit")
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
