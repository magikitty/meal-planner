package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/magikitty/meal-planner/src/generatemeals"
	"github.com/magikitty/meal-planner/src/utils"
)

// Index handler for root directory
func Index(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, utils.MenuMessages().WelcomeMessage+"\n"+utils.MenuMessages().MenuInstructions)
	if err != nil {
		fmt.Print(err)
	}

	var s string
	menu := utils.MenuMainOptions
	for i := 1; i <= len(menu); i++ {
		s = fmt.Sprintf("%v. %s\n", i, menu[strconv.Itoa(i)])
		_, err = fmt.Fprint(w, s)
	}
}

// NewMealPlan handler for /new-meal-plan directory
func NewMealPlan(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	var stringMealPlan string
	mealPlan, err := generatemeals.GetMealPlan()
	stringMealPlan, err = generatemeals.StringMealPlan(mealPlan)
	_, err = fmt.Fprint(w, stringMealPlan)
	if err != nil {
		fmt.Print(err)
	}

}

// NewRecipe handler for /new-recipe
func NewRecipe(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, "We are going to add a new meal! Yay!")
	if err != nil {
		fmt.Print(err)
	}
}
