package generatemeals

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/magikitty/meal-planner/src/utils"
	"github.com/magikitty/menu"
)

// GetMealPlan returns meal plan ready to be converterd to string for front end
func GetMealPlan() ([]utils.Meal, error) {
	// duration := getMealPlanDurationInput()
	duration := 10
	return makeMealPlan(duration, utils.FilePaths()["dataMeal"]), nil
}

// TODO: Have user input duration in frontend
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

func makeMealPlan(totalTargetDuration int, filePath string) []utils.Meal {
	var mealPlan []utils.Meal
	mealData := utils.GetFileData(filePath)
	allMeals := getAllMealsFromData(mealData).Meals
	for len(mealPlan) != totalTargetDuration {
		// keep copy of allMeals data in memory instead of reading from file each loop
		allMealsCopy := make([]utils.Meal, len(allMeals))
		copy(allMealsCopy, allMeals)

		mealPlan = addRandomMealsToPlan(allMealsCopy, mealPlan, totalTargetDuration)
	}

	return mealPlan
}

func addRandomMealsToPlan(allMeals, mealPlan []utils.Meal, targetDuration int) []utils.Meal {
	currentTargetDuration := targetDuration - len(mealPlan)
	for duration := 0; duration < currentTargetDuration; {
		randomNum := utils.GetRandomPositiveNumber(len(allMeals))
		randomMeal := allMeals[randomNum]
		if randomMeal.PortionSize == 1 {
			mealPlan = append(mealPlan, randomMeal)
			duration++
		} else if mealFitsPlan(randomMeal, currentTargetDuration-duration) {
			mealPlan = append(mealPlan, addAllPortionsOfMeal(randomMeal)...)
			duration += randomMeal.PortionSize
		}

		allMeals, _ = removeMeal(allMeals, randomNum)
		if len(allMeals) == 0 {
			break
		}
	}
	return mealPlan
}

func getAllMealsFromData(data []byte) utils.AllMeals {
	var allMeals utils.AllMeals
	err := json.Unmarshal(data, &allMeals)
	if err != nil {
		log.Fatal(err)
	}
	return allMeals
}

func mealFitsPlan(meal utils.Meal, duration int) bool {
	if meal.PortionSize <= duration {
		return true
	}
	return false
}

// addAllPortionsOfMeal returns collection of a meal of meal's portion property length
func addAllPortionsOfMeal(meal utils.Meal) []utils.Meal {
	var mealCollection []utils.Meal
	for portion := 1; portion <= meal.PortionSize; portion++ {
		mealCollection = append(mealCollection, meal)
	}
	return mealCollection
}

// removeMeal returns copy of mealsSlice without specified item
func removeMeal(mealsSlice []utils.Meal, mealIndex int) ([]utils.Meal, error) {
	if mealIndex < 0 || mealIndex > len(mealsSlice)-1 {
		log.Fatal(utils.CustomErrors()["invalidIndexToRemove"])
	}
	mealsSlice[mealIndex] = mealsSlice[len(mealsSlice)-1]
	return mealsSlice[:len(mealsSlice)-1], nil
}
