package main

import "fmt"

//Stack allocated variable
func newIntStack() *int {
	vv := new(int) //allocated on the stack
	// since vv doesn't escape
	return vv
}

func main() { fmt.Println(*newIntStack()) }

// func main() {
// 	x := "GOPHERCON-2018"
// 	fmt.Println(x)
// }
