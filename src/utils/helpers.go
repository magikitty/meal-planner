package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/magikitty/menu"
)

// GetInputAsInt returns user input converted to int
func GetInputAsInt() (int, error) {
	inputString := menu.GetUserInput()
	inputInt, err := strconv.Atoi(inputString)
	return inputInt, err
}

// GetRandomPositiveNumber returns number between 0, max
func GetRandomPositiveNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
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
