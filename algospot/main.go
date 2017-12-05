package main

import (
	"fmt"
	"os"

	"github.com/GeunhoKim/go-algorithm/algospot/fanmeeting"
	"github.com/GeunhoKim/go-algorithm/algospot/fence"
	"github.com/GeunhoKim/go-algorithm/algospot/quadtree"
	"github.com/GeunhoKim/go-algorithm/algospot/quantization"
	"github.com/GeunhoKim/go-algorithm/algospot/matchorder"
	"github.com/GeunhoKim/go-algorithm/algospot/boardcover2"
	"github.com/GeunhoKim/go-algorithm/algospot/josephus"
)

/*
Algospot
	Main Executor
*/
const DEBUG = true

type Problem interface {
	GetProblemTitle() string
	ReadProblem()
	SolveProblem() interface{}
}

var problems = map[string]Problem{
	"fanmeeting": &fanmeeting.FanMeeting{},
	"quadtree":   &quadtree.QuadTree{},
	"fence":      &fence.Fence{},
	"quantization": &quantization.Quantization{},
	"matchorder": &matchorder.MatchOrder{},
	"boardcover2": &boardcover2.BoardCover2{},
	"josephus": &josephus.Josephus{},
}

func main() {
	problemName := os.Args[1]

	problem := problems[problemName]

	title := problem.GetProblemTitle()
	printd(title)

	cases := readNumberOfCases()
	for cases > 0 {
		problem.ReadProblem()
		result := problem.SolveProblem()

		fmt.Println(result)

		cases--
	}
}

func readNumberOfCases() int {
	var cases int
	fmt.Scanf("%d", &cases)

	return cases
}

func printd(args ...interface{}) {
	if DEBUG {
		fmt.Println(args)
	}
}
