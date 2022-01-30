package stats

import (
	"reflect"
	"testing"

	"github.com/Yessentemir/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_000},
		{ID: 2, Category: "food", Amount: 2_000_000},
		{ID: 3, Category: "auto", Amount: 3_000_000},
		{ID: 4, Category: "auto", Amount: 4_000_000},
		{ID: 5, Category: "fun", Amount: 5_000_000},
		{ID: 6, Category: "food", Amount: 4_000_000},
		{ID: 7, Category: "auto", Amount: 4_000_000},
		{ID: 8, Category: "fun", Amount: 3_000_000},
	}
	expected := map[types.Category]types.Money{
		"auto": 3_000_000,
		"food": 3_000_000,
		"fun":  4_000_000,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}

}
