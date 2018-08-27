package main

import (
	"os"
	"time"
)

func loggingMonitorErr(files []string) {
	for range time.Tick(time.Minute) {
		for range files {
			//files coming in through the channel.
			fp := os.Open("files")
			defer fp.Close()
			//process file
		}
	}
}

func main() {
	fp, err := os.Open("path/to/file.text")
	if err != nil {
		//handle error gracefully
	}
	defer fp.Close()
}

func loggingMonitorFix(files []string) {
	for range time.Tick(time.Minute) {
		for range files {
			//files coming in through the channel.
			func() {
				fp := os.Open("files")
				defer fp.Close()
				//process file
			}()
		}
	}
}
