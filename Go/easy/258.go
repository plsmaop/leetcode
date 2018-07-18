import (
	"strconv"
)

func addDigits(num int) int {
	for num >= 10 {
		tmp := 0
		for _, digit := range strconv.Itoa(num) {
			temp, _ := strconv.Atoi(string(digit))
			tmp += temp
		}
		num = tmp
	}
	return num
}