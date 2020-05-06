package utils

import (
	"fmt"
	"strconv"
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
