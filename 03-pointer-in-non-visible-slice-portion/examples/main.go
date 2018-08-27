package main

import (
	"fmt"
)

//Slice error prior to Go 1.2
func main() {

	a := []*int{new(int), new(int)}
	fmt.Println(a)
	a = a[:1]
	fmt.Println(a) // [&0]

	// second element is not garbage collected, because it's *still* accessible
	a = a[:2] //[&0[
	fmt.Println(a)

	// // Using the 3- index slice operation
	// a = a[:1:1]
	// fmt.Println(a)
}

// //
// func main() {
// 	bs := []byte{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd'}
// 	fmt.Printf("%d %d %q\n", len(bs), cap(bs), bs[:cap(bs)])
// 	ps := bytes.Split(bs, []byte{','})
// 	for _, p := range ps {
// 		fmt.Printf("%d %d %q\n", len(p), cap(p), p[:cap(p)])
// 	}
// }

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	t := s[0:3:3]
// 	fmt.Println("s", len(s), cap(s), s)
// 	fmt.Println("t", len(t), cap(t), t)

// 	t = t[:4]
// 	t[3] = 0
// 	t = t[:3]

// 	fmt.Println("s", len(s), cap(s), s)
// 	fmt.Println("t", len(t), cap(t), t)
// }
