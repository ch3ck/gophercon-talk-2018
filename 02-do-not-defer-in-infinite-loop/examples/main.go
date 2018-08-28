package main

import (
	"os"
	"time"
)

//START OMIT
func loggingMonitorErr(files []string) {
	for range time.Tick(time.Minute) {
		for range files {
			//files coming in through the channel.
			fp := os.Open("files")
			defer fp.Close() //<-- Does not execute -->
			//process file
		}
	}
}

//END OMIT

func main() {
	fp, err := os.Open("path/to/file.text")
	if err != nil {
		//handle error gracefully
	}
	defer fp.Close()
}

//START FIX
func loggingMonitorFix(files []string) {
	for range time.Tick(time.Minute) {
		for range files {
			//files coming in through the channel.
			func() {
				fp := os.Open("files")
				defer fp.Close() //<-- Executes since defer is bound to func context -->
				//process file
			}()
		}
	}
}

//END FIX
