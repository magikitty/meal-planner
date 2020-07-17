package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

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

// GetRandomMeal returns a random meal from AllMeals struct
func GetRandomMeal(allMeals AllMeals) (meal Meal) {
	rand.Seed(time.Now().UnixNano())
	max := len(allMeals.Meals)
	randomNum := rand.Intn(max)
	return allMeals.Meals[randomNum]
}
