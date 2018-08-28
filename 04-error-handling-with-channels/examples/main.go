package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/profile"
)

func doSomething(name string) error {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	return fmt.Errorf("something went wrong with " + name)
}

func doSomethingTwice() error {
	// without the 2  goroutines this will leak a goroutine
	errc := make(chan error, 1)
	//errc := make(chan error, 2)
	go func() {
		defer fmt.Println("done with a")
		errc <- doSomething("a")
	}()
	go func() {
		defer fmt.Println("done with b")
		errc <- doSomething("b")
	}()
	err := <-errc
	return err
}

func main() {
	defer profile.Start().Stop()
	rand.Seed(time.Now().Unix())
	for range time.Tick(100 * time.Millisecond) {
		fmt.Println(doSomethingTwice())
	}

}
