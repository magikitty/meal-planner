package utils

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// ConvertStringToInt rconverts string to int
func ConvertStringToInt(input string) (int, error) {
	inputInt, err := strconv.Atoi(input)
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
