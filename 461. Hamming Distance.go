/*
汉明距离广泛应用于多个领域。在编码理论中用于错误检测，在信息论中量化字符串之间的差异。
两个整数之间的汉明距离是对应位置上数字不同的位数。
*/
func hammingDistance(x int, y int) (ans int) {
    for s := x ^ y; s > 0; s >>= 1 {
        ans += s & 1
    }
    return
}