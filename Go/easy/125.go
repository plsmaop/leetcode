import "unicode"

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	var tmp string
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			tmp += string(unicode.ToLower(c))
		}
	}
	length := int(len(tmp) / 2)
	for i := 0; i < length; i++ {
		if tmp[i] != tmp[len(tmp)-1-i] {
			return false
		}
	}
	return true
}