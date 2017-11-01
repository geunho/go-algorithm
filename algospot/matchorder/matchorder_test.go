package matchorder

import (
	"testing"
	"fmt"
)

func TestMaxwinCase1(t *testing.T) {
	enemies := []int{3000, 2700, 2800, 2200, 2500, 1900}
	alias := []int{2800, 2750, 2995, 1800, 2600, 2000}
	result := maxwin(alias, enemies)
	solution := 5

	fmt.Println(result)

	if result != solution {
		t.Error(
			"alias", alias,
			"enemies", enemies,
			"exepected", solution,
			"got", result,
		)
	}
}

func TestMaxwinCase2(t *testing.T) {
	enemies := []int{1, 2, 3}
	alias := []int{3, 2, 1,}
	result := maxwin(alias, enemies)
	solution := 3

	fmt.Println(result)

	if result != solution {
		t.Error(
			"alias", alias,
			"enemies", enemies,
			"exepected", solution,
			"got", result,
		)
	}
}

func TestMaxwinCase3(t *testing.T) {
	enemies := []int{2, 3, 4, 5}
	alias := []int{1, 2, 3, 4}
	result := maxwin(alias, enemies)
	solution := 3

	fmt.Println(result)

	if result != solution {
		t.Error(
			"alias", alias,
			"enemies", enemies,
			"exepected", solution,
			"got", result,
		)
	}
}