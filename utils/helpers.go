package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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
		return GetMealPlanDuration()
	}
	return durationInputInt
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

// GetRandomMeal returns a random meal from AllMeals struct
func GetRandomMeal(allMeals AllMeals) (meal Meal) {
	rand.Seed(time.Now().UnixNano())
	max := len(allMeals.Meals)
	randomNum := rand.Intn(max)
	return allMeals.Meals[randomNum]
}
