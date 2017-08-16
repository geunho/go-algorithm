package main

import (
	"github.com/GeunhoKim/go-algorithm/algospot/picnic"
	"fmt"
)

/*
Algospot
	Main Executor
*/
const DEBUG = true

func main() {
	printd(picnic.ProblemTitle)

	picnic.ReadNumberOfCases()

	nCases := picnic.Cases
	for nCases > 0 {
		picnic.ReadProblem()


	}

}

func printd(args ...string) {
	if DEBUG {
		fmt.Println(args)
	}
}