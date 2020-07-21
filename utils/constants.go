package utils

// MenuMessages function returning struct object with menu messages
func MenuMessages() menuMessages {
	menuMessages := menuMessages{}
	menuMessages.WelcomeMessage = "Welcome to the Meal Planner!"
	menuMessages.MenuInstructions = "\nWhat do you want to do? Press the number of your choice."
	menuMessages.MenuMainOptions = MenuMainOptions
	menuMessages.MealPlanDuration = "How many days do you want to create a meal plan for?"
	menuMessages.DurationNotViable = "You do not have enough recipes for that. Create a meal plan for less days or add more recipes."
	menuMessages.ConfirmQuit = "Are you sure you want to quit? y/n"
	menuMessages.DisplayMealPlan = "\nHere is your meal plan:"
	menuMessages.QuitYes = "y"
	menuMessages.QuitNo = "n"
	return menuMessages
}

type menuMessages struct {
	WelcomeMessage    string
	MenuInstructions  string
	MenuMainOptions   map[string]string
	MealPlanDuration  string
	DurationNotViable string
	ConfirmQuit       string
	DisplayMealPlan   string
	QuitYes           string
	QuitNo            string
}

// MenuMainOptions contains options for the main menu
var MenuMainOptions = map[string]string{
	"1": "Generate a meal plan",
	"2": "Add a new meal",
	"3": "Quit",
}

// QuitProgram is a bool that is set to true when user wants to quit program
var QuitProgram = false

// FilePaths function returning struct object with filePaths
func FilePaths() filePaths {
	FilePaths := filePaths{}
	FilePaths.JSONMealsData = "./meals/meals.json"
	return FilePaths
}

type filePaths struct {
	JSONMealsData string
}
