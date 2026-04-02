package functions

func Median(data []int, cnt int) (res int) {
	if cnt == 0 {
		return 0
	}
	switch cnt % 2 {
	case 1:
		res = data[cnt/2]
	case 0:
		switch (data[cnt/2] + data[(cnt/2)-1]) % 2 {
		case 1:
			res = ((data[cnt/2] + data[(cnt/2)-1]) / 2) + 1
		case 0:
			res = (data[cnt/2] + data[(cnt/2)-1]) / 2
		}
	}
	return res
}
