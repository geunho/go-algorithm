package main

import (
	"testing"
	"fmt"
)

func TestCounting(t *testing.T) {
	r0 := counting("00110")
	fmt.Println(r0)

	//r1 := counting("10001")
	//fmt.Println(r1)
	//
	//r2 := counting("10101")
	//fmt.Println(r2)
}


/*
func main() {
	var names []string
	var names_size int
	fmt.Scanf("%d", &names_size)
	​
	for i := 0; i < names_size; i++ {
		var names_item string
		fmt.Scanf("%s", &names_item)
		names = append(names, names_item)
		​
	}
	​
	var res []string = getSortedList(names)
	for res_i := range res {
		fmt.Println(res[res_i])
	}
}
*/
/*
func main() {
	var dates []string
	var dates_size int
	fmt.Scanf("%d", &dates_size)

	for i := 0; i < dates_size; i++ {
		var dates_item string
		fmt.Scanf("%s", &dates_item)
		dates = append(dates, dates_item)

	}

	var res []string = reformatDate(dates)
	for res_i := range res {
		fmt.Println(res[res_i])
	}
}
*/