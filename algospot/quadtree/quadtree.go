package quadtree

import "fmt"

/*
Algospot
	Chapter 7.2 QUADTREE
	(https://algospot.com/judge/problem/read/QUADTREE)
 */

const ProblemTitle = "Revsere Quad Tree"

type QuadTree struct {
	compressed string
}

var problem = QuadTree{}

func (p QuadTree)GetProblemTitle() string {
	return ProblemTitle
}

func (p QuadTree)ReadProblem() {
	fmt.Scanf("%s", &(problem.compressed))
}

func (p QuadTree)SolveProblem() interface{} {
	result, _ := reverse([]rune(problem.compressed))

	return result
}

func reverse(compressed []rune) (string, int) {
	tree := []string{}

	iterations := 0

	for iterations < len(compressed) {

		if compressed[iterations] == 'x' {

			partialResult, delta := reverse(compressed[iterations+1:])

			tree = append(tree, partialResult)
			iterations += delta
			if iterations >= len(compressed) {
				break
			}
		} else {

			tree = append(tree, string(compressed[iterations]))
			iterations++
		}

		if len(tree) == 4 {
			return "x" + string(tree[2]) + string(tree[3]) + string(tree[0]) + string(tree[1]), iterations + 1
		}
	}

	result := ""
	for _, t := range tree {
		result += t
	}

	return result, 0
}