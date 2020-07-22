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
	// replace with new getInputAsInt() return value (durationInput, err)
	durationInput := GetUserInput()
	durationInputInt, err := strconv.Atoi(durationInput)
	if err != nil {
		return GetMealPlanDuration()
	}
	return durationInputInt
}

func checkDurationValid(mealPlanDuration int, totalPortions int) (durationValid bool) {
	durationValid = false
	if mealPlanDuration > totalPortions {
		fmt.Println(MenuMessages().DurationNotValid)
	} else {
		durationValid = true
	}
	return durationValid
}

// getSumAllMealsPortions returns the sum of all the portions of all the meals combined
func getSumAllMealsPortions(meals AllMeals) (sumAllPortions int) {
	sumAllPortions = 0
	for _, meal := range meals.Meals {
		sumAllPortions += meal.Portions
	}
	return sumAllPortions
}

// MakeMealPlan returns a collection of length duration containing randomly picked meals
func MakeMealPlan(duration int, mealData AllMeals) []Meal {
	mealPlan := []Meal{}
	for i := 0; i < duration; {
		meal := GetRandomMeal(mealData)
		duplicateMeal := checkDuplicateMeal(mealPlan, meal)

		if duplicateMeal == false {
			if meal.Portions > 1 && meal.Portions <= (duration-i) {
				mealCollectionPortions := addPortions(meal)
				mealPlan = append(mealPlan, mealCollectionPortions...)
				i = i + meal.Portions
			} else if meal.Portions == 1 {
				mealPlan = append(mealPlan, meal)
				i++
			}
		}

	}
	return mealPlan
}

// GetRandomMeal returns a random meal from AllMeals struct
func GetRandomMeal(allMeals AllMeals) (meal Meal) {
	rand.Seed(time.Now().UnixNano())
	max := len(allMeals.Meals)
	randomNum := rand.Intn(max)
	return allMeals.Meals[randomNum]
}

// checkDuplicateMeal checks if a meal is already in the meal plan and returns true if it is
func checkDuplicateMeal(mealPlan []Meal, meal Meal) (mealInPlan bool) {
	mealInPlan = false
	for i := 0; i < len(mealPlan); i++ {
		if mealPlan[i].Name == meal.Name {
			mealInPlan = true
		}
	}
	return mealInPlan
}

// Rename for dumb people like me who can't easily understand what it does

// addPortions returns a collection containing a meal the number of times equal to its portion size
func addPortions(meal Meal) []Meal {
	mealCollection := []Meal{}
	for portion := 1; portion <= meal.Portions; portion++ {
		fmt.Println(portion)
		mealCollection = append(mealCollection, meal)
	}
	return mealCollection
}
