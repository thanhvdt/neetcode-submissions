func removeElement(nums []int, val int) int {
    // in-place algorithm is not using another data structure that costs more than O(1) space
    // 2 pointers: 1 for keep track of the valid num, the other iterate through the array
    i := 0
    for j := 0; j < len(nums); j++ {
        if nums[j] != val {
            nums[i] = nums[j]
            i++
        }
    }

    return i
}
