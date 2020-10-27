package utils

import (
	"errors"
)

/************************************************************************
* Struct initialiser functions return structs with page-specific values	*
*	- Structs populated by values stored in maps						*
*	- Structs passed to handlers' html.Execute functions for passing to	*
*		pages' HTML														*
*************************************************************************/

// GetConstantsIndex returns struct with constant values for Home page
func GetConstantsIndex() ConstantsIndex {
	c := ConstantsIndex{}
	c.AddressIndex = PageAddresses()["index"]
	c.AddressNewPlan = PageAddresses()["newPlan"]
	c.AddressNewRecipe = PageAddresses()["newRecipe"]
	c.MessageWelcome = MessagesIndex()["welcome"]
	c.MessageInstructions = MessagesIndex()["instructions"]
	c.ButtonOptionsNewPlan = MessagesIndex()["optionNewPlan"]
	c.ButtonOptionsNewRecipe = MessagesIndex()["optionNewRecipe"]
	c.NavIndex = MessagesGlobal()["index"]
	c.TitlePage = MessagesGlobal()["index"]
	c.TitleApp = MessagesGlobal()["nameApp"]
	return c
}

// ConstantsIndex struct for Home page constants
type ConstantsIndex struct {
	AddressIndex           string
	AddressNewPlan         string
	AddressNewRecipe       string
	MessageWelcome         string
	MessageInstructions    string
	ButtonOptionsNewPlan   string
	ButtonOptionsNewRecipe string
	NavIndex               string
	TitlePage              string
	TitleApp               string
}

// GetConstantsNewPlan returns struct with constant values for New Plan page
func GetConstantsNewPlan(mealPlan []MealStringified, mealPlanString []string) ConstantsNewPlan {
	c := ConstantsNewPlan{}
	c.AddressIndex = PageAddresses()["index"]
	c.DurationMealPlan = MessagesNewMealPlan()["durationMealPlan"]
	c.NavIndex = MessagesGlobal()["index"]
	c.TitlePage = MessagesGlobal()["titleNewPlan"]
	c.TitleApp = MessagesGlobal()["nameApp"]
	// TODO: Remove once implemented string collection rendering in HTML
	c.MealPlan = mealPlan
	c.MealPlanString = mealPlanString
	return c
}

// ConstantsNewPlan struct for New Plan page constants
type ConstantsNewPlan struct {
	AddressIndex     string
	DurationMealPlan string
	NavIndex         string
	TitlePage        string
	TitleApp         string
	// TODO: Remove once implemented string collection rendering in HTML
	MealPlan       []MealStringified
	MealPlanString []string
}

// GetConstantsNewRecipe returns struct with constant values for New Recipe page
func GetConstantsNewRecipe() ConstantsNewRecipe {
	c := ConstantsNewRecipe{}
	c.AddressIndex = PageAddresses()["index"]
	c.MessageNewRecipe = MessagesNewRecipe()["messageNewRecipe"]
	c.NavIndex = MessagesGlobal()["index"]
	c.TitlePage = MessagesGlobal()["titleNewRecipe"]
	c.TitleApp = MessagesGlobal()["nameApp"]
	return c
}

// ConstantsNewRecipe struct for New Recipe page constants
type ConstantsNewRecipe struct {
	AddressIndex     string
	MessageNewRecipe string
	NavIndex         string
	NameApp          string
	TitlePage        string
	TitleApp         string
}

/********************************************************************************
* Map initialiser functions return maps containing values, acting as constants	*
* 	- Maps' values used both:													*
* 		- In backend, e.g. as paths to files and custom errors					*
* 		- In frontend, e.g. as text displayed to user and page metadata			*
*********************************************************************************/

// FilePaths initialiser function returns map of all file paths
func FilePaths() map[string]string {
	var filePaths = map[string]string{
		"dataMeal":      "./data/meals/meals.json",
		"pageIndex":     "./html/home.html",
		"pageNewRecipe": "./html/newRecipe.html",
		"pageNewPlan":   "./html/newMealPlan.html",
	}
	return filePaths
}

// MessagesIndex initialiser function returns map of Home page messages
func MessagesIndex() map[string]string {
	var messagesIndex = map[string]string{
		"welcome":         "Welcome to the Meal Planner!",
		"instructions":    "What do you want to do?",
		"optionNewPlan":   "Make a new meal plan",
		"optionNewRecipe": "Add a new recipe",
	}
	return messagesIndex
}

// MessagesNewMealPlan initialiser function returns map of New Plan page messages
func MessagesNewMealPlan() map[string]string {
	var messagesNewMealPlan = map[string]string{
		"durationMealPlan": "How many days do you want to make a meal plan for?",
		"invalidDuration":  "Invalid meal plan duration. Please try again.",
	}
	return messagesNewMealPlan
}

// MessagesNewRecipe initialiser function returns map of New Recipe page messages
func MessagesNewRecipe() map[string]string {
	var messagesNewRecipe = map[string]string{
		"messageNewRecipe": "We are going to add a new meal! Yay!",
	}
	return messagesNewRecipe
}

// MessagesGlobal initialiser function returns map of global page messages
func MessagesGlobal() map[string]string {
	var messagesGlobal = map[string]string{
		"index":          "Home",
		"nameApp":        "Meal Planner",
		"titleNewRecipe": "New Recipe",
		"titleNewPlan":   "New Meal Plan",
	}
	return messagesGlobal
}

// PageAddresses initialiser function returns map of all page addresses
func PageAddresses() map[string]string {
	var pageAddresses = map[string]string{
		"index":     "/",
		"newPlan":   "/new-plan",
		"newRecipe": "/new-recipe",
	}
	return pageAddresses
}

// CustomErrors initialiser function returns map of all custom errors
func CustomErrors() map[string]error {
	var customErrors = map[string]error{
		"invalidIndexToRemove":    errors.New("cannot remove item as index invalid"),
		"stringifyMealPlanFailed": errors.New("cannot stringify meal plan"),
	}
	return customErrors
}

// GetFormatStrings initialization function returns map of string formatting values
func GetFormatStrings() map[string]string {
	var formatStrings = map[string]string{
		"day":           "Day",
		"portionSize":   "Portion size",
		"space":         " ",
		"colon":         ":",
		"bracketOpen":   "(",
		"bracketClosed": ")",
	}
	return formatStrings
}
