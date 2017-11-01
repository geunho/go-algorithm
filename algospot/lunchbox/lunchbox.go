package lunchbox

import (
	"fmt"
	"sort"
)


/*
Algospot
	Chapter 10.2 LUNCH BOX
	(https://algospot.com/judge/problem/read/LUNCHBOX)
*/

const ProblemTitle = "Lunch Box"

type LunchBox struct {
	cook []int
	eat []int
}

type TimeTaken struct {
	timeToCook int
	timeToEat int
}

func (p LunchBox) GetProblemTitle() string {
	return ProblemTitle
}

func (p *LunchBox) ReadProblem() {
	var nLunchBox int
	fmt.Scanf("%d", &nLunchBox)

	cook := make([]int, nLunchBox)

	for i := 0; i < nLunchBox; i++ {
		var timeTaken int
		fmt.Scanf("%d", &timeTaken)

		cook[i] = timeTaken
	}

	eat := make([]int, nLunchBox)

	for i := 0; i < nLunchBox; i++ {
		var timeTaken int
		fmt.Scanf("%d", &timeTaken)

		eat[i] = timeTaken
	}

	p.cook = cook
	p.eat = eat
}

func (p LunchBox) SolveProblem() interface{} {
	return short(p.cook, p.eat)
}

func short(cook []int, eat []int) int {
	timeTakens := make([]TimeTaken, len(cook))

	for i:=0; i <len(cook); i++ {
		timeTakens[i] = TimeTaken{ cook[i], eat[i] }
	}

	sort.Slice(timeTakens, func(i, j int) bool {
		return timeTakens[i].timeToEat > timeTakens[j].timeToEat
		//if timeTakens[i].timeToCook == timeTakens[j].timeToCook {
		//	return timeTakens[i].timeToEat < timeTakens[j].timeToEat
		//} else {
		//	return timeTakens[i].timeToCook < timeTakens[j].timeToCook
		//}
	})

	totalTimeTaken := 0

	for i:=0; i< len(timeTakens); i++ {
		totalTimeTaken += timeTakens[i].timeToCook
		if i >0 && timeTakens[i-1].timeToEat > timeTakens[i].timeToCook {
			totalTimeTaken += timeTakens[i-1].timeToEat - timeTakens[i].timeToCook
		}
	}

	totalTimeTaken += timeTakens[len(timeTakens)-1].timeToEat

	return totalTimeTaken
}