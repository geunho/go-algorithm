package main

import (
	"fmt"
)


/*
 * Complete the function below.
 */
func counting(s string) uint32 {
	var count uint32

	runes := []rune(s)
	size := len(runes)
	length := 0

	// O(n*(m + m)) = O(nm)
	for i :=0; i< size-1; i++ {
		length = 0 // init

		for j := i; j<size; j++ {
			if runes[i] == runes[j] {
				length++
			} else {
				break
			}
		}

		nextBitsStartIndex := i + length
		nextBitsEndIndex := nextBitsStartIndex + length-1
		if nextBitsEndIndex > size-1 {
			continue // 다음 비트가 길이 초과
		} else {
			for j:=nextBitsStartIndex; j<nextBitsEndIndex+1; j++ {
				if runes[nextBitsStartIndex] != runes[j] {
					break// length 길이만큼의 비트가 연속되지 못함
				} else if j == nextBitsEndIndex {
					count++ // 연속된 비트가 존재
				}
			}
		}
	}

	return count
}

func main() {
	var s string
	fmt.Scanf("%s", &s)

	var res uint32 = counting(s)
	fmt.Println(res)
}
