package main

import (
	"bytes"
	"fmt"
)

func main() {
	bs := []byte{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd'}
	fmt.Printf("%d %d %q\n", len(bs), cap(bs), bs[:cap(bs)])
	ps := bytes.Split(bs, []byte{','})
	for _, p := range ps {
		fmt.Printf("%d %d %q\n", len(p), cap(p), p[:cap(p)])
	}
}
