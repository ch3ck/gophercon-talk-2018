package main

import "fmt"

func newIntStack() *int {

	vv := new(int) // HL

	return vv
}

func main() { fmt.Println(*newIntStack()) }
