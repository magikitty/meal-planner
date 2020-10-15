package generatemeals

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/magikitty/meal-planner/src/utils"
	"github.com/magikitty/menu"
)

func getMealPlan() ([]Meal, error) {
	duration := getMealPlanDurationInput()
	return makeMealPlan(duration, utils.FilePaths().JSONMealsData), nil
}

// ensure get valid duration input from user
func getMealPlanDurationInput() (durationInputInt int) {
	fmt.Println(utils.MenuMessages().MealPlanDuration)
	durationInput := utils.ConvertStringToInt(menu.GetUserInput())
	if !durationValid(durationInput) {
		return getMealPlanDurationInput()
	}
	return durationInput
}

func durationValid(mealPlanDuration int) (durationValid bool) {
	if mealPlanDuration > 0 {
		return true
	}
	fmt.Println(utils.MenuMessages().InputNotValid)
	return false
}

func makeMealPlan(totalTargetDuration int, filePath string) []Meal {
	var mealPlan []Meal
	mealData := utils.GetFileData(filePath)
	allMeals := getAllMealsFromData(mealData).Meals
	for len(mealPlan) != totalTargetDuration {
		mealPlan = addRandomMealsToPlan(allMeals, mealPlan, totalTargetDuration)
	}

	return mealPlan
}

func addRandomMealsToPlan(allMeals, mealPlan []Meal, targetDuration int) []Meal {
	currentTargetDuration := targetDuration - len(mealPlan)

	for duration := 0; duration < currentTargetDuration; {
		randomNum := utils.GetRandomPositiveNumber(len(allMeals))
		randomMeal := allMeals[randomNum]

		if randomMeal.Portions == 1 {
			mealPlan = append(mealPlan, randomMeal)
			duration++
		} else if mealFitsPlan(randomMeal, currentTargetDuration-duration) {
			mealPlan = append(mealPlan, addAllPortionsOfMeal(randomMeal)...)
			duration += randomMeal.Portions
		}

		allMeals, _ = removeMeal(allMeals, randomNum)
		if len(allMeals) == 0 {
			break
		}
	}
	return mealPlan
}

func getAllMealsFromData(data []byte) AllMeals {
	var allMeals AllMeals
	err := json.Unmarshal(data, &allMeals)
	if err != nil {
		log.Fatal(err)
	}
	return allMeals
}

func mealFitsPlan(meal Meal, duration int) bool {
	if meal.Portions <= duration {
		return true
	}
	return false
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
func removeMeal(mealsSlice []Meal, mealIndex int) ([]Meal, error) {
	if mealIndex >= 0 && mealIndex <= len(mealsSlice)-1 {
		mealsSlice[mealIndex] = mealsSlice[len(mealsSlice)-1]
		return mealsSlice[:len(mealsSlice)-1], nil
	} else {
		return mealsSlice, utils.InvalidIndexToRemove
	}
}
