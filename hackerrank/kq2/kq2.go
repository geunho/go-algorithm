package main

import (
	"fmt"
)

import "sort"

func main() {
	var names []string
	var names_size int
	fmt.Scanf("%d", &names_size)

	for i := 0; i < names_size; i++ {
		var names_item string
		fmt.Scanf("%s", &names_item)
		names = append(names, names_item)

	}

	fmt.Println(names)

	var res []string = getSortedList(names)
	for res_i := range res {
		fmt.Println(res[res_i])
	}
}

var ROMANS = map[string]int {
	"I":1, "V":5, "X":10, "L":50,
}

type RoyalName struct {
	name string
	birthOrder string
	order int
}
/*
 * Complete the function below.
 */
func getSortedList(names []string) []string {
	royalNames := make([]RoyalName, len(names))
	for i := 0; i < len(names); i++ {
		royalNames[i] = convertToRoyalName(names[i])
	}

	sort.Slice(royalNames, func(i, j int) bool {
		if royalNames[i].name == royalNames[j].name {
			return royalNames[i].order < royalNames[j].order
		} else {
			return compareName(royalNames[i].name, royalNames[j].name)
		}
	})

	result := make([]string, len(names))

	for i:=0; i<len(royalNames); i++ {
		result[i] = formatStringRoyalName(royalNames[i])
	}

	return result
}

func formatStringRoyalName(royalName RoyalName) string {
	return fmt.Sprintf("%s %s", royalName.name, royalName.birthOrder)
}

// name의 순서가 targetName 보다 빠른 경우 true 반환
func compareName(name string, targetName string) bool {
	nr := []rune(name)
	tn := []rune(targetName)
	loops := 0
	if len(nr) < len(tn) {
		loops = len(nr)
	} else {
		loops = len(tn)
	}

	for i:=0; i<loops; i++ {
		if nr[i] == tn[i] {
			continue
		} else if nr[i] < tn[i] {
			return true
		} else {
			return false
		}
	}

	return true
}

// strings 패키지를 사용할 수 없는 제약이 있음
func convertToRoyalName(royalName string) RoyalName {
	nameRunes := []rune(royalName)
	splitIndex := 0
	for index, rune := range nameRunes {
		// whitespace: 32
		if rune == 32 {
			splitIndex = index
			break
		}
	}

	name := nameRunes[:splitIndex]
	birthOrder := nameRunes[splitIndex+1:]
	order := convertBirthOrder(string(birthOrder))

	return RoyalName{string(name), string(birthOrder), order}
}

func convertBirthOrder(birthOrder string) int {
	// XL -> 10 + (50 - 10*2) = 40
	// IX -> 1 + (10 - 1*2)  = 9
	// 이전 문자보다 값이 큰 경우 이전 문자의 값은 차감해야 한다.

	birthRunes := []rune(birthOrder)

	result := 0
	for i:=0; i<len(birthRunes); i++ {
		result += ROMANS[string(birthRunes[i])]

		if i > 0 && (ROMANS[string(birthRunes[i])] > ROMANS[string(birthRunes[i-1])]) {
			result -= ROMANS[string(birthRunes[i-1])] * 2
		}
	}

	return result
}