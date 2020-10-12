package main

import (
	"fmt"
)

func getAvg(scores ...float64) (result float64) {
	total := 0.0

	for _, val := range scores {
		total += val
	}

	result = total / float64(len(scores))
	return
}

func checkPets(name string) bool {
	pets := map[string]string{
		"Theo":   "Cat",
		"Clancy": "Dog",
		"Spike":  "Dog",
		"Fig":    "Cat",
	}

	_, ok := pets[name]
	return ok
}

func addToList(items ...string) []string {
	groceries := [...]string{"banana", "apple", "pineapple", "mango", "pear"}
	slicedGroceries := groceries[1:4]
	finalList := slicedGroceries

	for _, value := range items {
		finalList = append(finalList, value)
	}

	return finalList
}

func main() {
	fmt.Println(getAvg(30.0, 50.0, 65.0))
	fmt.Println(checkPets("NA"))
	fmt.Println(addToList("cherry", "orange"))
}
