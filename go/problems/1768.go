package problems

func MergeAlternately(word1 string, word2 string) string {
	newWord := ""
	longest := getLongest(word1, word2)
	i, j := 0, 0

	for i < len(longest) {
		if i < len(word1) {
			newWord += string(word1[i])
		}
		if j < len(word2) {
			newWord += string(word2[j])
		}
		i++
		j++
	}

	return newWord
}

func getLongest(word1 string, word2 string) string {
	longest := ""

	if len(word1) > len(word2) {
		longest = word1
	} else {
		longest = word2
	}
	return longest
}
