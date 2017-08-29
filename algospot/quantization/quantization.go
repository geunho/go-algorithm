package quantization

import (
	"fmt"
	"reflect"
	"math"
)


/*
Algospot
	Chapter 8.9 QUANTIZATION
	(https://algospot.com/judge/problem/read/QUANTIZE)
 */

const ProblemTitle = "Quantization Problem"

type Quantization struct {
	numbers []int
	base int
}

var problem = Quantization{}

func (p Quantization) GetProblemTitle() string {
	return ProblemTitle
}

func (p Quantization) ReadProblem() {
	var size int
	fmt.Scanf("%d", &size)

	var numberOfBase int
	fmt.Scanf("%d", &numberOfBase)

	problem.base = numberOfBase

	numbers := make([]int, size)

	for i := 0; i < size; i++ {
		var n int
		fmt.Scanf("%d", &n)

		numbers[i] = n
	}

	problem.numbers = numbers
}

func (p Quantization) SolveProblem() interface{} {

	msd := math.MaxUint32
	for nbase := problem.base; nbase > 0; nbase-- {
		fmt.Println(nbase)
		fmt.Println(problem.numbers)
		calculated := quantize(problem.numbers, nbase)
		if calculated < msd {
			msd = calculated
		}
		fmt.Println(calculated)
	}

	return msd
}

func quantize(numbers []int, base int) int {

	groups := [][]int{}

	longer := &numbers

	for len(groups) < base {
		avg := average(longer)


		left, right := split(longer, avg)

		numberOfLefts, numberOfRights := len(left), len(right)

		arrays := append(groups, left, right)
		longer = findLongerArray(&arrays)

		findAndRemove(&groups, longer)

		if numberOfLefts == 0 || numberOfRights == 0 {
			groups = append(groups, *longer)
			break

		} else if numberOfLefts < numberOfRights {
			groups = append(groups, left)

		} else if numberOfLefts > numberOfRights {
			groups = append(groups, right)
		} else {
			groups = append(groups, left)
			groups = append(groups, right)
		}
	}

	msd := 0 // sum of mean square deviation
	for _, list := range groups {
		msd += deviation(&list)
	}

	return msd
}

// TODO: 값 비교가 최선은 아니겠지...
func findAndRemove(groups *[][]int, target *[]int) {
	for i, list := range *groups {
		if reflect.DeepEqual(list, *target) {
			*groups = append((*groups)[:i], (*groups)[i+1:]...)
		}
	}
}


func findLongerArray(arrays *[][]int) *[]int {
	llen := 0
	var larr *[]int

	for _, arr := range *arrays {
		if len(arr) > llen {
			llen = len(arr)
			larr = &arr
		}
	}

	return larr
}

// Mean Square Deviation
func deviation(list *[]int) int {
	avg := average(list)
	sum := 0

	for _, elem := range *list {
		dev := avg - elem
		sum += dev * dev
	}

	return sum
}


func average(list *[]int) int {
	sum := 0.0

	for _, elem := range *list {
		sum += float64(elem)
	}

	result := (sum / float64(len(*list))) + 0.5 /*rounds */

	return int(result)
}

func split(list *[]int, average int) ([]int, []int) {
	left := []int{}
	right := []int{}

	for _, elem := range *list {
		if elem <= average {
			left = append(left, elem)
		} else {
			right = append(right, elem)
		}
	}

	return left, right
}