package picnic

import (
	"fmt"
	"strings"
)

/*
Algospot
	Chapter 6.2 BOGGLE
	(https://algospot.com/judge/problem/read/BOGGLE)
*/

var ProblemTitle = "Picnic Problem"

var cases int

var nStudents int
var nfriendsPairs int

type FriendsPair struct {
	a, b int
}

func readCases() {
	fmt.Scanf("%d", &cases)
}

func readProblems() {
	var numberStr string
	fmt.Scan(&numberStr)

	var numberArr = strings.Split(numberStr, " ")

	nStudents = numberArr[0]
	nfriendsPairs = numberArr[1]

}
