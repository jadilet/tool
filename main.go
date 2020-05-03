package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GetMd5URL get md5 hash of http request response
func GetMd5URL(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	digest := md5.New()
	if _, err := io.Copy(digest, resp.Body); err != nil {
		return "", err
	}

	return hex.EncodeToString(digest.Sum(nil)), nil
}

func main() {
	parallel := flag.Int("parallel", 10, "Limit the number of parallel requests")
	flag.Parse()

	// Limit the concurrent goroutines
	concurrentGoroutines := make(chan struct{}, *parallel)

	for i := 0; i < *parallel; i++ {
		concurrentGoroutines <- struct{}{}
	}

	done := make(chan bool)
	waitForAllJobs := make(chan bool)

	go func() {
		for i := 0; i < len(flag.Args()); i++ {
			<-done

			concurrentGoroutines <- struct{}{}
		}

		waitForAllJobs <- true
	}()

	for index := range flag.Args() {
		<-concurrentGoroutines

		if u, err := url.Parse(flag.Arg(index)); err == nil {
			u.Scheme = "http"

			go func(url string) {
				digest, err := GetMd5URL(url)
	
				if err != nil {
					fmt.Printf("%s", err.Error())
				}
	
				fmt.Println(url, digest)
				done <- true
			}(u.String())
		}
	}

	<-waitForAllJobs
}
