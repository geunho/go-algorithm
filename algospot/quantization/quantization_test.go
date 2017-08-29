package quantization

import (
	"testing"
	"reflect"
	"fmt"
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
	average := 9

	left, right := split(&list, average)

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

func TestDeviation(t *testing.T) {
	list := []int { 15, 6, 7 }
	solution := 49

	dev := deviation(&list)

	fmt.Println(dev)

	if dev != solution {
		t.Error(
			"For", list,
			"expected", solution,
			"got", dev,
		)
	}
}

func TestFindLongerArray(t *testing.T) {
	arrays := [][]int{
		{ 1, 2 },
		{ 4 },
		{ 7, 8, 9, 10 },
	}

	result := findLongerArray(&arrays)

	fmt.Println(*result)
}

func TestFindAndRemove(t *testing.T) {
	arrays := [][]int{
		{ 1, 2 },
		{ 4 },
		{ 7, 8, 9, 10 },
	}

	target := []int{ 1, 2 }

	findAndRemove(&arrays, &target)

	fmt.Println(arrays)
}

func TestQuantize(t *testing.T) {
	list := []int { 1, 744, 755, 4, 897, 902, 890, 6, 777 }
	depth := 3

	//list := []int { 3, 3, 3, 1, 2, 3, 2, 2, 2, 1 }
	//depth := 3

	result := quantize(&list, depth)

	fmt.Println(result)
}