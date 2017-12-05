package boardcover2

import (
	"testing"
	"fmt"
)

func TestBoardCover2_maxblock(t *testing.T) {

}

func TestBoardCover2_Index_GetLineIndex(t *testing.T) {
	// . . . . . . .
	// . . . . . . .
	// . # . . . . .

	i := Index{1, 2}
	w := 7
	s := 15

	r := i.GetLineIndex(w)

	if s != r {
		t.Error("인덱스 계산이 잘 못 되었습니다.", r,)
	}
}

func TestBoardCover2_convertBlockToIndex(t *testing.T) {
	i := [][]rune { {35,35,35}, {35,46,46}, }
	s := []Index { {0,0}, {1,0}, {2,0}, {0,1} }

	r := convertBlockToIndex(i)

	if fmt.Sprint(s) != fmt.Sprint(r) {
		t.Error("잘 못 변형 되었습니다.", r)
	}
}

func TestBoardCover2_convertInputBlock(t *testing.T) {
	i := [][]rune { {35,35,35}, {35,46,46}, }
	s := []rune { 35,35,35,35,46,46 }

	r := convertInputBlock(i)

	if fmt.Sprint(s) != fmt.Sprint(r) {
		t.Error("잘 못 변형 되었습니다.", r)
	}
}

func TestBoardCover2_rotateBlock90DegreeClockwise(t *testing.T) {
	// i1   i2  i3   i4
	// ###  ##  ..#  #.
	// #..  .#  ###  #.
	//      .#       ##
	// [# 35] [. 46]
	i1 := [][]rune { {35,35,35}, {35,46,46}, }
	i2 := [][]rune { {35,35}, {46,35}, {46,35}, }
	i3 := [][]rune{ {46,46,35}, {35,35,35}, }
	i4 := [][]rune{ {35,46}, {35,46}, {35,35} }

	r1 := rotateBlock90DegreeClockwise(i1)
	if fmt.Sprint(r1) != fmt.Sprint(i2) {
		t.Error("변형된 값이 서로 다릅니다.", r1, )
	}

	r2 := rotateBlock90DegreeClockwise(r1)
	if fmt.Sprint(r2) != fmt.Sprint(i3) {
		t.Error("변형된 값이 서로 다릅니다.", r2, )
	}

	r3 := rotateBlock90DegreeClockwise(r2)
	if fmt.Sprint(r3) != fmt.Sprint(i4) {
		t.Error("변형된 값이 서로 다릅니다.", r3, )
	}
}