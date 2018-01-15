package ites

import "fmt"

/*
Algospot
	Chapter 19.6 ITES
	(https://algospot.com/judge/problem/read/ITES)
*/

const ProblemTitle = "ITES Analysis"

type ITES struct {
	K int // partial sum
	N int  // number of signals
}

func (p ITES) GetProblemTitle() string {
	return ProblemTitle
}

func (p *ITES) ReadProblem() {
	var K int
	fmt.Scanf("%d", &K)

	var N int
	fmt.Scanf("%d", &N)


}

func (p ITES) SolveProblem() interface{} {
	return 0
}

func generateNext(beforeNum uint) uint {
	var number = (beforeNum * 214013 + 2531011) % 4294967296 /* 2^32 */
	return number
}

func generateSignal(beforeNum uint) uint{
	var number = beforeNum % 10000 + 1
	return number
}

//cnt: 5
// in:
//{1,4,2,1,4,3,1,6} 에서 합이 7 인 부분 수열은 모두 5개
func analyze(K int, N int) int {
	signals := make([]uint, 0, N)
	var sum uint = 0
	count := 0

	var input uint = 1983
	var signal uint
	for i:=0; i<N; i++ {
		if sum == uint(K) {
			count++
		}

		signal = generateSignal(input)

		sum += signal
		signals = append(signals, signal)


		if sum >= uint(K) {
			if len(signals) > 0 {
				sum -= signals[0]

				if len(signals) > 1 {
					signals = signals[1:]
				} else {
					signals = signals[:0]
				}

			} else {
				// 큐는 비어있음
				signals = signals[:0]
			}
		}

		input = generateNext(input)
		//fmt.Println(signals)
		//fmt.Println(sum)
	}

	return count
}