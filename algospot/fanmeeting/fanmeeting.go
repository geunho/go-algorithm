package fanmeeting

import (
	"fmt"
)

/*
Algospot
	Chapter 7.6 FANMEETING
	(https://algospot.com/judge/problem/read/FANMEETING)
*/

const ProblemTItle = "Fan Meeting Problem"

const M = true
const F = false

var Cases int

var _members []bool
var _fans []bool

func ReadNumberOfCases() {
	fmt.Scanf("%d", &Cases)
}

func ReadProblem() {
	var membersStr string
	fmt.Scan(&membersStr)
	_members = transformGenderInputs(membersStr)

	var fansStr string
	fmt.Scan(&fansStr)
	_fans = transformGenderInputs(fansStr)
}

func transformGenderInputs(gendersStr string) []bool {
	gendersRune := []rune(gendersStr)
	genders := []bool{}
	for _, r := range gendersRune {
		var gender bool

		switch r {
		case 'M':
			gender = M
		case 'F':
			gender = F
		}

		genders = append(genders, gender)
	}

	return genders
}

func SolveProblem(members []bool, fans []bool) int {
	nAllMembersHug := 0

	nMembers := len(members)
	nIterations := len(fans) - nMembers + 1

	menInMembers := findMenIndicesInMembers(&members)

	if len(menInMembers) == 0 {
		return nIterations
	}

	for i := 0; i < nIterations; i++ {
		targetFans := fans[i : i+nMembers]

		if isAllMemberHug(&menInMembers, &targetFans) {
			nAllMembersHug++
		}
	}

	return nAllMembersHug
}

func findMenIndicesInMembers(members *[]bool) []int {
	menInMembers := []int{}

	for index, member := range *members {
		if member == M {
			menInMembers = append(menInMembers, index)
		}
	}

	return menInMembers
}

func isAllMemberHug(menInMembers *[]int, fans *[]bool) bool {
	for _, menIndex := range *menInMembers {
		if (*fans)[menIndex] == M {
			return false
		}
	}

	return true
}
