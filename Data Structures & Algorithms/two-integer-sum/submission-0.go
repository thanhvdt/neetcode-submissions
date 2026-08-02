func twoSum(nums []int, target int) []int {
    mapValueToIndex := make(map[int]int)
	var secondValue int
	
	for i := 0; i < len(nums); i++ {
		secondValue = target - nums[i]
		secondIndex, exists := mapValueToIndex[secondValue]
		if exists {
			return []int{secondIndex, i}
		}
		// save the current iteration to hashmap
		mapValueToIndex[nums[i]] = i
	}

	return nil
}
