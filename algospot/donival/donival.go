package donival

const F = false
const T = true

var numberOfNodes, startNode = 5, 0

var connected = [5][5]bool {
	{F, T, T, T, F},
	{T, F, F, F, T},
	{T, F, F, F, F},
	{T, F, F, F, F},
	{F, T, F, F, F},
}

var numConnected = [5]int {3, 2, 1, 1, 1}

var searched = [5][2]float64{
	{-1.0, -1.0},
	{-1.0, -1.0},
	{-1.0, -1.0},
	{-1.0, -1.0},
	{-1.0, -1.0},
}

func search(currentNode int, length int) float64 {
	if length == 0 {
		if currentNode == startNode {
			return 1.0
		} else {
			return 0.0
		}
	}

	var result = &searched[currentNode][length-1]

	if *result > 0 {
		return *result
	}

	*result = 0.0

	for beforeNode := 0; beforeNode < numberOfNodes; beforeNode++ {
		if connected[currentNode][beforeNode] {
			*result += search(beforeNode, length - 1) / float64(numConnected[beforeNode])
		}
	}

	return *result
}

