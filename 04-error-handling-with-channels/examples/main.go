package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/profile" //PROFILE PACKAGE HELPS WITH PROFILING CODE
)

//START OMIT
func doSomething(name string) error {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	return fmt.Errorf("something went wrong with " + name)
}

func doSomethingTwice() error {
	// without the 2  goroutines this will leak a goroutine
	errc := make(chan error) // ISSUE OCCURS HERE
	//errc := make(chan error, 2) // FIX TO ISSUE
	go func() {
		defer fmt.Println("done wth a")
		errc <- doSomething("a")
	}()
	go func() {
		defer fmt.Println("done with b")
		errc <- doSomething("b")
	}()
	err := <-errc
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
