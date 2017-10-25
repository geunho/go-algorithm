package quantization2

import (
	"fmt"
	"sort"
	"math"
)


/*
Algospot
	Chapter 8.9 QUANTIZATION
	(https://algospot.com/judge/problem/read/QUANTIZE)
 */

const ProblemTitle = "Quantization Problem"

type Quantization2 struct {
	numbers []int
	base int
	psum []int
	psqsum []int
	cache [][]int
}

var problem = Quantization2{}

func (p Quantization2) GetProblemTitle() string {
	return ProblemTitle
}

func (p Quantization2) ReadProblem() {
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

	initialize()
}

func initialize() {
	sort.Ints(problem.numbers)

	problem.cache = [][]int{}
	problem.psum = make([]int, len(problem.numbers))
	problem.psqsum = make([]int, len(problem.numbers))

	for i, number := range problem.numbers {
		if i == 0 {
			problem.psum[0] = number
			problem.psqsum[0] = number * number
		} else {
			problem.psum[i] = number + problem.psum[i - 1]
			problem.psqsum[i] = number * number + problem.psqsum[i - 1]
		}
	}

	for i := 0; i < len(problem.numbers); i ++ {
		row := make([]int, problem.base)
		for j := range row {
			row[j] = -1
		}
	}
}

func deviation(start int, end int) int {
	// devi of a to b
	// sum(x - avg)^2
	// sum(x^2 - 2*avg*x + avg^2)
	// sum(x^2) - 2*avg*sum(x) + avg^2*(b-a+1)
	// sqsum - 2*avg*sum + avg^2z
	length := end - start + 1
	sum, sqsum := 0, 0

	if start == 0 {
		sum = problem.psum[end]
		sqsum = problem.psqsum[end]
	} else {
		sum = problem.psum[end] - problem.psum[start-1]
		sqsum = problem.psqsum[end] - problem.psqsum[start-1]
	}

	avg := average(sum, length); fmt.Println(avg)

	devi := sqsum - (2*avg*sum) + avg*avg*length

	return devi
}

func average(sum int, length int) int {
	result := float64(sum / length) + 0.5  /*rounds */

	return int(result)
}

func quantization(start int, base int) int {
	if start == problem.base {
		return 0
	}

	if base == 0 {
		return math.MaxInt32
	}

	ret := &problem.cache[start][base]
	if *ret != -1 {
		return *ret
	}

	*ret = math.MaxInt32

	for ps := 1; start + ps < problem.base; ps++ {
		target := deviation(start, start+ps-1) + quantization(start + ps, base-1)

		if target < *ret {
			*ret = target
		}
	}

	return *ret
}

