func findTheDifference(s string, t string) byte {
	xor := 0
	for _, c := range s {
		xor ^= int(c)
	}
	for _, c := range t {
		xor ^= int(c)
	}
	return byte(rune(xor))
}