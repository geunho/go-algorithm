package lunchbox

import (
	"testing"
	"fmt"
)

func TestShortCase1(t *testing.T) {
	cook := []int{2, 2, 2}
	eat := []int{2, 2, 2}
	result := short(cook, eat)
	solution := 8

	fmt.Println(result)

	if result != solution {
		t.Error(
			"exepected", solution,
			"got", result,
		)
	}
}

func TestShortCase2(t *testing.T) {
	cook := []int{1, 2, 3}
	eat := []int{1, 2, 1}
	result := short(cook, eat)
	solution := 7

	fmt.Println(result)

	if result != solution {
		t.Error(
			"exepected", solution,
			"got", result,
		)
	}
}
