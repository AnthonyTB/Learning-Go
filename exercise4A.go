package main

import (
	"fmt"
)

func average(nums ...float64) (averageVal float64) {
	total := 0.0
	for _, val := range nums {
		total += val
	}
	averageVal = total / float64(len(nums))
	return
}

func mainoldd() {
	fmt.Println(average(60.50, 55.00, 90.50))
}
