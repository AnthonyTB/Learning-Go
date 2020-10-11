package main

import (
	"fmt"
)

func main() {
	sentence := "This is a sentence"

	for index, strVal := range sentence {
		if index%2 != 0 {
			fmt.Println(string(strVal))
		}
	}
}
