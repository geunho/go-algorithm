package treap

import "math/rand"

type KeyType int

type Treap struct {
	Root *Node
	Nodes map[KeyType]*Node
}

type Node struct {
	Priority, Size int
	Key KeyType
	Left, Right *Node
}

func NewNode(key int) *Node {
	return &Node{rand.Int(), 1, KeyType(key),nil, nil }
}

func (n *Node)calcSize() {
	n.Size = 1
	if n.Left != nil {
		n.Size += n.Left.Size
	}
	if n.Right != nil {
		n.Size += n.Right.Size
	}
}

func (n *Node)setLeft(left *Node) {
	n.Left = left
	n.calcSize()
}

func (n *Node)setRight(right *Node) {
	n.Right = right
	n.calcSize()
}

func split(root *Node, key KeyType) (*Node, *Node) {
	if root == nil {
		return nil, nil
	}

	if root.Key < key {
		left, right := split(root.Right, key)

		root.setRight(left)
		return root, right
	} else {
		left, right := split(root.Left, key)

		root.setLeft(right)
		return left, root
	}
}

func (tr *Treap)Insert(node *Node) {
	root := tr.Root

	tr.Nodes[node.Key] = node
	tr.Root = insert(root, node)
}

func insert(root *Node, node *Node) *Node {
	if root == nil {
		return node
	}

	if root.Priority < node.Priority {
		left, right := split(root, node.Key)

		node.Left = left
		node.Right = right

		root = node
		return root

	} else if root.Key < node.Key {
		root.setRight(insert(root.Right, node))
	} else {
		root.setLeft(insert(root.Left, node))
	}

	return root
}

func merge(left *Node, right *Node) *Node {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.Priority < right.Priority {
		right.setLeft(merge(left, right.Left))
		return right

	} else {
		left.setRight(merge(left.Right, right))
		return left
	}
}

func (tr *Treap)Delete(key KeyType) {
	root := tr.Root

	tr.Root = del(root, key)
	delete(tr.Nodes, key)
}

func del(root *Node, key KeyType) *Node {
	if root == nil {
		return root
	}

	if root.Key == key {
		newRoot := merge(root.Left, root.Right)
		return newRoot
	}

	if root.Key < key  {
		root.setRight(del(root.Right, key))
	} else {
		root.setLeft(del(root.Left, key))
	}
	return root
}


