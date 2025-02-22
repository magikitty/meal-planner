package endpoints

import (
	"html/template"
	"log"
	"net/http"

	"github.com/magikitty/meal-planner/src/generatemeals"
	"github.com/magikitty/meal-planner/src/ui"
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
	if err != nil {
		log.Fatal(err)
	}
}

// NewMealPlan handler for /new-meal-plan directory
func NewMealPlan(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	html, err := template.ParseFiles(utils.FilePaths()["pageNewPlan"])
	if err != nil {
		log.Fatal(err)
	}
	mealPlan, err := ui.StringifyMealPlan(generatemeals.GetMealPlan())
	if err != nil {
		log.Fatal(err)
	}
	mealString := utils.GetConstantsNewPlan(mealPlan, ui.ReformatMealStringPlan(mealPlan))
	err = html.Execute(w, mealString)
	if err != nil {
		log.Fatal(err)
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
