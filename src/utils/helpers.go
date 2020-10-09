package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/magikitty/menu"
)

// GetInputAsInt returns user input converted to int
func GetInputAsInt() (int, error) {
	inputString := menu.GetUserInput()
	inputInt, err := strconv.Atoi(inputString)
	return inputInt, err
}

// ReadDataFromFile returns unmarshalled AllMeals struct from specified JSON file
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
