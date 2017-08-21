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

const ProblemTItle = "Fan Meeting Problem"

const (
	M = true
	F = false
)

func ReadNumberOfCases() int {
	var cases int
	fmt.Scanf("%d", &cases)

	return cases
}

func ReadProblem() ([]bool, []bool){
	var membersStr string
	fmt.Scan(&membersStr)
	members := transformGenderInputs(membersStr)

	var fansStr string
	fmt.Scan(&fansStr)
	fans := transformGenderInputs(fansStr)

	return members, fans
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
	nMembers := len(members)
	nIterations := len(fans) - nMembers + 1

	menInMembers := findMenIndicesInMembers(&members)

	if len(menInMembers) == 0 {
		return nIterations
	}

	menInFans := findMenIndicesInMembers(&fans)

	if len(menInFans) == 0 {
		return nIterations
	}

	result := asyncHug(&menInMembers, &fans, nIterations, nMembers)
	//result := hug(&menInMembers, &fans, nIterations, nMembers)

	menInMembers = nil
	menInFans = nil

	return result
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

	iterGroup := sync.WaitGroup{}

	var lock sync.Mutex

	for i := 0; i < nIterations; i++ {
		targetFans := (*fans)[i : i+nMembers]
		if  i%4 == 0 {
			if isAllMemberHug(menInMembers, &targetFans) {
				lock.Lock()
				nAllMembersHug++
				lock.Unlock()
			}

			targetFans = nil
		} else {
			iterGroup.Add(1)
			go func() {
				defer iterGroup.Done()
				if isAllMemberHug(menInMembers, &targetFans) {
					lock.Lock()
					defer lock.Unlock()
					nAllMembersHug++
				}
				targetFans = nil
			}()
		}
	}

	iterGroup.Wait()

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
