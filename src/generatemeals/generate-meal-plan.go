package generatemeals

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/magikitty/meal-planner/src/utils"
)

// GetMealPlan returns meal plan ready to be converterd to string for front end
func GetMealPlan() ([]Meal, error) {
	duration := getMealPlanDurationInput()
	return makeMealPlan(duration, utils.FilePaths()["dataMeal"]), nil
}

// TODO: allow user input via UX in frontend to set duration; currently set to 10
// ensure get valid duration input from user
func getMealPlanDurationInput() (durationInputInt int) {
	// fmt.Println(utils.MenuMessages().MealPlanDuration)
	// durationInput := utils.ConvertStringToInt(menu.GetUserInput())
	// if !durationValid(durationInput) {
	// 	return getMealPlanDurationInput()
	// }
	// return durationInput
	return 10
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
		// keep copy of allMeals data in memory instead of reading from file each loop
		allMealsCopy := make([]Meal, len(allMeals))
		copy(allMealsCopy, allMeals)

		mealPlan = addRandomMealsToPlan(allMealsCopy, mealPlan, totalTargetDuration)
	}

	return mealPlan
}

func addRandomMealsToPlan(allMeals, mealPlan []Meal, targetDuration int) []Meal {
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

func getAllMealsFromData(data []byte) AllMeals {
	var allMeals AllMeals
	err := json.Unmarshal(data, &allMeals)
	if err != nil {
		log.Fatal(err)
	}
	return allMeals
}

func mealFitsPlan(meal Meal, duration int) bool {
	if meal.PortionSize <= duration {
		return true
	}
	return false
}

// addAllPortionsOfMeal returns collection of a meal of meal's portion property length
func addAllPortionsOfMeal(meal Meal) []Meal {
	var mealCollection []Meal
	for portion := 1; portion <= meal.PortionSize; portion++ {
		mealCollection = append(mealCollection, meal)
	}
	return mealCollection
}

// removeMeal returns copy of mealsSlice without specified item
func removeMeal(mealsSlice []Meal, mealIndex int) ([]Meal, error) {
	if mealIndex < 0 || mealIndex > len(mealsSlice)-1 {
		log.Fatal(utils.ErrInvalidIndexToRemove)
	}
	mealsSlice[mealIndex] = mealsSlice[len(mealsSlice)-1]
	return mealsSlice[:len(mealsSlice)-1], nil
}
