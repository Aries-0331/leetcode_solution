/*1. 贪心
要想选出k件物品，使和最大，那么贪心地选择前k大的物品时最优的，根据k取值，分为一下情况：
- k <= numOnes，选择k个1，最大和为k
- numOnes < k <= numOnes + numZeros，选择全部1和部分0，最大和为numOnes
- numOnes + numZeros < k，选择全部1和0，先择部分-1，数目为 k-numOnes-numZeros，最大和为numOnes-(k-numOnes-numZeros)
*/
func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
    if k <= numOnes {
        return k
    } else if k <= numOnes + numZeros {
        return numOnes
    } else {
        return numOnes - (k - numOnes - numZeros)
    }
}

2.
func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
    i := 0
    sum := 0
    for i < numOnes && i < k{
        sum += 1
        i++
    }
    for i < numOnes + numZeros && i < k{
        sum += 0
        i++
    }
    for i < numOnes + numZeros + numNegOnes && i < k{
        sum += -1
        i++
    }
    return sum
}