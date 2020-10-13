package main

import (
	math "fem-intro-to-go/05_toolkit/code/utils"
	"fmt"
)

func calculateImportantData() int {
	totalValue := math.Add(1, 2, 3, 4)
	return totalValue
}

func main() {
	total := calculateImportantData()
	fmt.Println(total)
	fmt.Println("Packages!")
}
