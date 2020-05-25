package utils

/**********
Messages
***********/

// MenuInstructions explains how to make a selection in the menu
const MenuInstructions = `
--------------------------------------------------------
What do you want to do? Press the number of your choice.
`

// MessageNewFriend displays after a new friend has been generated
const MessageNewFriend = "\n***********************\nHere's your new friend:\n"

// MessageSelectGender asks user to specify gender
const MessageSelectGender = "\nSelect a gender for your new friend.\n"

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

// MenuGenerateFriendOptions contains the options for generating a friend
var MenuGenerateFriendOptions = map[string]string{
	"1": randomFriend,
	"2": customFriend,
	"3": mainMenu,
}

// Options for the Generate Friend menu
var randomFriend = "Generate random friend"
var customFriend = "Generate custom friend"
var mainMenu = "Main menu"

// MenuCustomGenderOptions contains the options for generating a friend with a specified gender
var MenuCustomGenderOptions = map[string]string{
	"1": female,
	"2": male,
}

// Options for MenuCustomGenderOptions
var female = "Female"
var male = "Male"

/********************
APIs and file paths
*********************/

// API is the endpoint where API calls are made from
const API = "https://randomuser.me/api/"

// APIGender is the endpoint where calls are made for a specified gender (added to end of api call)
const APIGender = "https://randomuser.me/api/?gender="

// JSONDataFilePath is the path for the file where data is stored
var JSONDataFilePath = directory + fileName
var directory = "./jsonresponses/"
var fileName = "friends.json"

/**************
Error messages
***************/

// ErrorInvalidSelection for when user makes an invalid selection in a menu
const ErrorInvalidSelection = "invalid selection"

/***************
New Friend Info
****************/

// Name of new friend
var Name = "Name: %v %v\n"

// Age of new friend
var Age = "Age: %v\n"

// Gender of new friend
var Gender = "Gender: %v\n"

// Location of new friend
var Location = "Lives in: %v, %v\n"
