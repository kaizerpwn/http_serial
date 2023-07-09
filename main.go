package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSavePage(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is not responding, it is probably down.\n", url)
	} else {
		defer resp.Body.Close()

		fmt.Printf("%s is responding, status code : %d\n", url, resp.StatusCode)

		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			file := strings.Split(url, "//")[1]
			file += ".txt"
			fmt.Printf("Saving page to %s\n", file)

			//  Saving page to file
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	wg.Done()
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.youtube.com", "https://facebook889989.com"}

	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSavePage(url, &wg)
	}

	fmt.Println("Number of Go Routines:", runtime.NumGoroutine())

	wg.Wait()
}
