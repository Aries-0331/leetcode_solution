//哈希表
func findRepeatNumber(nums []int) int {
    m := make(map[int]int)
    for _, v := range nums {
        if _, ok := m[v]; ok {
            return v
        }
        m[v] = 1
    }
    return -1
}

//原地交换
func findRepeatNumber(nums []int) int {
    i := 0
    for i < len(nums) {
        if nums[i] == i {
            i++
            continue
        }
        if nums[nums[i]] == nums[i] {
            return nums[i]
        }
        nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
    }
    return -1
}