package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetUserInput returns user input without leading or trailing white space
func GetUserInput() string {
	reader := bufio.NewReader(os.Stdout)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	inputTrimmedSpaces := strings.TrimSpace(userInput)
	return inputTrimmedSpaces
}

// GetMealPlanDuration returns number of days meal plan should last
func GetMealPlanDuration() (durationInputInt int) {
	fmt.Println(MenuMessages().MealPlanDuration)
	durationInput := GetUserInput()
	durationInputInt, err := strconv.Atoi(durationInput)
	if err != nil {
		GetMealPlanDuration()
	}
	return durationInputInt
}
