func longestCommonPrefix(strs []string) string {
    // iterate through a string at a time, and find the common prefix
	// then pass the common prefix to the next iteration
	// after the loop that common prefix is the longest
	// time complexity: O(n*k) with n is strs length, k is longest length of string inside strs
	commonPrefix := strs[0]
	for _, str := range strs {
		for !strings.HasPrefix(str, commonPrefix) {
			commonPrefix = commonPrefix[:len(commonPrefix)-1]
		}
	}
	
	return commonPrefix
}
