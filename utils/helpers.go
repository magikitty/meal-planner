package utils

import (
	"bufio"
	"fmt"
	"os"
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
func GetMealPlanDuration() (durationInput string) {
	fmt.Println(MealPlanDuration)
	durationInput = GetUserInput()
	return durationInput
}

// func ValidateMealPlanDurationInput(durationInput string) (durationInputInt int) {
// 	inputValid := false
// 	for inputValid == false {
// 		durationInputInt, err := strconv.Atoi(durationInput)
// 		// fmt.Println("days:", durationInput)
// 		if err == nil {
// 			inputValid = true
// 		}
// 	}
// 	return durationInputInt
// }
