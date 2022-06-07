package stats

import (
	"reflect"
	"testing"

	"github.com/ShavqatKavrakov/Lesson11_bank/v2/pkg/types"
)

func TestCategoriesAvg_empty(t *testing.T) {
	payments := []types.Payment{}
	result := CategoriesAvg(payments)
	if len(result) != 0 {
		t.Error("result len !=0")
	}
}

func TestCategoriesAvg_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 3_000_00,
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, result: %v", expected, result)
	}
}
func TestCategoriesAvg_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 4_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 5_000_00},
		{ID: 6, Category: "fun", Amount: 7_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 3_000_00,
		"food": 2_000_00,
		"fun":  6_000_00,
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, result: %v", expected, result)
	}
}
