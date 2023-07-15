/*
思路1：
根据杨辉三角的特性，可以构建二维数组，递推每层的值，
每当计算出第i行的值，都可以在线性时间复杂度内计算出第i+1行的值
*/
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
/*
优化：
思路1中的二维数组，计算第i+1行时，仅用到第i行的数据，
可使用滚动数组的思想优化空间复杂度。
ps：滚动数组是DP 中的一种编程思想。 简单的理解就是让数组滚动起来，每次都使用固定的几个存储空间，来达到压缩，起到优化空间，节省存储空间的作用。 主要应用在递推或动态规划中（如01背包问题）。
*/
func getRow(rowIndex int) []int {
	var cur, pre []int
	for i := 0; i < rowIndex; i++ {
		cur = make([]int, i+1)
		cur[0],cur[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = pre[j-1] + pre[j]
		}
		pre = cur
	}
	return pre
}
