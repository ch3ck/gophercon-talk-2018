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
		switch err := err.(type) {
		case *net.OpError:
			fmt.Println(err)
			fmt.Printf("failed to connect to %s because %v\n", err.Net, err.Err)
			return
		default:
			log.Fatal(err)
		}
	}
	defer conn.Close()
}

//END OMIT
