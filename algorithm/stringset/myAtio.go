package stringset

import (
	"math"
)

func myAtoi(s string) int {
	flag := 1
	num := 0

	begin := len(s)

	for i, k := range s {
		if k == 32 {
			continue
		}

		if k == 45 {
			flag = -1
			begin = i + 1
			break
		}

		if k == 43 {
			flag = 1
			begin = i + 1
			break
		}

		if k >= 48 && k <= 57 {
			begin = i
			break
		}

		return 0
	}

	for i := begin; i < len(s); i++ {
		k := s[i]
		if k >= 48 && k <= 57 {
			num = num*10 + numMap[k]
			if num*flag < math.MinInt32 {
				return math.MinInt32
			}

			if num*flag > math.MaxInt32 {
				return math.MaxInt32
			}
		} else {
			break
		}
	}

	return num * flag
}

var numMap = map[uint8]int{
	48: 0,
	49: 1,
	50: 2,
	51: 3,
	52: 4,
	53: 5,
	54: 6,
	55: 7,
	56: 8,
	57: 9,
}
