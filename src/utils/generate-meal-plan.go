package utils

import "fmt"

// getMealPlanDuration returns number of days meal plan must last
func getMealPlanDuration() (durationInputInt int) {
	fmt.Println(MenuMessages().MealPlanDuration)
	durationInput, err := GetInputAsInt()
	if err != nil {
		return getMealPlanDuration()
	}
	return durationInput
}

// return sum of portions of all meals
func getSumAllMealsPortions(meals AllMeals) (sumAllPortions int) {
	sumAllPortions = 0
	for _, meal := range meals.Meals {
		sumAllPortions += meal.Portions
	}
	return sumAllPortions
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

// makeMealPlan returns collection of random meals of duration length
func makeMealPlan(duration int, allMealsSlice []Meal) []Meal {
	var mealPlan []Meal
	for i := 0; i < duration; {
		randomNum := getRandomMealIndex(allMealsSlice)
		randomMeal := allMealsSlice[randomNum]

		if randomMeal.Portions > 1 && randomMeal.Portions <= (duration-i) {
			mealPlan = append(mealPlan, addAllPortionsOfMeal(randomMeal)...)
			i = i + randomMeal.Portions
		} else if randomMeal.Portions == 1 {
			mealPlan = append(mealPlan, randomMeal)
			i++
		} else {
			fmt.Println(MenuMessages().MealPlanCreationFailed)
			MenuMain()
		}
		allMealsSlice = removeMeal(allMealsSlice, randomNum)
	}
	return mealPlan
}
