package utils

import (
	"math/rand"
	"time"
)

// getRandomMealIndex returns random valid index of slice
func getRandomMealIndex(allMealsSlice []Meal) int {
	rand.Seed(time.Now().UnixNano())
	max := len(allMealsSlice)
	return rand.Intn(max)
}

// removeMeal returns copy of mealsSlice without specified item
func removeMeal(mealsSlice []Meal, mealIndex int) []Meal {
	mealsSlice[mealIndex] = mealsSlice[len(mealsSlice)-1]
	return mealsSlice[:len(mealsSlice)-1]
}

// addAllPortionsOfMeal returns collection of a meal of meal's portion property length
func addAllPortionsOfMeal(meal Meal) []Meal {
	var mealCollection []Meal
	for portion := 1; portion <= meal.Portions; portion++ {
		mealCollection = append(mealCollection, meal)
	}
	return mealCollection
}
