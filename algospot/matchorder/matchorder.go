package matchorder

import (
	"fmt"
	"sort"
)

/*
Algospot
	Chapter 10.1 MATCH ORDER
	(https://algospot.com/judge/problem/read/MATCHORDER)
*/

const ProblemTitle = "Match Order"

type MatchOrder struct {
	alias []int
	enemies []int
}

func (p MatchOrder) GetProblemTitle() string {
	return ProblemTitle
}

func (p *MatchOrder) ReadProblem() {
	var nPlayers int
	fmt.Scanf("%d", &nPlayers)

	enemies := make([]int, nPlayers)

	for i := 0; i < nPlayers; i++ {
		var rating int
		fmt.Scanf("%d", &rating)

		enemies[i] = rating
	}

	alias := make([]int, nPlayers)

	for i := 0; i < nPlayers; i++ {
		var rating int
		fmt.Scanf("%d", &rating)

		alias[i] = rating
	}

	p.alias = alias
	p.enemies = enemies
}

func (p MatchOrder) SolveProblem() interface{} {
	return maxwin(p.alias, p.enemies)
}

func maxwin(alias []int, enemies []int) int {
	sort.Ints(alias)
	sort.Ints(enemies)

	nWins := 0

	for i:= 0; i <len(enemies); i++ {
		selectedIdx := 0

		for j:= 0; j <len(alias); j++ {
			if enemies[i] <= alias[j] {
				nWins++
				selectedIdx = j
				break
			}
		}

		alias = append(alias[:selectedIdx], alias[selectedIdx+1:]...)
	}

	return nWins
}