package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/magikitty/menu"
)

// GetInputAsInt returns an error and user input as an int
func GetInputAsInt() (int, error) {
	inputString := menu.GetUserInput()
	inputInt, err := strconv.Atoi(inputString)
	return inputInt, err
}

// ReadDataFromFile reads JSON data from a file and returns data in AllMeals struct
func ReadDataFromFile(filePath string) AllMeals {
	var allMeals AllMeals
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonData, &allMeals)
	if err != nil {
		log.Fatal(err)
	}
	return allMeals
}
