// variables need to decleared before
package main

import "fmt"

func main() {
	// fmt.Println(speed)  this won't work because, a variable cannot be called before initialization.
	var speed int // will be initialized with a value of 0
	fmt.Println(speed)
	speed = 100 // speed reassigned to 100
	fmt.Println(speed)
}