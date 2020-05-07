package utils

/*******
Messages
********/

// MenuInstructions explains how to make a selection in the menu
const MenuInstructions = `
--------------------------------------------------------
What do you want to do? Press the number of your choice.
`

// MessageQuit is a message for when the user quits the program
const MessageQuit = "You quit the program. Bye bye!"

// MessageWelcome greets user when they start the application
const MessageWelcome = "****** (^ ^) Welcome to the Friend Generator! (^ ^) ******"

/********************
Menu maps and options
*********************/

// MenuMainOptions contains messages and options for the main menu
var MenuMainOptions = map[string]string{
	"1": generateFriend,
	"2": quit,
}

// Options for the main menu
var generateFriend = "Generate Friend"
var quit = "Quit"

// MenuGeneratorOptions contains the options for generating a friend
var MenuGeneratorOptions = map[string]string{
	"1": randomFriend,
	"2": customFriend,
	"3": mainMenu,
}

// Options for the generator menu
var randomFriend = "Generate random friend"
var customFriend = "Generate custom friend"
var mainMenu = "Main menu"
