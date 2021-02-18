package main

import (
	"fmt"
	"os"

	"github.com/flofriday/websearch/crawl"
)

func main() {
	downloader := crawl.NewCachedDownloader(crawl.NewDefaultDownloader(), "cache")
	input, _, err := downloader.Load(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	os.Stdout.WriteString(crawl.ExtractText(input))
}
