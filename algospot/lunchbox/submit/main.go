package main

import (
	"sort"
	"fmt"
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

type TimeTakenSorter []TimeTaken
func (a TimeTakenSorter) Len() int           { return len(a) }
func (a TimeTakenSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TimeTakenSorter) Less(i, j int) bool { return a[i].timeToEat > a[j].timeToEat }


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

func short(cook, eat []int) int {
	timeTakens := make([]TimeTaken, len(cook))

	for i:=0; i <len(cook); i++ {
		timeTakens[i] = TimeTaken{ cook[i], eat[i] }
	}


	// >=go 1.8
	//sort.Slice(timeTakens, func(i, j int) bool {
	//	return timeTakens[i].timeToEat > timeTakens[j].timeToEat
	//})
	sort.Sort(TimeTakenSorter(timeTakens))

	totalTimeToCook := 0
	for i:=0; i< len(timeTakens); i++ {
		totalTimeToCook += timeTakens[i].timeToCook
	}

	cookWindows := totalTimeToCook

	for i:=0; i< len(timeTakens); i++ {
		cookWindows -= timeTakens[i].timeToCook

		if timeTakens[i].timeToEat > cookWindows {
			needMoreTimeToEat := timeTakens[i].timeToEat - cookWindows
			cookWindows += needMoreTimeToEat
		}
	}

	totalTimeTaken := totalTimeToCook + cookWindows

	return totalTimeTaken
}

func main() {
	cases := readNumberOfCases()
	for cases > 0 {
		problem := &LunchBox{}
		problem.ReadProblem()
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