package main

import (
	"testing"
	"fmt"
)

func TestConvertBirthOrder(t *testing.T) {
	result := convertBirthOrder("XLVII")
	solution := 47

	fmt.Println(result)

	if result != solution {
		t.Error(
			"expected", solution,
			"result", result,
		)
	}
}

func TestCompareName(t *testing.T) {
	result := compareName("Abcde", "Accce")
	solution := true

	fmt.Println(result)

	if result != solution {
		t.Error(
			"expected", solution,
			"result", result,
		)
	}
}

func TestGetSortedList(t *testing.T) {
	result := getSortedList([]string{"2", "Philippe II", "Louis IX"})

	fmt.Println(result)
}