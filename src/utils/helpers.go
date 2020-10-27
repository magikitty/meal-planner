package utils

import (
	"math/rand"
	"strconv"
	"time"
)

// ConvertStringToInt rconverts string to int
func ConvertStringToInt(input string) int {
	inputInt, _ := strconv.Atoi(input)
	return inputInt
}

// GetRandomPositiveNumber returns number between 0, max
func GetRandomPositiveNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// IsLargerThanZero returns if provided number is larger than zero
func IsLargerThanZero(mealPlanDuration int) (isLargerThanZero bool) {
	if mealPlanDuration > 0 {
		return true
	}
	return false
}

// NumberFitsLimit checks provided value less than or equal to provided limit
func NumberFitsLimit(number, limit int) bool {
	if number <= limit && number > 0 {
		return true
	}
	return false
}
