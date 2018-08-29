package main

import (
	"fmt"
)

//START OMIT
func main() {
	a := []*int{new(int), new(int)} // HL
	fmt.Println(a)
	b := a[:1] // HL
	fmt.Println(b)

	// second element is not garbage collected, because it's *still* accessible
	c := b[:2] // HL
	fmt.Println(c)
}

//END OMIT
