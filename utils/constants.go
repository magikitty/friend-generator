package utils

// MenuMainOptions contains messages and options for the main menu
var MenuMainOptions = map[string]string{
	"1": generateFriend,
	"2": quit,
}

// Messages for the main menu
var generateFriend = "Generate Friend"
var quit = "Quit"

// MenuInstructions explains how to make a selection in the menu
const MenuInstructions = "\n--------------------------------------------------------\nWhat do you want to do? Press the number of your choice.\n"

// MessageMenuWrongInput if user did not press number corresponding to menu item
const MessageMenuWrongInput = "Press the number of a menu item."

// MessageSelectionUser for showing the user what they selected
var MessageSelectionUser = "You selected:"

// MessageWelcome greets user when they start the application
const MessageWelcome string = "****** (^ ^) Welcome to the Friend Generator! (^ ^) ******\n"

// MenuGeneratorOptions contains the options for generating a friend
var MenuGeneratorOptions = map[string]string{
	"1": randomFriend,
	"2": customFriend,
}

var randomFriend = "Generate random friend"
var customFriend = "Generate custom friend"
