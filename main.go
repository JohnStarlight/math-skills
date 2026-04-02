package main

import (
	"fmt"

	"math-skills/functions"
)

func main() {
	data, cnt := functions.SortList()
	fmt.Printf("Average: %d\nMedian: %d\nVariance: %d\nStandard Deviation: %d\n", functions.Average(data, cnt), functions.Median(data, cnt), functions.Variance(data, cnt), functions.StanDev(data, cnt))
}
