package fence

import (
	"testing"
	"fmt"
	"reflect"
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

func TestSolveProblem(t *testing.T) {
	problem.panels = []int{ 7, 1, 5, 9, 6, 7, 3}
	solution := 20

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
