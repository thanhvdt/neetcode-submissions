func longestConsecutive(nums []int) int {
	// convert nums into a set
	setNums := make(map[int]bool)
	for _, num := range nums {
		setNums[num] = true
	}

	max := 0
	for num := range setNums {
		// found starting num of consecutive sequence
		if !setNums[num-1] {
			sequenceLength := 1
			for i := 1; setNums[num+i]; i++ {	
				sequenceLength++
			}

			if sequenceLength > max {
				max = sequenceLength
			}
		}
		
	}

	return max
}
