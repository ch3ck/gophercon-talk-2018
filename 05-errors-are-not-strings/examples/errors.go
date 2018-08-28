package main

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
)

//START OMIT
func connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		switch err := err.(type) {
		case *net.OpError:
			// return fmt.Errorf("failed to connect to %s: %v", err.Net, err)
			return errors.Wrapf(err, "failed to connect to %s", err.Net)
		default:
			// return fmt.Errorf("unkown error: %v", err)
			return errors.Wrap(err, "unknown error")
		}
	}

	defer conn.Close()

	return nil
}

//END OMIT

func main() {
	if err := connect("goooooooooooogle.com:80"); err != nil {
		fmt.Printf("%T\n", err)
		fmt.Printf("%T\n", errors.Cause(err)) //checking what possibly caused this error
	}
}
