package endpoints

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/magikitty/meal-planner/src/generatemeals"
	"github.com/magikitty/meal-planner/src/utils"
)

// Index handler for root directory
func Index(w http.ResponseWriter, _ *http.Request) {

	w.WriteHeader(http.StatusOK)
	html, err := template.ParseFiles(utils.FilePaths()["pageIndex"])
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, utils.GetConstantsIndex())
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
	html, err := template.ParseFiles(utils.FilePaths()["pageNewRecipe"])
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, utils.GetConstantsNewRecipe())
}
