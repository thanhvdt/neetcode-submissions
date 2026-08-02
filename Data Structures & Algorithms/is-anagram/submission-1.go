func isAnagram(s string, t string) bool {
    firstCharFreq := countCharsInString(s)
    secondCharFreq := countCharsInString(t)

    if len(firstCharFreq) != len(secondCharFreq) {
        return false
    }
    
    for key, val := range(firstCharFreq) {
        if val != secondCharFreq[key] {
            return false
        }
    }

    return true
}

func countCharsInString(s string) map[rune]int {
    charFreq := make(map[rune]int)

    for _, char := range(s) {
        if _, ok := charFreq[char]; !ok {
            charFreq[char] = 1
        } else {
            charFreq[char] += 1
        }
    }

    return charFreq
}
