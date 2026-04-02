package functions

import "math"

func StanDev(data []int, cnt int) (res int) {
	av := Variance(data, cnt)
	res = int(math.Round(math.Sqrt(float64(av))))
	return res
}
