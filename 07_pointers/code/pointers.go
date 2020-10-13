package main

import "fmt"

// func main() {
// 	var name string
// 	var namePointer *string

// 	name = "Beyonce"
// 	namePointer = &name
// 	nameValue := *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name value:", nameValue)
// }

// // ******************************************************

// func main() {
// 	var name string = "Beyonce"
// 	var namePointer *string = &name
// 	var nameValue = *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name Value:", nameValue)

// }

// // ******************************************************

// func changeName(n string) {
// 	n = strings.ToUpper(n)
// }

// func main() {
// 	name := "Elvis"
// 	changeName(name)
// 	fmt.Println(name)
// }

// // ******************************************************

// Coordinates type
type Coordinates struct {
	X, Y float64
}

var c = Coordinates{X: 10.0, Y: 20.0}

func main() {
	coordinatesMemoryAddress := &c
	coordinatesMemoryAddress.X = 200

	fmt.Println(coordinatesMemoryAddress)
}
