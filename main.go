package main

import (
	"fmt"
	"meal-planner/utils"
)

func main() {
	fmt.Println(utils.MenuMessages().WelcomeMessage)

	for utils.QuitProgram == false {
		utils.MainMenu()
	}
}
