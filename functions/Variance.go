package functions

import "math"

func Variance(data []int, cnt int) (res int) {
	av := Average(data, cnt)
	tmp := 0
	for i := 0; i < cnt; i++ {
		tmp += (data[i] - av) * (data[i] - av)
	}
	res = int(math.Round(float64(tmp) / float64(cnt)))
	return res
}
