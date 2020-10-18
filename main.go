package main

import (
	"log"
	"net/http"
	"os"

	"github.com/magikitty/meal-planner/endpoints"
)

func main() {
	http.HandleFunc("/", endpoints.Index)
	http.HandleFunc("/new-plan", endpoints.NewMealPlan)
	http.HandleFunc("/new-recipe", endpoints.NewRecipe)

	log.Fatal(http.ListenAndServe(port(), nil))
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
