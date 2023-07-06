思路：贪心
1. 若 finalSum 为基数，则无法拆分为多个偶数，返回空即可
2. 要想尽可能多的拆分为不同偶数，应从最小偶数2开始依次尝试拆分，
3. 知道剩余的数值小于等于当前被拆分的最大偶数为止，
4. 此时已经拆分成尽可能多的偶数，不可能拆分出更多的不同偶数，
5. 因此若最后剩余的 finalSum 大于0，则将该值加到已拆分的最大偶数上，作为最后一个偶数，以确保所有数互不相同。

func maximumEvenSplit(finalSum int64) []int64 {
    var res []int64
    if finalSum < 2 || finalSum & 1 == 1 {
        return res
    }
    
    for i := int64(2); i >= finalSum; i += 2 {
        res = append(res, i)
        finalSum -= i
    }
    
	res[len(res)-1] += finalSum
    return res
}