package main

import "fmt"

/*
Algospot
	Chapter 18.5 Josephus Problem
	(https://algospot.com/judge/problem/read/JOSEPHUS)
*/

type Josephus struct {
	N int /* Number of person  */
	K int /* Offset */
}

func ReadProblem() Josephus {
	var numbers int
	fmt.Scanf("%d", &numbers)

	var offset int
	fmt.Scanf("%d", &offset)

	return Josephus{numbers, offset }
}

func (p Josephus) SolveProblem() interface{} {
	first, second := solve(p.N, p.K)

	result := fmt.Sprintf("%d %d", first, second)

	return result
}

func solve(n int, k int) (int, int) {
	// init
	menAlive := make([]int, n)

	for i := 0; i < n; i++ {
		menAlive[i] = i + 1
	}

	targetIndex := 0

	for len(menAlive) > 2 {
		menAlive = append(menAlive[:targetIndex], menAlive[targetIndex+1:]...)

		targetIndex = nextIndex(targetIndex + k, len(menAlive))
	}

	return menAlive[0], menAlive[1]
}

func nextIndex(index int, length int) int {
	return (index-1) % length
}


func main() {
	cases := readNumberOfCases()
	for cases > 0 {
		problem := ReadProblem()
		result := problem.SolveProblem()

		fmt.Println(result)

		cases--
	}
}

func readNumberOfCases() int {
	var cases int
	fmt.Scanf("%d", &cases)

	return cases
}