package functions

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func SortList() (data []int, cnt int) {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <file>")
		os.Exit(1)
	}
	path := os.Args[1]
	BS, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	S := string(BS)
	SS := strings.Fields(S)
	for _, token := range SS {
		num, err := strconv.Atoi(token)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid value %q: skipping\n", token)
			continue
		}
		data = append(data, num)
		cnt++
	}
	if cnt == 0 {
		fmt.Println("Error: no valid numbers found in file")
		os.Exit(1)
	}
	slices.Sort(data)
	return data, cnt
}
