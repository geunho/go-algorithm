package traversal

import "fmt"

/*
Algospot
	Chapter 21.3 TRAVERSAL
	(https://algospot.com/judge/problem/read/TRAVERSAL)

Traverses
	PreOrder
	루트 노드에서 시작해서,
	1. 노드를 방문한다.
	2. 왼쪽 서브 트리를 전위 순회한다.
	3. 오른쪽 서브 트리를 전위 순회한다.

	InOrder
	1. 왼쪽 서브 트리를 중위 순회한다.
	2. 노드를 방문한다.
	3. 오른쪽 서브 트리를 중위 순회한다.

	PostOrder
	1. 왼쪽 서브 트리를 후위 순회한다.
	2. 오른쪽 서브 트리를 후위 순회한다.
	3. 노드를 방문한다.
*/

const ProblemTitle = "Change Order of Traversal Tree"

type Traversal struct {
	N int  // number of tree nodes
	PreOrder []int
	InOrder []int
}

func (p Traversal) GetProblemTitle() string {
	return ProblemTitle
}

func (p *Traversal) ReadProblem() {
	var N int
	fmt.Scanf("%d", &N)

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

func (p Traversal) SolveProblem() interface{} {
	return 0
 }

 /* Binary Tree Node */
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

func rebuildTree(traversal *Traversal) map[int]*BTNode {
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

	return nodeList
 }

func getTreeRoot(preOrder *[]int) BTNode {
	return BTNode {(*preOrder)[0], nil, nil, nil}
}