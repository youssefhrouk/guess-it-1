package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func Median(slice []int) int {
	length := len(slice)
	sort.Ints(slice)
	if length%2 == 0 {
		return (slice[length/2-1] + slice[length/2]) / 2
	}
	return slice[length/2]
}

func MedianAbsoluteDev(numbers []int) int {
	median := Median(numbers)
	absDeviations := make([]int, len(numbers))

	for i, num := range numbers {
		absDeviations[i] = int(math.Abs(float64(num) - float64(median)))
	}

	return Median(absDeviations)
}

func main() {
	var num int
	arr := []int{}

	for {
		_, err := fmt.Fscan(os.Stdin, &num)
		if err != nil {
			os.Exit(0)
		}
		arr = append(arr, num)

		if len(arr) > 1 {
			mad := MedianAbsoluteDev(arr)
			med := Median(arr)

			low := med - int(float64(mad)*1.5)
			high := med + int(float64(mad)*1.5)

			fmt.Println(low, high)
		}
	}
}
