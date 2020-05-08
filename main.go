package main

import (
	"fmt"
	"friend-generator/menus"
	"friend-generator/utils"
)

func main() {
	fmt.Println(utils.MessageWelcome)
	menus.MenuMain()
}
