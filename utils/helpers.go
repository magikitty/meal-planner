package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Move into magikitty/menu library or use pre-existing menu library function
// Update to return int from input string (durationInputInt, err := strconv.Atoi(durationInput) or err

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

// ReadDataFromFile reads JSON data from a file and returns data in AllMeals struct
func ReadDataFromFile(filePath string) AllMeals {
	var allMeals AllMeals
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonData, &allMeals)
	if err != nil {
		log.Fatal(err)
	}
	return allMeals
}

// Move into menus.go
// Improve Printf formatting
func displayMealPlan(mealPlan []Meal) {
	fmt.Println(MenuMessages().DisplayMealPlan)
	for i, meal := range mealPlan {
		fmt.Printf("%v. %v\n", i+1, meal)
	}
}
