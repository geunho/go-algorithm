package quadtree

import "fmt"

/*
Algospot
	Chapter 7.2 QUADTREE
	(https://algospot.com/judge/problem/read/QUADTREE)
 */

const ProblemTitle = "Revsere Quad Tree"

func ReadNumberOfCases() int {
	var cases int
	fmt.Scanf("%d", &cases)

	return cases
}

func ReadProble() string {
	var compressed string
	fmt.Scanf("%s", compressed)

	return compressed
}

func SolveProblem(compressed string) string {
	result, _ := reverse([]rune(compressed))

	return result
}

func reverse(compressed []rune) (string, int) {
	fmt.Println(compressed)
	tree := []string{}

	iterations := 0

	for iterations < len(compressed) {

		if compressed[iterations] == 'x' {

			partialResult, delta := reverse(compressed[iterations+1:])

			tree = append(tree, partialResult)
			iterations += delta
			if iterations >= len(compressed) {
				break;
			}
		} else {

			tree = append(tree, string(compressed[iterations]))
			iterations++
			fmt.Println(tree)
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