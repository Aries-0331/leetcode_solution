/*
思路：Brian Kernighan 算法，对于任意整数 x，令 x=x & (x−1)，该运算将 x 的二进制表示的最后一个 1 变成 0，因此对 x 重复此操作直至 x 为 0，操作次数即为 x 的二进制表示包含的 1 的数量。
*/
func countBits(n int) []int {
    ret := make([]int, n+1)
    for i := 0; i <= n; i++ {
        count, tmp := 0, i
        for tmp != 0 {
            tmp = tmp & (tmp - 1)
            count++
        }
        ret[i] = count
    }
    return ret
}

/*
动态规划
*/
func countBits(n int) []int {
    bits := make([]int, n+1)
    highBit := 0
    for i := 1; i <= n; i++ {
        if i&(i-1) == 0 {
            highBit = i
        }
        bits[i] = bits[i-highBit] + 1
    }
    return bits
}