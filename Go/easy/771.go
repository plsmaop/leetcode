func numJewelsInStones(J string, S string) int {
    var table [130]bool
    for _, j := range[]rune(J) {
        table[j] = true
    }
    ans:= 0
    for _, s := range[]rune(S) {
        if (table[s]) {
            ans++   
        }
    }
    return ans
}