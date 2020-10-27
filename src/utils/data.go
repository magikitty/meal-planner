package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// GetFileData returns byte stream of data from file
func GetFileData(filePath string) []byte {
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}

// GetAllMealsFromData unmarshals and returns JSON data as a struct of all meals
func GetAllMealsFromData(data []byte) AllMeals {
	var allMeals AllMeals
	err := json.Unmarshal(data, &allMeals)
	if err != nil {
		log.Fatal(err)
	}
	return allMeals
}
