package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GetMealPlanDuration returns number of days meal plan should last
func GetMealPlanDuration() (durationInputInt int) {
	fmt.Println(MenuMessages().MealPlanDuration)
	durationInput, err := GetInputAsInt()
	if err != nil {
		return GetMealPlanDuration()
	}
	return durationInput
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
func MakeMealPlan(duration int, allMealsSlice []Meal) []Meal {
	var mealPlan []Meal
	for i := 0; i < duration; {
		fmt.Println("1......... All meals are:", allMealsSlice) // debugging
		randomNum := GetRandomMealIndex(allMealsSlice)
		meal := allMealsSlice[randomNum]
		fmt.Println("2......... Got meal:", meal) // debugging

		if meal.Portions > 1 && meal.Portions <= (duration-i) {
			mealCollectionPortions := addAllPortionsOfMeal(meal)
			mealPlan = append(mealPlan, mealCollectionPortions...)
			i = i + meal.Portions
		} else if meal.Portions == 1 {
			mealPlan = append(mealPlan, meal)
			i++
		} else {
			fmt.Println(MenuMessages().MealPlanCreationFailed)
			//MainMenu()
		}

		removeMeal(allMealsSlice, randomNum)
	}
	return mealPlan
}

// GetRandomMealIndex returns an int for a random meal index from a slice of meals
func GetRandomMealIndex(allMealsSlice []Meal) int {
	rand.Seed(time.Now().UnixNano())
	max := len(allMealsSlice)
	return rand.Intn(max)
}

// removeMeal removes an item from a mealsSlice and returns the slice without the removed item
func removeMeal(mealsSlice []Meal, mealIndex int) []Meal {
	mealsSlice[mealIndex] = mealsSlice[len(mealsSlice)-1]
	return mealsSlice[:len(mealsSlice)-1]
}

// addAllPortionsOfMeal returns a collection containing a meal the number of times equal to its portion size
func addAllPortionsOfMeal(meal Meal) []Meal {
	var mealCollection []Meal
	for portion := 1; portion <= meal.Portions; portion++ {
		mealCollection = append(mealCollection, meal)
	}
	return mealCollection
}
