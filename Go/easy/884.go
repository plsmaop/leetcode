import "strings"

func uncommonFromSentences(A string, B string) []string {
    table, table2 := make(map[string]int), make(map[string]int)
    words := strings.Split(A, " ")
    for _, word := range words {
        table[word]++
    }
    tmp := []string{}
    for _, word := range words {
        if table[word] == 1 {
            tmp = append(tmp, word)
        }
    }
    words2 := strings.Split(B, " ")
    
    for _, word := range words2 {
        table2[word]++
    }
    tmp2 := []string{}
    for _, word := range words2 {
        if table2[word] == 1 {
            tmp2 = append(tmp2, word)
        }
    }
    ans := []string{}
    for _, word := range tmp {
        if table2[word] == 0 {
            ans = append(ans, word)
        }
    }
    for _, word := range tmp2 {
        if table[word] == 0 {
            ans = append(ans, word)
        }
    }
    return ans
}