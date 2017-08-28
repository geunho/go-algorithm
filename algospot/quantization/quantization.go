package quantization

import "fmt"


/*
Algospot
	Chapter 8.9 QUANTIZATION
	(https://algospot.com/judge/problem/read/QUANTIZE)
 */

const ProblemTitle = "Quantization Problem"

type Quantization struct {
	list []int
}

var problem = Quantization{}

func (p Quantization) GetProblemTitle() string {
	return ProblemTitle
}

func (p Quantization) ReadProblem() {
	var membersStr string
	fmt.Scan(&membersStr)
}

func (p Quantization) SolveProblem() interface{} {


	return
}

func average(list *[]int) int {
	sum := 0.0

	for _, elem := range *list {
		sum += float64(elem)
	}

	result := (sum / float64(len(*list))) + 0.5 /*rounds */

	return int(result)
}

func split(list *[]int) ([]int, []int) {
	left := []int{}
	right := []int{}

	avg := average(list)

	for _, elem := range *list {
		if elem <= avg {
			left = append(left, elem)
		} else {
			right = append(right, elem)
		}
	}

	return left, right
}