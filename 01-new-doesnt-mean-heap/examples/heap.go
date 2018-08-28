package main

import "fmt"

// START OMIT
//Stack allocated variable
func newIntStack() *int {
	vv := new(int) //allocated on the stack
	// since vv doesn't escape
	return vv
}

func main() { fmt.Println(*newIntStack()) }

// END OMIT
