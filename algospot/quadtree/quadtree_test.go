package quadtree

import (
	"testing"
	"fmt"
)

func TestSolveProblem(t *testing.T) {
	problem.compressed = "xxwwwbxwxwbbbwwxxxwwbbbwwwwbb"
	//problem.compressed = "xbwwb"
	result := problem.SolveProblem()

	fmt.Println(result)
}
