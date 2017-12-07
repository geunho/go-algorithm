package fence2

import "fmt"

/*
Algospot
	Chapter 19.3 FENCE (by queue)
	(https://algospot.com/judge/problem/read/FENCE)
*/

const ProblemTitle = "Fence Cutting 2"

type Fence2 struct {
	panels []int
}

func (p Fence2) GetProblemTitle() string {
	return ProblemTitle
}

func (p *Fence2) ReadProblem() {
	var nPanels int
	fmt.Scanf("%d", &nPanels)

	panels := make([]int, nPanels)
	size := nPanels
	for i := 0; i < size; i++ {
		var h int
		fmt.Scanf("%d", &h)

		panels[i] = h
	}

	p.panels = panels
}

func (p Fence2) SolveProblem() interface{} {

	return 0
}