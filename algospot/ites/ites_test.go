package ites

import (
	"testing"
	"fmt"
)

func TestITES_generateNext(t *testing.T) {
	var result uint = 1983

	for i:=0; i < 10; i++ {
		result = generateNext(result)
		fmt.Println(result)
	}
}

func TestITES_generateSignal(t *testing.T) {
	// 1983, 426918790, 3629214769, 1552047664, 2788443187
	// 1984, 8791, 4770, 7665, 3188

	A := []uint {
		1983, 426918790, 3629214769, 1552047664, 2788443187,
	}

	signals := []uint { 1984, 8791, 4770, 7665, 3188 }

	for i:=0; i<len(A); i++ {
		calc := generateSignal(A[i])
		fmt.Println(calc)
		if calc != signals[i] {
			t.Error("Wrong result: ", calc)
		}
	}
}

func TestITES_analyze_TestCase1(t *testing.T) {
	sol := 1
	calc := analyze(8791, 20)

	if calc != sol {
		t.Error("Wrong calculation.", calc)
	}
}

func TestITES_analyze_TestCase2(t *testing.T) {
	sol := 4
	calc := analyze(5265, 5000)

	if calc != sol {
		t.Error("Wrong calculation.", calc)
	}
}

func TestITES_analyze_TestCase3(t *testing.T) {
	sol := 1049
	calc := analyze(3578452, 5000000)

	if calc != sol {
		t.Error("Wrong calculation.", calc)
	}
}