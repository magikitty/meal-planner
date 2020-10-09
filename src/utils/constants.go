package utils

// MenuMainOptions contains main menu options
var MenuMainOptions = map[string]string{
	"1": "Generate a meal plan",
	"2": "Add a new meal",
	"3": "Quit",
}

// MenuMessages returns struct with menu messages
func MenuMessages() menuMessages {
	menuMessages := menuMessages{}
	menuMessages.ConfirmQuit = "Are you sure you want to quit? y/n"
	menuMessages.DisplayMealPlan = "\nHere is your meal plan:"
	menuMessages.DisplayPlanFormatting = "%v. Meal: %v\n   Ingredients: %v\n   Portions: %v \n"
	menuMessages.DurationNotValid = "You do not have enough recipes for that. Create a meal plan for less days or add more recipes."
	menuMessages.MealPlanDuration = "How many days do you want to create a meal plan for?"
	menuMessages.MealPlanCreationFailed = "Couldn't find meals with suitable portion sizes for that number of days. Please try again with a different number of days."
	menuMessages.MenuInstructions = "\nWhat do you want to do? Press the number of your choice."
	menuMessages.MenuMainOptions = MenuMainOptions
	menuMessages.QuitYes = "y"
	menuMessages.QuitNo = "n"
	menuMessages.WelcomeMessage = "Welcome to the Meal Planner!"
	return menuMessages
}

// FilePaths returns struct with file paths
func FilePaths() filePaths {
	FilePaths := filePaths{}
	FilePaths.JSONMealsData = "./data/meals/meals.json"
	return FilePaths
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

type filePaths struct {
	JSONMealsData string
}
