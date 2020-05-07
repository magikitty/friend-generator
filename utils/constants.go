package utils

// MenuMainMessages contains messages and options for the main menu
var MenuMainMessages = map[string]string{
	"1": generateFriend,
	"2": quit,
}

// Messages for the main menu
var generateFriend = "Generate Friend"
var quit = "Quit"

// MenuInstructions explains how to make a selection in the menu
const MenuInstructions = "What do you want to do? Press the number of your choice.\n"

// MenuMessageWrongInput if user did not press number corresponding to menu item
const MenuMessageWrongInput = "Press the number of a menu item."

// SelectionUser for showing the user what they selected
var SelectionUser = "You selected:"

// MessageWelcome greets user when they start the application
const MessageWelcome string = "****** (^ ^) Welcome to the Friend Generator! (^ ^) ******\n\n"
