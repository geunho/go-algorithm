package picnic

import (
	"fmt"
)

/*
Algospot
	Chapter 6.3 PICNIC
	(https://algospot.com/judge/problem/read/PICNIC)
*/

const ProblemTitle = "Picnic Problem"


var Cases int
/*
{
  0: { 1:true, 3:true },
  1: { 0:true, 2:true },
  ...
}
*/
// 짝을 이룰 수 있는 친구 매트릭스
var aMatchMatrix map[int]([]int)

var nStudents int
var nFriendsPairs int
var nSolutions = 0

func ReadNumberOfCases() {
	fmt.Scanf("%d", &Cases)
}

func ReadProblem() {
	fmt.Scan(&nStudents)
	fmt.Scan(&nFriendsPairs)

	aMatchMatrix = make(map[int]([]int))

	for i := 0; i < nStudents; i++ {
		aMatchMatrix[i] = []int{}
	}

	for i := 0; i < (nFriendsPairs*2)-1; i += 2 {
		var first int
		var second int

		fmt.Scan(&first)
		fmt.Scan(&second)

		if first < second {
			aMatchMatrix[first] = append(aMatchMatrix[first], second)
		} else {
			aMatchMatrix[second] = append(aMatchMatrix[second], first)
		}
	}
	fmt.Println(aMatchMatrix)
}

func SolveProblem() {
	// 기저: 학생 수가 홀 수일 때
	if nStudents % 2 != 0 {
		PrintSolution(0)
	}

	toSelect := len(aMatchMatrix[0])
	selected := make(map[int]bool)
	selectPairs(toSelect, &selected)


}

func PrintSolution(solution int) {
	fmt.Println(solution)
}

func selectPairs(toSelect int, selected *map[int]bool) {
	if toSelect == 0 {
		PrintSolution(nSolutions)
	}
	for i := 0; i < nStudents; i++ {
		(*selected)[i] = true

		for len(aMatchMatrix[i]) > 0 {

		}
		//selectPairs(nPairs - 1, selected)
	}
}