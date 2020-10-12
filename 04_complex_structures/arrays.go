package main

import (
	"fmt"
)

func main() {
	scores := [...]float64{9, 1.5, 4.5, 7, 8}
	for _, val := range scores {
		fmt.Println(val)
	}
}
