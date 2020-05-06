package utils

import "fmt"

// PrintMenu prints messages with or without a number
func PrintMenu(menu map[int]string, showNumber bool) {
	for key, menuItem := range menu {
		if showNumber == true {
			fmt.Printf("%v. %s\n", key, menuItem)
		} else {
			fmt.Printf("%s\n", menuItem)
		}
	}
}
