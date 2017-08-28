package quantization

import "fmt"


/*
Algospot
	Chapter 8.9 QUANTIZATION
	(https://algospot.com/judge/problem/read/QUANTIZE)
 */

const ProblemTitle = "Quantization Problem"

type Quantization struct {
	bases []int
	numbers []int
	rs map[int][]int
	// e.g.

}

var problem = Quantization{}

func (p Quantization) GetProblemTitle() string {
	return ProblemTitle
}

func (p Quantization) ReadProblem() {
	var numbersStr string
	fmt.Scan(&numbersStr)

	var numberOfBase int
	fmt.Scanf("%d", &numberOfBase)

	problem.bases = make([]int, numberOfBase)
	problem.numbers = transformNumberString(numbersStr)
}

func (p Quantization) SolveProblem() interface{} {


	return 0
}

func quantize(numbers []int, depth int) []([]int) {

	groups := make([][]int, depth)

	longer := numbers

	for i := 0; i < depth; i++ {
		avg := average(&longer)

		left, right := split(&longer, avg)

		numberOfLefts, numberOfRights := len(left), len(right)

		if numberOfLefts == 0 || numberOfRights == 0 {
			break
		}

		if numberOfLefts < numberOfRights {
			groups[i] = left
		} else {
			groups[i] = right
		}
	}


}

func findLongerArray(arrays ...*[]int) []int {
	llen := 0
	var larr []int

	for _, arr := range *arrays {
		if len(arr) > llen {
			llen = len(*arr)
			larr = *arr
		}
	}

	return larr
}

func transformNumberString(numberStr string) []int {


	return []int{}
}

// Mean Square Deviation
func deviation(list *[]int, average int) int {
	sum := 0

	for _, elem := range *list {
		dev := average - elem
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