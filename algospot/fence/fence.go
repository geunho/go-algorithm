package fence

import (
	"fmt"
)

/*
Algospot
	Chapter 7.4 FENCE
	(https://algospot.com/judge/problem/read/FENCE)
*/

const ProblemTitle = "Fence Cutting"

type Fence struct {
	panels []int
}

var problem = Fence{}

func (p Fence) GetProblemTitle() string {
	return ProblemTitle
}

func (p Fence) ReadProblem() {
	var nPanels int
	fmt.Scanf("%d", &nPanels)

	panels := make([]int, nPanels)
	size := nPanels
	for i := 0; i < size; i++ {
		var h int
		fmt.Scanf("%d", &h)

		panels[i] = h
	}

	problem.panels = panels
}

func (p Fence) SolveProblem() interface{} {
	return fence(problem.panels)
}

// brute force
func fence(panels []int) int {
	maxArea := 0

	for i := 0; i < len(panels); i++ {
		panel := panels[i]

		left := leftmax(panels[:i], panel)
		right := rightmax(panels[i+1:], panel)
		area := panel * (1 + left + right)

		if maxArea < area {
			maxArea = area
		}
	}

	return maxArea
}

func leftmax(partialPanels []int, baseHeight int) int {
	sum := 0

	startIndex := len(partialPanels) - 1

	for i := startIndex; i >= 0; i-- {
		panel := partialPanels[i]

		if baseHeight <= panel {
			sum++
		} else {
			break
		}
	}

	return sum
}

func rightmax(partialPanels []int, baseHeight int) int {
	sum := 0

	for _, panel := range partialPanels {
		if baseHeight <= panel {
			sum++
		} else {
			break
		}
	}

	return sum
}
