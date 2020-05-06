package utils

// MenuMainMessages contains messages and options for the main menu
var MenuMainMessages = map[int]string{
	1: generateFriend,
	2: quit,
}

// MenuInstructions explains how to make a selection in the menu
const MenuInstructions = "What do you want to do? Press the number of your choice.\n"

// Messages for the main menu
var generateFriend = "Generate Friend"
var quit = "Quit"

// MessageWelcome greets user when they start the application
const MessageWelcome string = "****** (^ ^) Welcome to the Friend Generator! (^ ^) ******\n\n"
