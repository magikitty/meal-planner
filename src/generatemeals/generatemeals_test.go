package generatemeals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMeals_mealFitsPlan(t *testing.T) {
	meal := Meal{"Soup", []string{"carrots, potatoes"}, 4}
	expectedBoolFits := true
	gotBoolFits := mealFitsPlan(meal, 5)

	expectedBoolDoesNotFit := false
	gotBoolDoesNotFit := mealFitsPlan(meal, 2)

	assert.Equal(t, expectedBoolFits, gotBoolFits, "Meal portions should be less than or equal to duration")
	assert.Equal(t, expectedBoolDoesNotFit, gotBoolDoesNotFit, "Meal portions should be less than or equal to duration")
}
