package quantization2

import (
	"testing"
	"fmt"
)

func TestDeviation(t *testing.T) {
	problem.numbers = []int{ 10, 4, 1, 5, 7 }
	initialize()

	result := deviation(3, 4)

	fmt.Println(result)
}

func TestQuantization(t *testing.T) {
	problem.numbers = []int{ 1, 744, 755, 4, 897, 902, 890, 6, 777 }
	problem.base = 3
	initialize()

	ret := quantization(0, problem.base)

	fmt.Println(ret)

}