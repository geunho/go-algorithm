package brackets2


import "fmt"

/*
Algospot
	Chapter 19.4 BRACKETS2
	(https://algospot.com/judge/problem/read/BRACKETS2)
*/

const ProblemTitle = "Brackets 2"

type Brackets2 struct {
	input []rune // (: 40, ):41, {: 123, }:125, [:91, ]:93
}

func (p Brackets2) GetProblemTitle() string {
	return ProblemTitle
}

func (p *Brackets2) ReadProblem() {
	var input string
	fmt.Scanln(&input)

	p.input = []rune(input)
}

func (p Brackets2) SolveProblem() interface{} {

	result := validate(&p.input)

	if result {
		return "YES"
	} else {
		return "NO"
	}
}

func validate(input *[]rune) bool {
	brackets := *input
	stack := make([]rune, 0, len(brackets))

	for _, bracket := range brackets {
		if bracket == 40 || bracket == 123 || bracket == 91 {
			stack = append(stack, bracket)

		} else {
			// 닫는 괄호가 먼저 나오면 짝이 맞지 않는 것
			if len(stack) < 1 {
				return false
			}

			lastIdx := len(stack)-1
			topElement := stack[lastIdx]

			if arePair(topElement, bracket) {
				// 마지막 원소를 삭제
				stack = stack[:lastIdx]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func arePair(open rune, close rune) bool {
	switch open {
	case 40: // (
		return close == 41 // )
	case 91: // [
		return close == 93 // ]
	case 123: // {
		return close == 125 // }
	default:
		return false
	}
}