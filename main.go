package main

import (
	"log"
	"net/http"
	"os"

	"github.com/magikitty/meal-planner/endpoints"
	"github.com/magikitty/meal-planner/src/utils"
)

func main() {
	http.HandleFunc(utils.PageAddresses()["index"], endpoints.Index)
	http.HandleFunc(utils.PageAddresses()["newPlan"], endpoints.NewMealPlan)
	http.HandleFunc(utils.PageAddresses()["newRecipe"], endpoints.NewRecipe)

	log.Fatal(http.ListenAndServe(port(), nil))
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
