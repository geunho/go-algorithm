package main

import (
	"testing"
	"fmt"
)

func TestConvertBirthOrder(t *testing.T) {
	order := convertBirthOrder("XLVII")

	fmt.Println(order)
}

func TestCompareName(t *testing.T) {
	result := compareName("Abcde", "Accce")

	fmt.Println(result)
}

func TestGetSortedList(t *testing.T) {
	result := getSortedList([]string{"2", "Louis IX","Philippe II"})

	fmt.Println(result)
}