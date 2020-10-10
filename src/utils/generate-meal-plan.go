package utils

import (
	"fmt"
)

func getMealPlan() ([]Meal, error) {
	duration := ensureMealPlanDurationInput()
	mealData := GetMealsFromFile(FilePaths().JSONMealsData)
	sumPortions := getSumMealsPortions(mealData)

	if durationValid(duration, sumPortions) {
		return makeMealPlan(duration, mealData.Meals), nil
	}
	return nil, CustomErrors().InvalidMealDuration
}

// ensureMealPlanDurationInput gets required length of meal plan from user
func ensureMealPlanDurationInput() (durationInputInt int) {
	fmt.Println(MenuMessages().MealPlanDuration)
	durationInput, err := GetInputAsInt()
	if err != nil {
		fmt.Printf(MenuMessages().InputNotValid)
		return ensureMealPlanDurationInput()
	}
	return durationInput
}

// return sum of portions of all meals
func getSumMealsPortions(meals AllMeals) (sumPortions int) {
	sumPortions = 0
	for _, meal := range meals.Meals {
		sumPortions += meal.Portions
	}
	return sumPortions
}

func durationValid(mealPlanDuration int, sumPortions int) (durationValid bool) {
	if mealPlanDuration <= sumPortions && mealPlanDuration > 0 {
		return true
	}
		fmt.Println(MenuMessages().DurationNotValid)
	return false
}

// makeMealPlan returns collection of random meals of duration length
func makeMealPlan(duration int, allMealsSlice []Meal) []Meal {
	var mealPlan []Meal
	for i := 0; i < duration; {
		randomNum := GetRandomPositiveNumber(len(allMealsSlice))
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

// addAllPortionsOfMeal returns collection of a meal of meal's portion property length
func addAllPortionsOfMeal(meal Meal) []Meal {
	var mealCollection []Meal
	for portion := 1; portion <= meal.Portions; portion++ {
		mealCollection = append(mealCollection, meal)
	}
	return mealCollection
}

// removeMeal returns copy of mealsSlice without specified item
func removeMeal(mealsSlice []Meal, mealIndex int) []Meal {
	mealsSlice[mealIndex] = mealsSlice[len(mealsSlice)-1]
	return mealsSlice[:len(mealsSlice)-1]
}
