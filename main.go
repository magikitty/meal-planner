package main

import (
	"fmt"
	"github.com/magikitty/meal-planner/src/utils"
)

func main() {
	fmt.Println(utils.MenuMessages().WelcomeMessage)

	for true {
		utils.MainMenu()
	}
}
