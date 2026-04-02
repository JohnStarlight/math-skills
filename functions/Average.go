package functions

import "math"

func Average(data []int, cnt int) (res int) {
	if cnt == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < cnt; i++ {
		sum += data[i]
	}
	res = int(math.Round(float64(sum) / float64(cnt)))
	return res
}
