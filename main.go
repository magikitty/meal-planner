package main

import (
	"fmt"
	"meal-planner/utils"

	"github.com/magikitty/menu"
)

func main() {
	fmt.Println(utils.WelcomeMessage)
	mainMenu()
}

func mainMenu() {
	fmt.Println(utils.MenuInstructions)
	menu.PrintMenu(utils.MenuMainOptions, true)
}
