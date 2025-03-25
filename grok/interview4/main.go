package main

import (
	"fmt"
	"net/http"
	"time"
)

type httpResp struct {
	resp *http.Response
	err  error
}

func fetch(sites []string) map[string]bool {
	goodSites := map[string]bool{}
	out := make(chan httpResp)
	for _, v := range sites {
		site := "https://" + v
		go func() {
			resp, err := http.Get(site)
			out <- httpResp{resp: resp, err: err}
		}()
		r := <-out
		if r.err != nil {
			fmt.Println(r.err)
		}
		if r.resp.StatusCode == 200 {
			goodSites[v] = true
		}
	}
	return goodSites
}

func main() {
	now := time.Now()
	sites := []string{"google.com", "yahoo.com", "microsoft.com", "github.com", "amazon.com"}
	results := fetch(sites)
	for k, v := range results {
		fmt.Printf("%s = %v\n", k, v)
	}
	fmt.Println(time.Since(now))
}
