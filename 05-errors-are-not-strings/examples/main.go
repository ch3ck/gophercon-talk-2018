package main

import (
	"fmt"
	"log"
	"net"
)

//START OMIT
func main() {
	conn, err := net.Dial("tcp", "goooooooooooogle.com:80")
	if err != nil {
		fmt.Printf("%T\n", err) // HL
		log.Fatal(err)
	}
	defer conn.Close()
}

//END OMIT
