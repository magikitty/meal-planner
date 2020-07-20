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

// MakeMealPlan returns a collection of length duration containing random meals
func MakeMealPlan(duration int, mealData AllMeals) []Meal {
	mealPlan := []Meal{}
	for i := 0; i < duration; i++ {
		randomMeal := GetRandomMeal(mealData)
		duplicateMeal := checkDuplicateMeal(mealPlan, randomMeal)
		if duplicateMeal == false {
			mealPlan = append(mealPlan, randomMeal)
		} else {
			i--
		}
	}
	return mealPlan
}

// checkDuplicateMeal checks if a meal is already in the meal plan and returns true if it is
func checkDuplicateMeal(mealPlan []Meal, meal Meal) (mealInPlan bool) {
	mealInPlan = false
	for i := 0; i < len(mealPlan); i++ {
		if mealPlan[i].Name == meal.Name {
			mealInPlan = true
			fmt.Printf("checked if %v is the same as %v and returned %v\n", mealPlan[i].Name, meal.Name, mealInPlan) // debugging
		}
	}
	return mealInPlan
}

func addPortions(meal Meal) []Meal {
	mealCollection := []Meal{}
	for portion := 1; portion <= meal.Portions; portion++ {
		mealCollection = append(mealCollection, meal)
		fmt.Printf("added meal -- %v -- to plan\n", meal)
	}
	return mealCollection
}
