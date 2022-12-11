package t54

func spiralOrder(matrix [][]int) []int {
	rowLen, colLen, row, col := len(matrix), len(matrix[0]), 0, 0
	directionMap, direction := map[int][]int{
		0: {1, 0},
	}, 0
	var res []int
	for len(res) < rowLen * colLen {
		res = append(res, matrix[row][col])
		dVals := directionMap[direction]
		row = row + dVals[0]
		col = col + dVals[1]
		if row {
			
		}
	}

}