package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/profile" //PROFILE PACKAGE HELPS WITH PROFILING CODE
)

func doSomething(name string) error {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	return fmt.Errorf("something went wrong with " + name)
}

// START OMIT
func doSomethingTwice() error {
	// Issue occurs below
	errc := make(chan error, 2) // HL

	go func() {
		defer fmt.Println("done wth a")
		errc <- doSomething("a")
	}()
	go func() {
		defer fmt.Println("done with b")
		errc <- doSomething("b")
	}()
	err := <-errc // HL
	return err
}

//END OMIT

func main() {
	defer profile.Start(profile.TraceProfile).Stop() //ADD TRACING TOOL

	rand.Seed(time.Now().Unix())
	for range time.Tick(1000 * time.Millisecond) {
		fmt.Println("------")
		fmt.Println(doSomethingTwice())
	}

}
