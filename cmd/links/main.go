package main

import (
	"fmt"
	"os"

	"github.com/flofriday/websearch/crawl"
)

func main() {
	downloader := crawl.NewCachedDownloader(crawl.NewDefaultDownloader(), "cache")
	input, url, err := downloader.Load(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	allLinks := crawl.ExtractLinks(input, url)
	links := make(map[string]bool)
	for _, link := range allLinks {
		links[link] = true
	}

	i := 0
	for link, _ := range links {
		i++
		fmt.Printf("%4d) %s\n", i, link)
	}
}
