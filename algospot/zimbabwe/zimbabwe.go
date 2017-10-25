package zimbabwe

import (
	"strconv"
	"fmt"
)

var e = "422"
var c = "422"
var m = 2

var counter = 0

func predict(x string, selected []bool) {
	if len(x) == len(e) {
		var convx, _ = strconv.Atoi(x)
		var conve, _ = strconv.Atoi(e)

		if convx < conve {
			if convx % m == 0 {
				fmt.Println(convx)
				counter++
			}
		}
	} else {
		for i := 0; i < len(e); i++ {
			if !selected[i] && (i == 0 || c[i-1] != c[i] || selected[i-1]) {
				selected[i] = true
				predict(x + string(c[i]), selected)
				selected[i] = false
			}
		}
	}
}
