package utils

import (
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

// GetFileData returns byte stream of data from file
func GetFileData(filePath string) []byte {
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}
