package quantization

import (
	"testing"
	"reflect"
)

func TestAverage(t *testing.T) {
	list := []int { 15, 6, 7 }
	solution := 9

	result := average(&list)

	if result != solution {
		t.Error(
			"For", list,
			"expected", solution,
			"result", result,
		)
	}
}

func TestSplit(t *testing.T) {
	list := []int { 15, 6, 7 }
	sleft := []int {6, 7}
	sright := []int {15}

	left, right := split(&list)

	err := func (solution *[]int, target *[]int) {
		t.Error(
			"For", list,
			"expected", *solution,
			"got", *target,
		)
	}

	if !reflect.DeepEqual(sleft, left) {
		err(&sleft, &left)
	}

	if !reflect.DeepEqual(sright, right) {
		err(&sright, &right)
	}
}