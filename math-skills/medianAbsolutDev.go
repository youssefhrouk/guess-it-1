package mad

import "math"

func MedianAbsoltDev(numbers []int) int {
	median := Median(numbers)

	// Calculate the absolute deviations from the median
	absDeviations := make([]int, len(numbers))
	for i, num := range numbers {
		absDeviations[i] = int(math.Abs(float64(num) - float64(median)))
	}

	// Calculate the median of the absolute deviations
	return Median(absDeviations)
}
