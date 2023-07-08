
func getRow(rowIndex int) []int {
	row := make([][]int, rowIndex+1)
	for i := range row {
		row[i] = make([]int, i+1)
		row[i][0], row[i][i] = 1, 1
		for j := 1; j < i; j++ {
			row[i][j] = row[i-1][j-1] + row[i-1][j]
		}
	}
	return row[rowIndex]
}