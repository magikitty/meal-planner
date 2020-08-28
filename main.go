package main

import (
	"fmt"
	"meal-planner/src/utils"
)

// Remove unnecessary debugging comments
func main() {
	fmt.Println(utils.MenuMessages().WelcomeMessage)

	// Investigate replacing global variable with something else (b/c anything can assign values to this var and accidentally terminate your app)
	for utils.QuitProgram == false {
		utils.MainMenu()
	}
}
