package utils

// MenuMessages function returning struct object with menu messages
func MenuMessages() menuMessages {
	menuMessages := menuMessages{}
	menuMessages.ConfirmQuit = "Are you sure you want to quit? y/n"
	menuMessages.DisplayMealPlan = "\nHere is your meal plan:"
	menuMessages.DisplayPlanFormatting = "%v. Meal: %v\n   Ingredients: %v\n   Portions: %v \n"
	menuMessages.DurationNotValid = "You do not have enough recipes for that. Create a meal plan for less days or add more recipes."
	menuMessages.MealPlanDuration = "How many days do you want to create a meal plan for?"
	menuMessages.MealPlanCreationFailed = "Failed to create a meal plan. Please try again."
	menuMessages.MenuInstructions = "\nWhat do you want to do? Press the number of your choice."
	menuMessages.MenuMainOptions = MenuMainOptions
	menuMessages.QuitYes = "y"
	menuMessages.QuitNo = "n"
	menuMessages.WelcomeMessage = "Welcome to the Meal Planner!"
	return menuMessages
}

type menuMessages struct {
	ConfirmQuit            string
	DisplayMealPlan        string
	DisplayPlanFormatting  string
	DurationNotValid       string
	MealPlanDuration       string
	MealPlanCreationFailed string
	MenuInstructions       string
	MenuMainOptions        map[string]string
	QuitYes                string
	QuitNo                 string
	WelcomeMessage         string
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
	FilePaths.JSONMealsData = "./data/meals/meals.json"
	return FilePaths
}

type filePaths struct {
	JSONMealsData string
}
