func topKFrequent(nums []int, k int) []int {
	// create a map frequency to int
	mapIntToFreq := make(map[int]int)
	for _, num := range nums {
		mapIntToFreq[num]++
	}

	buckets := make([][]int, len(nums) + 1)
	for num, freq := range mapIntToFreq {
		buckets[freq] = append(buckets[freq], num)
	}

	// from the most frequent in map, append to result
	n := len(nums)
	var result []int
	for n > 0 && k > 0 {
		result = append(result, buckets[n]...)
		k -= len(buckets[n])
		n--
	}
	
	return result
}
