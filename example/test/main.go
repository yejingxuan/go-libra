package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	var eg errgroup.Group
	apps := make([]string, 0)
	apps = append(apps, "21223", "2234")
	fmt.Println(apps)
	for _, item := range apps {
		item := item
		eg.Go(func() (err error) {
			fmt.Println(item)
			return
		})
	}
	eg.Wait()
	//ExampleGroup_justErrors()
}



func ExampleGroup_justErrors() {
	var g errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			/*resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}*/
			fmt.Println(url)
			return nil
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}