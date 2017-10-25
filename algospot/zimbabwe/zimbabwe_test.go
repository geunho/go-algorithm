package zimbabwe

import (
	"testing"
	"fmt"
)

func TestPredict(t *testing.T) {
	predict("", []bool{false,false,false})

	fmt.Println(counter)
}