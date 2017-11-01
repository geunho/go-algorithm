package blockgame

import (
	"fmt"
	"testing"
)

const T = true
const F = false

func TestCanWin(t *testing.T) {
	/*
		TEST CASE 1)
		.....
		.##..
		##..#
		#.###
		..#..
	*/
	case1 := [][]bool{
		{T, T, T, T, T, F},
		{T, F, F, T, T, F},
		{F, F, T, T, F, F},
		{F, T, F, F, F, F},
		{T, T, F, T, T, F},
		{F, F, F, F, F, F},
	}

	result1 := canWin(case1)
	fmt.Println(result1)

	/*
	TEST CASE 2)
	.....
	.....
	.....
	.....
	.....
	*/
		case2 := [][]bool{
		{T, T, T, T, T, F},
		{T, T, T, T, T, F},
		{T, T, T, T, T, F},
		{T, T, T, T, T, F},
		{T, T, T, T, T, F},
		{F, F, F, F, F, F},
	}
	result2 := canWin(case2)
	fmt.Println(result2)

	case3 := [][]bool{
		{F, T, T, F, F, F},
		{F, F, T, F, F, F},
		{F, F, T, F, F, F},
		{F, T, T, T, F, F},
		{F, F, T, F, F, F},
		{F, F, F, F, F, F},
	}
	result3 := canWin(case3)
	fmt.Println(result3)
}

func TestGetMemKey(t *testing.T) {
	case3 := [][]bool{
		{F, T, T, F, F, F},
		{F, F, T, F, F, F},
		{F, F, T, F, F, F},
		{F, T, T, T, F, F},
		{F, F, T, F, F, F},
		{F, F, F, F, F, F},
	}
	result := getMemKey(case3)

	fmt.Println(result)
}
