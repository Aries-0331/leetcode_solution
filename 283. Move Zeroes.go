/*
思路：把非0数据逐个左移，一个指针遍历非0数据，一个指针记录要左移的位置，然后把最后一个记录的位置后全置0。
*/
func moveZeroes(nums []int)  {
    if nums == nil {
        return
    }
    j := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[j] = nums[i]
            j++
        }
    }
    for i := j; i < len(nums); i++ {
        nums[i] = 0
    }
}

/*
思路：累死快速排序的做法，以0作为基准，不为0的移到左边，为0的移到右边，使用两个指针，一个遍历数据，一个记录位置用于交换
*/
func moveZeroes(nums []int)  {
    if nums == nil {
        return
    }
    j := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0{
            tmp := nums[i]
            nums[i] = nums[j]
            nums[j] = tmp
            j++
        }
    }
}