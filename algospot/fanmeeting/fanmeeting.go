package fanmeeting

import (
	"fmt"
	"sync"
)

/*
Algospot
	Chapter 7.6 FANMEETING
	(https://algospot.com/judge/problem/read/FANMEETING)
*/

const ProblemTitle = "Fan Meeting Problem"

type FanMeeting struct {
	members []bool
	fans    []bool
}

const (
	M = true
	F = false
)

func (p FanMeeting) GetProblemTitle() string {
	return ProblemTitle
}

func (p *FanMeeting) ReadProblem() {
	var membersStr string
	fmt.Scan(&membersStr)
	p.members = transformGenderInputs(membersStr)

	var fansStr string
	fmt.Scan(&fansStr)
	p.fans = transformGenderInputs(fansStr)
}

func (p FanMeeting) SolveProblem() interface{} {
	nMembers := len(p.members)
	nIterations := len(p.fans) - nMembers + 1

	menInMembers := findMenIndicesInMembers(&(p.members))

	if len(menInMembers) == 0 {
		return nIterations
	}

	menInFans := findMenIndicesInMembers(&(p.fans))

	if len(menInFans) == 0 {
		return nIterations
	}

	//result := asyncHug(&menInMembers, &(p.fans), nIterations, nMembers)
	result := hug(&menInMembers, &(p.fans), nIterations, nMembers)

	return result
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

func hug(menInMembers *[]int, fans *[]bool, nIterations int, nMembers int) int {
	nAllMembersHug := 0

	for i := 0; i < nIterations; i++ {
		targetFans := (*fans)[i : i+nMembers]

		if isAllMemberHug(menInMembers, &targetFans) {
			nAllMembersHug++
		}
	}

	return nAllMembersHug
}

func asyncHug(menInMembers *[]int, fans *[]bool, nIterations int, nMembers int) int {
	nAllMembersHug := 0

	var lock sync.Mutex
	for i := 0; i < nIterations; i++ {
		targetFans := (*fans)[i : i+nMembers]

		go func() {
			if isAllMemberHug(menInMembers, &targetFans) {
				lock.Lock()
				defer lock.Unlock()
				nAllMembersHug++
			}
			targetFans = nil
		}()
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
