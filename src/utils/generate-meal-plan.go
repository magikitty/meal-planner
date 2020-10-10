package utils

import (
	"fmt"
)

func getMealPlan() ([]Meal, error) {
	duration := ensureMealPlanDurationInput()

	if durationValid(duration) {
		return makeMealPlan(duration), nil
	}
	return nil, CustomErrors().InvalidMealDuration
}

// ensureMealPlanDurationInput gets required length of meal plan from user
func ensureMealPlanDurationInput() (durationInputInt int) {
	fmt.Println(MenuMessages().MealPlanDuration)
	durationInput, err := GetInputAsInt()
	if err != nil {
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

func durationValid(mealPlanDuration int) (durationValid bool) {
	if mealPlanDuration > 0 {
		return true
	}
	fmt.Println(MenuMessages().DurationNotValid)
	return false
}

func makeMealPlan(totalTargetDuration int) []Meal {
	var mealPlan []Meal

	for len(mealPlan) != totalTargetDuration {
		mealPlan = addRandomMealsToPlan(mealPlan, totalTargetDuration)
	}

	return mealPlan
}

func addRandomMealsToPlan(mealPlan []Meal, targetDuration int) []Meal {
	currentTargetDuration := targetDuration - len(mealPlan)
	mealData := GetMealsFromFile(FilePaths().JSONMealsData)
	allMeals := mealData.Meals

	for duration := 0; duration < currentTargetDuration; {
		randomNum := GetRandomPositiveNumber(len(allMeals))
		randomMeal := allMeals[randomNum]

		if randomMeal.Portions == 1 {
			mealPlan = append(mealPlan, randomMeal)
			duration++
		} else if mealFitsPlan(randomMeal, currentTargetDuration-duration) {
			mealPlan = append(mealPlan, addAllPortionsOfMeal(randomMeal)...)
			duration += randomMeal.Portions
		}

		allMeals = removeMeal(allMeals, randomNum)
		if len(allMeals) == 0 {
			break
		}
	}
	return mealPlan
}

func mealFitsPlan(meal Meal, duration int) bool {
	if meal.Portions <= duration {
		return true
	}
	return false
}

func mealPlanFailed() {
	fmt.Println(MenuMessages().MealPlanCreationFailed)
	MenuMain()
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
