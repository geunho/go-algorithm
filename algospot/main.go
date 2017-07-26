package main

import (
	"fmt"
)

/*
Algospot
	Problem 1.2 FESTIVAL
	(https://algospot.com/judge/problem/read/FESTIVAL)
*/
var DEBUG = true

func main() {
	printd("Festival Problem")

	// read INPUTS
	// -------------------------------------------------
	var cases int
	fmt.Scanf("%d", &cases)

	for cases > 0 {
		var N int
		var L int

		fmt.Scanf("%d", &N)
		fmt.Scanf("%d", &L)

		prices := [3]int{1, 2, 3}
		solveProblem(1, 1, *prices)

		cases--
	}

}

func printd(args ...string) {
	if DEBUG {
		println(args)
	}
}

func solveProblem(N int, L int, prices []int) {
	var result float64

	// print OUTPUT
	// -------------------------------------------------
	fmt.Printf("%f\n", result)
}
