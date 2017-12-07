package brackets2

import "testing"

func TestBrackets2_arePair(t *testing.T) {
	openBrackets := []rune("([{")
	closeBrackets := []rune(")]}")

	for i:=0; i < len(openBrackets); i++ {
		if !arePair(openBrackets[i], closeBrackets[i]) {
			t.Error("짝이 맞지 않습니다.", i)
		}
	}

}

func TestBrackets2_validate_TestCase1(t *testing.T) {
	input := []rune("()()")
	solution := true

	result := validate(&input)

	if solution != result {
		t.Error("계산에 오류가 있습니다.", result)
	}
}

func TestBrackets2_validate_TestCase2(t *testing.T) {
	input := []rune("({[}])")
	solution := false

	result := validate(&input)

	if solution != result {
		t.Error("계산에 오류가 있습니다.", result)
	}
}

func TestBrackets2_validate_TestCase3(t *testing.T) {
	input := []rune("({}[(){}])")
	solution := true

	result := validate(&input)

	if solution != result {
		t.Error("계산에 오류가 있습니다.", result)
	}
}
