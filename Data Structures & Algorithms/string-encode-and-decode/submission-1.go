type Solution struct{}
const delimeter = '#'

func (s *Solution) Encode(strs []string) string {
	var encoded string
	for _, str := range strs {
		encoded += strconv.Itoa(len(str))
		encoded += string(delimeter)
		encoded += str
	}

	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	var strLen string
	
	i := 0
	for i < len(encoded) {
		if encoded[i] == delimeter {
			length, _ := strconv.Atoi(strLen)
			result = append(result, encoded[i+1:i+1+length])
			i = i+1+length
			strLen = ""
			continue
		}
		strLen = strLen + string(encoded[i])
		i++
	}

	return result
}
