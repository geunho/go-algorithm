package treap

import (
	"testing"
)

var treap *Treap

func compare(root *Node, t *testing.T) {
	if root == nil {
		return
	}

	if root.Left != nil &&
		root.Priority < root.Left.Priority {
		t.Error("Root priority should bigger than children's.")
	}
	if root.Right != nil &&
		root.Priority < root.Right.Priority {
		t.Error("Root priority should bigger than children's.")
	}

	if root.Left != nil &&
		root.Key < root.Left.Key{
		t.Error("Root Key should bigger than left node's.")
	}
	if root.Right != nil &&
		root.Key > root.Right.Key {
		t.Error("Root Key should smaller than right node's.")
	}

	compare(root.Left, t)
	compare(root.Right, t)
}

func TestTreap_Insert(t *testing.T) {
	if treap == nil {
		treap = &Treap{nil, make(map[KeyType]*Node) }
	}

	for i := 1; i < 8; i++ {
		treap.Insert(NewNode(i))
		root := treap.Root
		compare(root, t)
	}
}

func TestTreap_Delete(t *testing.T) {
	treap.Delete(7)
	root := treap.Root
	compare(root, t)

	treap.Delete(4)
	root = treap.Root
	compare(root, t)

	treap.Delete(3)
	root = treap.Root
	compare(root, t)

	treap.Delete(2)
	root = treap.Root
	compare(root, t)

	treap.Delete(5)
	root = treap.Root
	compare(root, t)
}