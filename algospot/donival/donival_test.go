package donival

import (
	"testing"
	"fmt"
)

func TestSearch(t *testing.T) {
	r0 := search(0, 2)
	r2 := search(2, 2)
	r4 := search(4, 2)

	fmt.Println(r0)
	fmt.Println(r2)
	fmt.Println(r4)
}