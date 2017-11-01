package blockgame

import "fmt"

/*
홀수번 블럭을 뒀을 때 빈 공간이 연속된 곳이 없는 경우가 '하나라도' 있다면 true, 그렇지 않은 경우 false

입력: status bool[6][6], turn int = 1 부터 시작

''  '   ''  ''  '    '
    '    '  '   ''  ''
*/

var mem = make(map[string]bool)

func canWin(status [][]bool) bool {
	if !isPossibleToPlayGame(status) {
		return false
	}

	if value, ok := mem[getMemKey(status)]; ok {
		return value
	}

	result := false

	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			// 6 Cases, base on specific point.
			// CASE 1: * *
			if status[x][y] && status[x+1][y] {
				status[x][y], status[x+1][y] = false, false
				result = result || !canWin(status) // 다음 턴의 결과가 패배라면 이번 턴은 승리한다.
				status[x][y], status[x+1][y] = true, true
			}
			// CASE 2:  * *
			//            *
			if status[x][y] && status[x+1][y] && status[x+1][y+1] {
				status[x][y], status[x+1][y], status[x+1][y+1] = false, false, false
				result = result || !canWin(status) // 다음 턴의 결과가 패배라면 이번 턴은 승리한다.
				status[x][y], status[x+1][y], status[x+1][y+1] = true, true, true
			}
			// CASE 3:  * *
			//          *
			if status[x][y] && status[x+1][y] && status[x][y+1] {
				status[x][y], status[x+1][y], status[x][y+1] = false, false, false
				result = result || !canWin(status) // 다음 턴의 결과가 패배라면 이번 턴은 승리한다.
				status[x][y], status[x+1][y], status[x][y+1] = true, true, true
			}
			// CASE 4:  *
			//          *
			if status[x][y] && status[x][y+1] {
				status[x][y], status[x][y+1] = false, false
				result = result || !canWin(status) // 다음 턴의 결과가 패배라면 이번 턴은 승리한다.
				status[x][y], status[x][y+1] = true, true
			}
			// CASE 5:  *
			//          * *
			if status[x][y] && status[x][y+1] && status[x+1][y+1] {
				status[x][y], status[x][y+1], status[x+1][y+1] = false, false, false
				result = result || !canWin(status) // 다음 턴의 결과가 패배라면 이번 턴은 승리한다.
				status[x][y], status[x][y+1], status[x+1][y+1] = true, true, true
			}
			// CASE 6:  *
			//        * *
			if x > 0 && status[x][y] && status[x][y+1] && status[x-1][y+1] {
				status[x][y], status[x][y+1], status[x-1][y+1] = false, false, false
				result = result || !canWin(status) // 다음 턴의 결과가 패배라면 이번 턴은 승리한다.
				status[x][y], status[x][y+1], status[x-1][y+1] = true, true, true
			}
		}
	}

	mem[getMemKey(status)] = result

	return result
}

func isPossibleToPlayGame(status [][]bool) bool {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if (status[x][y] && status[x+1][y]) || (status[x][y] && status[x][y+1]) {
				return true
			}
		}
	}

	return false
}

func getMemKey(status [][]bool) string {
	return fmt.Sprint(status)
}