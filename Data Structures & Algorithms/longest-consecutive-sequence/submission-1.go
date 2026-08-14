func longestConsecutive(nums []int) int {
	// convert nums into a set
	setNums := make(map[int]int)
	for _, num := range nums {
		setNums[num] = 1
	}

	max := 0
	for num := range setNums {
		// found starting num of consecutive sequence
		if _, exists := setNums[num-1]; !exists {
			flag := true
			sequenceLength := 1
			i := 1
			for flag {
				if _, nextNumExists := setNums[num+i]; !nextNumExists {
					flag = false
					continue
				}	
				i++
				sequenceLength++
			}

			if sequenceLength > max {
				max = sequenceLength
			}
		}
		
	}

	return max
}
