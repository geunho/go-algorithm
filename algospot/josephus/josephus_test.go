package josephus

import (
	"testing"
	"fmt"
)

func TestJosephus_solve_TestCase1(t *testing.T) {
	n, k := 6, 3
	s1, s2 := 3, 5

	r1, r2 := solve(n, k)

	fmt.Println(r1, r2)

	if s1 != r1 {
		t.Error("계산에 오류가 있습니다.", r1)
	}

	if s2 != r2 {
		t.Error("계산에 오류가 있습니다.", r2)
	}

}

func TestJosephus_solve_TestCase2(t *testing.T) {
	n, k := 40, 3
	s1, s2 := 11, 26

	r1, r2 := solve(n, k)

	fmt.Println(r1, r2)

	if s1 != r1 {
		t.Error("계산에 오류가 있습니다.", r1)
	}

	if s2 != r2 {
		t.Error("계산에 오류가 있습니다.", r2)
	}

}