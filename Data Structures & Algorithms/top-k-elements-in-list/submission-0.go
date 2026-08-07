func topKFrequent(nums []int, k int) []int {
	// create a map frequency to int
	mapIntToFreq := make(map[int]int)
	for _, num := range nums {
		mapIntToFreq[num]++
	}

	mapFreqToInt := make(map[int][]int)
	for num, freq := range mapIntToFreq {
		mapFreqToInt[freq] = append(mapFreqToInt[freq], num)
	}

	n := len(nums)
	var result []int
	for n > 0 && k > 0 {
		result = append(result, mapFreqToInt[n]...)
		k -= len(mapFreqToInt[n])
		n--
	}
	
	return result
}
