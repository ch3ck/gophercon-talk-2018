package main

import (
	"log"
	"time"

	"github.com/pkg/profile"
)

//START OMIT
type file string

func OpenFile(s string) file {
	log.Printf("opening %s", s) // HL
	return file(s)
}
func (f file) Close() { log.Printf("closing %s", f) } // HL

func loggingMonitorFix(files ...string) {
	for range time.Tick(time.Second) {
		for _, f := range files {
			//files coming in through the channel.
			func() { // HL
				fp := OpenFile(f) // HL
				defer fp.Close()  // HL
				//process file
			}() // HL
		}
	}
}

//END OMIT

func main() {
	defer profile.Start(profile.TraceProfile).Stop() //ADD TRACING TOOL

	loggingMonitorFix("one.txt", "two.txt")
}
