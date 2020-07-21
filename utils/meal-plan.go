package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)


// getTotalPortions returns the total portions of all the meals combined
func getTotalPortions(meals AllMeals) (totalPortions int) {
	totalPortions = 0
	for _, meal := range meals.Meals {
		totalPortions += meal.Portions
	}
	return totalPortions
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

// GetRandomMeal returns a random meal from AllMeals struct
func GetRandomMeal(allMeals AllMeals) (meal Meal) {
	rand.Seed(time.Now().UnixNano())
	max := len(allMeals.Meals)
	randomNum := rand.Intn(max)
	return allMeals.Meals[randomNum]
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

// addPortions returns a collection containing a meal the number of times equal to its portion size
func addPortions(meal Meal) []Meal {
	mealCollection := []Meal{}
	for portion := 1; portion <= meal.Portions; portion++ {
		mealCollection = append(mealCollection, meal)
	}
	return mealCollection
}
