package main

import (
	"fmt"
	"testing"
)

func TestTransform(t *testing.T) {
	result := transform([]string{"1st","Jan","2012"})
	fmt.Println(result)
}

func TestReformatDate(t *testing.T) {
	result := reformatDate([]string{
		"20th Oct 2052",
		"6th Jun 1933",
		"26th May 1960",
		"20th Sep 1958",
		"16th Mar 2068",
		"25th May 1912",
		"16th Dec 2018",
		"26th Dec 2061",
		"4th Nov 2030",
		"28th Jul 1963"})

	fmt.Println(result)
}
func TestSplit(t *testing.T) {
	result := split("14th Nov 2030")

	fmt.Println(result)

}