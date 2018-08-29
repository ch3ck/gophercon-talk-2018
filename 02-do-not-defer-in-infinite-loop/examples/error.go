package main

import (
	"log"
	"time"

	"github.com/pkg/profile"
)

//START OMIT
func loggingMonitorErr(files ...string) {
	for range time.Tick(time.Second) {
		for _, f := range files {
			//files coming in through the channel.
			fp := OpenFile(f)
			// The line below will never execute.
			defer fp.Close() // HL
			//process file
		}
	}
}

//END OMIT

type file string

func OpenFile(s string) file {
	log.Printf("opening %s", s) // HL
	return file(s)
}
func (f file) Close() { log.Printf("closing %s", f) } // HL

func main() {
	defer profile.Start(profile.TraceProfile).Stop() //ADD TRACING TOOL
	loggingMonitorErr("one.txt", "two.txt")          // HL
}
