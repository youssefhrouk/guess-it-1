package mad

import "sort"

func Median(slice []int) int {
	length := len(slice)
	sort.Ints(slice)
	if len(slice)%2 == 0 {
		return (slice[length/2-1] + slice[length/2]) / 2
	}

	return slice[length/2]
}
