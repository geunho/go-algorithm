package fanmeeting

import (
	"reflect"
	"testing"
)

func TestTransformGenderInputs(t *testing.T) {
	gendersStr := "FFFMMM"
	genders := []bool{F, F, F, M, M, M}

	transformed := transformGenderInputs(gendersStr)

	if !reflect.DeepEqual(genders, transformed) {
		t.Error(
			"For", gendersStr,
			"exepected", genders,
			"got", transformed,
		)
	}
}

func TestFindMenIndicesInMembers(t *testing.T) {
	members := []bool{F, F, F, M, M, M}
	menIndices := []int{3, 4, 5}

	foundIndices := findMenIndicesInMembers(&members)

	if !reflect.DeepEqual(menIndices, foundIndices) {
		t.Error(
			"For", members,
			"expected", menIndices,
			"got", foundIndices,
		)
	}

}

func TestIsAllMemberHug(t *testing.T) {
	menInMembers := []int{3, 4, 5}
	fans := []bool{M, M, M, F, F, F}

	isAllMemberHugged := isAllMemberHug(&menInMembers, &fans)

	if !isAllMemberHugged {
		t.Error(
			"For", menInMembers,
			"expected", true,
			"got", isAllMemberHugged,
		)
	}
}

func TestSolveProblem(t *testing.T) {
	//CASE 1)
	//members := []bool{F, F, F, M, M, M}
	//fans := []bool{M, M, M, F, F, F}
	//solution := 1

	//CASE 2)
	//members := []bool{F,F,F,F,F}
	//fans := []bool{F,F,F,F,F,F,F,F,F,F}
	//solution := 6

	//CASE 3)
	members := []bool{F, F, F, F, M}
	fans := []bool{F, F, F, F, F, M, M, M, M, F}
	solution := 2

	//CASE 4)
	//members := []bool{M, F, M, F, M, F, F, F, M, M, M, F, M, F}
	//fans := []bool{M, M, F, F, F, F, F, M, F, F, F, M, F, F, F, F,
	//	F, F, M, F, F, F, M, F, F, F, F, M, F, M, M, F, F, F, F, F, F, F}
	//solution := 2

	answer := SolveProblem(members, fans)

	if answer != solution {
		t.Error(
			"For", members,
			"expected", solution,
			"got", answer,
		)
	}
}
