package utils

import "fmt"

// PrintMessages prints messages with or without a number
func PrintMessages(messages []string, showNumber bool) {
	for index, message := range messages {
		if showNumber == true {
			fmt.Printf("%v. %s\n", index+1, message)
		} else {
			fmt.Printf("%s\n", message)
		}
	}
}
