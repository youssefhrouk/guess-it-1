package main

import (
	"fmt"
	"os"

	mad "mad/math-skills"
)

func main() {
	var num int
	arr := []int{}

	for {
		_, err := fmt.Fscan(os.Stdin, &num)
		if err != nil {
			os.Exit(0)
		}
		arr = append(arr, num)

		low := mad.Median(arr) - int(float64((arr))*1.5)
		hight := mad.Median(arr) + int(float64(mad.MedianAbsoltDev(arr))*1.5)

		fmt.Println(low, hight)
	}
}
