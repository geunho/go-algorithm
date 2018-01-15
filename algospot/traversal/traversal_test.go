package traversal

import (
	"testing"
	"os"
	"bytes"
	"io"
	"strings"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestTraversal_rebuildTree_And_printPostOrder_TestCase1(t *testing.T) {
	problem := Traversal{
		7,
		[]int{27, 16, 9, 12, 54, 36, 72},
		[]int{9, 12, 16, 27, 36, 54, 72},
		}

	tree := rebuildTree(&problem)

	result := captureStdout(func() {
		printPostOrder(tree[27])
	})

	if strings.EqualFold(result, "12 9 16 36 72 54 27") {
		t.Error("fail")
	}
}

func TestTraversal_rebuildTree_And_printPostOrder_TestCase2(t *testing.T) {
	problem := Traversal{
		6,
		[]int{409, 479, 10, 838, 150, 441},
		[]int{409, 10, 479, 150, 838, 441},
	}

	tree := rebuildTree(&problem)

	result := captureStdout(func() {
		printPostOrder(tree[409])
	})

	if strings.EqualFold(result, "10 150 441 838 479 409") {
		t.Error("fail")
	}
}