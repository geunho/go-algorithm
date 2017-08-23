package fence

import (
	"fmt"
	"strings"
	"strconv"
)

/*
Algospot
	Chapter 7.4 FENCE
	(https://algospot.com/judge/problem/read/FENCE)
*/

const ProblemTitle = "Fence Cutting"

type Fence struct {
	nPanels int
	panels []int
}

var problem = Fence{}

func (p Fence) GetProblemTitle() string {
	return ProblemTitle
}

func (p Fence) ReadProblem() {
	fmt.Scanf("%d", &(problem.nPanels))

	var panelsStr string
	fmt.Scan(&panelsStr)

	problem.panels = transformInput(panelsStr)
}

func transformInput(panelsStr string) []int {
	panels := make([]int, problem.nPanels)

	for i, s := range strings.Fields(panelsStr) {
		parsed, _ := strconv.ParseInt(s, 10, 64)
		panels[i] = int(parsed)
	}

	return panels
}

func (p Fence) SolveProblem() interface{} {
	maxArea := 0

	for i, panel := range problem.panels {

		left := leftmax(problem.panels[:i], panel)
		right := rightmax(problem.panels[i+1:], panel)
		area := panel * (1 + left + right)

		if maxArea < area {
			maxArea = area
		}

	}

	return maxArea
}

func leftmax(partialPanels []int, baseHeight int) int {
	sum := 0

	if len(partialPanels) == 0 {
		return sum
	}

	for i := len(partialPanels) -1; i < 0; i-- {
		panel := partialPanels[i]

		if baseHeight <= panel {
			sum++
		} else {
			break;
		}
	}

	return sum
}

func rightmax(partialPanels []int, baseHeight int) int {
	sum := 0
	for panel := range partialPanels {
		if baseHeight <= panel  {
			sum++
		} else {
			break;
		}
	}

	return sum
}


