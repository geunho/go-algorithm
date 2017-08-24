package fence

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestLefmax(t *testing.T) {
	partialPanels := []int{7, 1, 5, 4, 3, 6}
	baseHeight := 3
	solution := 4

	result := leftmax(partialPanels, baseHeight)

	if result != solution {
		t.Error(
			"For", partialPanels,
			"exepected", solution,
			"got", result,
		)
	}
}

func TestRightmax(t *testing.T) {
	partialPanels := []int{4, 3, 7, 1, 2}
	baseHeight := 3
	solution := 3

	result := rightmax(partialPanels, baseHeight)

	if result != solution {
		t.Error(
			"For", partialPanels,
			"exepected", solution,
			"got", result,
		)
	}
}

func TestSolveProblem(t *testing.T) {

	problem.panels = []int{7, 1, 5, 9, 6, 7, 3}
	solution := 20

	result := problem.SolveProblem()

	fmt.Println(result)

	if result != solution {
		t.Error(
			"For", problem.panels,
			"exepected", solution,
			"got", result,
		)
	}
}

func TestSolveProblems(t *testing.T) {
	cases := 50
	caseSize := 20000

	for cases > 0 {
		problem.panels = gencase(caseSize)
		result := problem.SolveProblem()

		fmt.Print(result)
		fmt.Print("  ")

		cases--
	}

	fmt.Println()
}

func gencase(size int) []int {
	min := 1
	max := 10000
	testCase := make([]int, size)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < size; i++ {
		testCase[i] = rand.Intn(max-min) + min
	}

	return testCase
}
