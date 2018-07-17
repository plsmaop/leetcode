func toLowerCase(str string) string {
    ans := ""
    for _, s := range[]rune(str) {
        if s <= 90 && s >= 65 {
            ans += (string(s+32))
        } else {
            ans += string(s)
        }
    }
    return ans
}