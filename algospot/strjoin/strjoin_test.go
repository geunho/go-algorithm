package strjoin

import (
	"testing"
	"fmt"
)

func TestMinCase1(t *testing.T) {
	result := min([]int{2, 2, 4})
	solution := 12

	fmt.Println(result)

	if result != solution {
		t.Error(
			"expected", solution,
			"got", result,
		)
	}
}

func TestMinCase2(t *testing.T) {
	result := min([]int{3, 1, 3, 4, 1})
	solution := 26

	fmt.Println(result)

	if result != solution {
		t.Error(
			"expected", solution,
			"got", result,
		)
	}
}

func TestMinCase3(t *testing.T) {
	result := min([]int{1, 1, 1, 1, 1, 1, 1, 2})
	solution := 27

	fmt.Println(result)

	if result != solution {
		t.Error(
			"expected", solution,
			"got", result,
		)
	}
}

func TestInsertElementAtSorttedSlice(t *testing.T) {
	slice := []int{1, 1, 4, 5, 8}
	elem := 3
	insertElementAtSorttedSlice(elem, &slice)

	fmt.Println(slice)

	if elem != slice[2] {
		t.Error(
			"expected", elem,
			"got", slice[2],
		)
	}

}

func TestDeleteFrontPairs(t *testing.T) {
	slice := []int{1, 1, 4, 5, 8}

	deleteFrontPairs(&slice)

	fmt.Println(slice)

	if slice[0] != 4 {
		t.Error(
			"expected", 4,
			"got", slice[0],
		)
	}
}