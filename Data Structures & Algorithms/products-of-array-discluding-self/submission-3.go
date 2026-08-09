func productExceptSelf(nums []int) []int {
	// brute-force: calculate the product of all and loop and do division at each num
	// without division ops: use 2 arrays, 1 is product of all nums to the left, 1 is product of all nums to the right
	leftProducts := make([]int, len(nums))
	leftProducts[0] = 1
	rightProducts := make([]int, len(nums))
	rightProducts[len(nums)-1] = 1
	leftProduct := 1
	rightProduct := 1
	for i := 0; i < len(nums)-1; i++ {
		leftProduct *= nums[i]
		leftProducts[i+1] = leftProduct
	}
	for i := len(nums)-1; i >= 1; i-- {
		rightProduct *= nums[i]
		rightProducts[i-1] = rightProduct
	}

	var result []int
	for i := 0; i < len(nums); i++ {
		result = append(result, leftProducts[i]*rightProducts[i])
	}

	return result
	// todo: optimize by using result array as the leftProducts, and iterate from the end of nums to calculate the final result, no need to create 2 new arrays -> O(1) extra space
}
