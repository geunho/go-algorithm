package fence

import (
	"testing"
	"fmt"
	"reflect"
	"time"
	"math/rand"
)

func TestTransformInputs(t *testing.T) {
	problem.nPanels = 7

	input := "7 1 5 9 6 7 3"
	transformed := transformInput(input)

	panels := []int{ 7, 1, 5, 9, 6, 7, 3}

	if !reflect.DeepEqual(transformed, panels) {
		t.Error(
			"For", input,
			"exepected", panels,
			"got", transformed,
		)
	}

}

func TestLefmax(t *testing.T) {
	partialPanels := []int{ 7,1,5, 4, 3, 6}
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
	partialPanels := []int{ 4, 3, 7, 1, 2}
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

	inputStr := "1 4 4 4 4 1 1"
	problem.nPanels = 7
	problem.panels = transformInput(inputStr)
	solution := 16

	result := problem.SolveProblem()

	fmt.Println(result)

	if result != solution  {
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
		inputStr := genTestCase(caseSize)
		problem.nPanels = caseSize
		problem.panels = transformInput(inputStr)

		result := problem.SolveProblem()

		fmt.Print(result)
		fmt.Print("  ")

		cases--
	}
}

func genTestCase(size int) string {
	min := 1
	max := 10000
	testCase := ""

	rand.Seed(time.Now().Unix())

	for size > 0 {
		testCase += string(rand.Intn(max - min) + min) + " "
		size--
	}

	return testCase
}
