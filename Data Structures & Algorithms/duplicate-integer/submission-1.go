func hasDuplicate(nums []int) bool {
    numFrequency := make(map[int]int)

    for _, val := range(nums) {
        _, ok := numFrequency[val]
        if !ok {
            numFrequency[val] = 1
        } else {
            return true
        }
    }

    return false
}
