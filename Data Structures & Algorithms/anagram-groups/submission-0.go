func groupAnagrams(strs []string) [][]string {
	mapFreqToStr := make(map[[26]int][]string)
	for _, singleStr := range(strs) {
		// count freq of character in singleStr
		var mapKey [26]int
		for _, singleChar := range(singleStr) {
			// find key number (singleChar - 'a')
			key := singleChar - 'a'
			// add freq to mapKey
			mapKey[key]++
		}

		// check if key exists in hashmap
		// - if yes, add string into the value of that element in hashmap
		// - if no, create a new element in the hashmap
		mapFreqToStr[mapKey] = append(mapFreqToStr[mapKey], singleStr)
	}

	var result [][]string
	// loop through all elements in hashmap and add in to an array to return
	for _, anagramGroup := range(mapFreqToStr) {
		result = append(result, anagramGroup)
	}

	return result
}
