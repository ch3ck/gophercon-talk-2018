package main

import (
	"fmt"
	"log"
	"net/http"
)

func openURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("Error opening url %+v", err)
		//Write error to stderr
		return err //exit with failure
	}
	defer resp.Body.Close()
	return nil
}

func main() {
	err := openURL("https://gooooogol.com")
	if err != nil {
		log.Fatal(err)
	}
}
