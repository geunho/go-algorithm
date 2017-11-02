package main

import (
	"testing"
	"fmt"
)

func TestCountingCase1(t *testing.T) {
	result := counting("00110")
	var solution uint32 = 3

	fmt.Println(result)

	if result != solution {
		t.Error(
			"expected", solution,
			"got", result,
		)
	}
}

func TestCountingCase2(t *testing.T) {
	result := counting("10001")
	var solution uint32 = 2

	fmt.Println(result)

	if result != solution {
		t.Error(
			"expected", solution,
			"got", result,
		)
	}
}

func TestCountingCase3(t *testing.T) {
	result := counting("10101")
	var solution uint32 = 4

	fmt.Println(result)

	if result != solution {
		t.Error(
			"expected", solution,
			"got", result,
		)
	}
}
