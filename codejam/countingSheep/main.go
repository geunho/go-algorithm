package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

/*
Google Code jam, 2016
Qualification Round
	Problem A. Counting Sheep
	(https://code.google.com/codejam/contest/dashboard?c=6254486)
 */
func main() {
	println("Counting Sheep Problem")

	pwd, _ := os.Getwd()
	//smallinput := "A-small-practice.in.txt"
	largeinput := "A-large-practice.in.txt"

	input, err := os.Open(pwd + "/codejam/countingSheep/" + largeinput)
	if err != nil { println(err) }
	defer input.Close()

	scanner := bufio.NewScanner(input)

	// read the first line as T, the number of tests
	scanner.Scan()
	firstline := scanner.Text()
	T, _ := strconv.Atoi(firstline)

	println("Number of tests: " + firstline)
	println()
	println("-----------------------------------")

	results := make([]string, T)

	index := 0
	for scanner.Scan() {
		line := scanner.Text()

		N, err := strconv.Atoi(line)
		if err == nil {
			solveProblem(N, T, &(results[index]))
		}

		index++
	}

	writeResults(&results)
}

func solveProblem(N int, T int, resultToReturn *string) {
	println("Target N:")
	println(N)

	checks := map[int]bool{
		0: false,
		1: false,
		2: false,
		3: false,
		4: false,
		5: false,
		6: false,
		7: false,
		8: false,
		9: false,
	}

	result := N
	resultIter := 0

	for i := 1; i <= T; i++ {
		var num = N * i
		checkNums(num, &checks)

		if checkRules(&checks) {
			result = num
			resultIter = i
			break
		}
	}

	*resultToReturn = strconv.Itoa(result)
	if(result == 0 || (resultIter == 100 && result == N)) {
		*resultToReturn = "INSOMNIA"
	}

	println("Sheep name that she fall asleep:")
	println(result)
	println("Iterations:")
	println(resultIter)
	println("-----------------------------------")
}

func checkRules(checks *map[int]bool) bool {
	for i := 0; i <= 9; i++ {
		if !(*checks)[i] {
			return false
		}
	}
	return true
}

func checkNums(num int, checks *map[int]bool) {
	digit := num % 10
	(*checks)[digit] = true;

	var next int = num / 10

	if next == 0 || next == digit {
		return
	} else {
		checkNums(num / 10, checks)
	}
}

func writeResults(results *[]string) {
	pwd, _ := os.Getwd()
	output, err := os.Create(pwd + "/codejam/countingSheep/output2.txt")
	if err != nil { println(err) }
	defer output.Close()

	writer := bufio.NewWriter(output)

	for index, result := range *results {
		line := fmt.Sprintf("Case #%d: %s\n", index+1, result)
		writer.WriteString(line)
	}
	writer.Flush()
}