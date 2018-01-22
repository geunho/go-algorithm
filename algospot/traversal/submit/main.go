package main

import "fmt"

func main() {

	problem := Traversal{}

	cases := readNumberOfCases()
	for cases > 0 {
		problem.ReadProblem()

		rootNodeLabel, nodeList := rebuildTree(&problem)
		printPostOrder(nodeList[rootNodeLabel])

		cases--
	}
}

func readNumberOfCases() int {
	var cases int
	fmt.Scanf("%d", &cases)

	return cases
}


type Traversal struct {
	N int  // number of tree nodes
	PreOrder []int
	InOrder []int
}

func (p *Traversal)ReadProblem() {
	var N int
	fmt.Scanf("%d", &N)
	p.N = N

	preOrder := make([]int, N)
	for i := 0; i < N; i++ {
		var number int
		fmt.Scanf("%d", &number)
		preOrder[i] = number
	}

	p.PreOrder = preOrder

	inOrder := make([]int, N)
	for i := 0; i < N; i++ {
		var number int
		fmt.Scanf("%d", &number)
		inOrder[i] = number
	}

	p.InOrder = inOrder
}

type BTNode struct {
	Label int
	Parent *BTNode
	Left *BTNode
	Right *BTNode
}

func printPostOrder(rootNode *BTNode) {
	if rootNode.Left != nil {
		printPostOrder(rootNode.Left)
	}

	if rootNode.Right != nil {
		printPostOrder(rootNode.Right)
	}

	fmt.Print(rootNode.Label, " ")
}

func rebuildTree(traversal *Traversal) (int, map[int]*BTNode) {
	p := *traversal

	nodeList := make(map[int]*BTNode)

	rootNode := getTreeRoot(&(p.PreOrder))
	nodeList[rootNode.Label] = &rootNode

	parentNode := &rootNode

	preOrderIdx := 0
	for inOrderIdx:= 0; inOrderIdx < p.N; inOrderIdx++ {
		inOrderNumber := p.InOrder[inOrderIdx]

		if nodeList[inOrderNumber] != nil {
			parentNode = nodeList[inOrderNumber]
			continue
		}

		for {
			preOrderIdx++
			preOrderNumber := p.PreOrder[preOrderIdx]

			nextNode := &BTNode {preOrderNumber, parentNode, nil, nil}
			nodeList[preOrderNumber] = nextNode

			if parentNode.Left == nil {
				parentNode.Left = nextNode
				parentNode = nextNode

			} else {
				parentNode.Right = nextNode
				parentNode = nextNode
			}

			if inOrderNumber == preOrderNumber {
				break
			}
		}
	}

	return p.PreOrder[0], nodeList
}

func getTreeRoot(preOrder *[]int) BTNode {
	return BTNode {(*preOrder)[0], nil, nil, nil}
}