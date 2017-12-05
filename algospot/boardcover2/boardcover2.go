package boardcover2

import "fmt"

/*
Algospot
	Chapter 11.3 Board Cover 2
	(https://algospot.com/judge/problem/read/BOARDCOVER2)
*/

const ProblemTitle = "Board Cover 2"

type BoardCover2 struct {
	Board []rune // unicode slice, [# 35] [. 46]
	Width int
	Height int
	BlockIndices [][]Index
}

type Index struct { // index of block
	x int
	y int
}

func (i Index) GetLineIndex(boardWidth int) int {
	return i.x + (i.y * boardWidth)
}

func (p BoardCover2) GetProblemTitle() string {
	return ProblemTitle
}

func (p *BoardCover2) ReadProblem() {
	var H, W, R, C int
	fmt.Scan(&H)
	fmt.Scan(&W)
	fmt.Scan(&R)
	fmt.Scan(&C)

	// Read a board
	p.Board = make([]rune, 0)
	p.Width, p.Height = W, H

	for i := 0; i < H; i++ {
		var boardLine string
		fmt.Scanln(&boardLine)

		p.Board = append(p.Board, []rune(boardLine)...)
	}

	// Read a block
	block := make([][]rune, R)

	for i := 0; i < R; i++ {
		var blockLine string
		fmt.Scanln(&blockLine)

		block[i] = make([]rune, C)

		for j, r := range []rune(blockLine) {
			block[i][j] = r
		}
	}

	// Convert a block to indices
	numberOfDirections := 4
	p.BlockIndices = make([][]Index, numberOfDirections)
	for i := 0; i < numberOfDirections; i++ {
		rotateCount := i

		for rotateCount > 0 {
			block = rotateBlock90DegreeClockwise(block)
			rotateCount--
		}

		p.BlockIndices[i] = convertBlockToIndex(block)
	}
}

func (p BoardCover2) SolveProblem() interface{} {

	return 0
}

func maxblock() {

}

func convertBlockToIndex(block [][]rune) []Index {
	r := make([]Index, 0)

	for i := 0; i < len(block); i++ {
		for j := 0; j < len(block[i]); j++ {
			if block[i][j] == 35 {
				r = append(r, Index{j, i})
			}
		}
	}

	return r
}

func convertInputBlock(block [][]rune) []rune {

	var r []rune

	for i := 0; i < len(block); i++ {
		r = append(r, block[i]...)
	}

	return r
}

func rotateBlock90DegreeClockwise(block [][]rune) [][]rune {
	// ###  ##  ..#  #.
	// #..  .#  ###  #.
	//      .#       ##
	//
	// [0,1]->[0,0]  [0,0]->[0,1]  [0,1]->[1,1]  [1,1]->[0,1]  [2,1]->[0,2]  [0,2]->[1,2]
	// (width, height) -> (height, width)
	// (y, x) -> (x, height-1 - y )
	// (0, 0) -> (0, 1 - 0) -> (0, 1)
	// (1, 0) -> (0, 1 - 1) -> (0, 0)

	h := len(block)
	w := len(block[0])
	r := make([][]rune, w)

	for y := 0; y < w; y++ {
		r[y] = make([]rune, h)
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r[x][h-1 - y] = block[y][x]
		}
	}

	return r
}

