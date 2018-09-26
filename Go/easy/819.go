import "strings"

func mostCommonWord(paragraph string, banned []string) string {
	bannedTable := make(map[string]bool)
	for _, word := range banned {
		bannedTable[word] = true
	}
	delimiter := " !?',;."
	wordList := strings.FieldsFunc(paragraph, func(c rune) bool {
		for _, d := range delimiter {
			if c == d {
				return true
			}
		}
		return false
	})
	wordCount := make(map[string]int)
	ans, ansCount := "", 0
	for _, word := range wordList {
		word = strings.ToLower(word)
		if bannedTable[word] {
			continue
		}
		wordCount[word]++
		if wordCount[word] > ansCount {
			ans = word
			ansCount = wordCount[word]
		}
	}
	return ans
}