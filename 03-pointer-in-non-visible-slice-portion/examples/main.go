package main

import (
	"fmt"
)

//Slice error prior to Go 1.2
func main() {
	a := []*int{new(int), new(int)}
	fmt.Println(a)
	b := a[:1]
	fmt.Println(b) // [&0]

	// second element is not garbage collected, because it's *still* accessible
	c := b[:2] //[&0[
	fmt.Println(c)
}

// //3 index slice
// func main() {
// 	a := []*int{new(int), new(int)}
// 	fmt.Println(a)

// 	// Using the 3- index slice operation
// 	b := a[:1:1]
// 	fmt.Println(b)
// 	c := b[:2]
// 	fmt.Println(c)
// }
