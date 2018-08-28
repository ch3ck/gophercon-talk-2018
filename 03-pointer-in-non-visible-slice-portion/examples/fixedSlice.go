package main

import (
	"fmt"
)

//START OMIT
func main() {
	a := []*int{new(int), new(int)}
	fmt.Println(a)

	// Using the 3- index slice operation
	b := a[:1:1]
	fmt.Println(b)
	c := b[:2]
	fmt.Println(c)
}

//END OMIT
