package strjoin

import (
	"fmt"
	"sort"
)

/*
Algospot
	Chapter 10.3 STRING JOIN
	(https://algospot.com/judge/problem/read/STRJOIN)
*/

const ProblemTitle = "String Join"

type StringJoin struct {
	lengths []int
}

func (p StringJoin) GetProblemTitle() string {
	return ProblemTitle
}

func (p *StringJoin) ReadProblem() {
	var nNumbers int
	fmt.Scanf("%d", &nNumbers)

	lengths := make([]int, nNumbers)

	for i := 0; i < nNumbers; i++ {
		var length int
		fmt.Scanf("%d", &length)

		lengths[i] = length
	}

	p.lengths = lengths
}

func (p StringJoin) SolveProblem() interface{} {
	return min(p.lengths)
}

func min(lengths []int) int {
	// O(nlogn)
	sort.Ints(lengths)

	return calc(lengths)
}

func calc(lengths []int) int {
	size := len(lengths)

	// 이미 다 계산 되었음
	if size < 2 {
		return 0
	}

	sum := lengths[0] + lengths[1]
	deleteFrontPairs(&lengths)
	insertElementAtSorttedSlice(sum, &lengths)

	return sum + calc(lengths)
}

func insertElementAtSorttedSlice(elem int, slice *[]int) {
	size := len(*slice)
	if size > 0 {
		// O(n)
		for i:=0; i<len(*slice); i++ {
			if elem <= (*slice)[i] {
				*slice = append(*slice, 0)
				copy((*slice)[i+1:], (*slice)[i:])
				(*slice)[i] = elem
				return
			}
		}
	}

	*slice = append(*slice, elem)
}

// index의 원소를 제거하고 제거한 원소를 반환한다.
func deleteFrontPairs(slice *[]int) {
	*slice = (*slice)[2:]
}