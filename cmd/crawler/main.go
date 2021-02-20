package main

import (
	"log"

	"github.com/flofriday/websearch/crawl"
	"github.com/flofriday/websearch/store"
)

func main() {
	waitingQueue := []string{"https://www.orf.at"}
	visited := make(map[string]bool)
	index := store.NewIndex("index")
	limit := 100

	downloader := crawl.NewCachedDownloader(crawl.NewDefaultDownloader(), "cache")
	for len(visited) < limit && len(waitingQueue) > 0 {
		var url string
		url, waitingQueue = waitingQueue[0], waitingQueue[1:]
		if _, ok := visited[url]; ok {
			continue
		}
		log.Printf("Downloading: %s", url)
		body, url, err := downloader.Load(url)
		if err != nil {
			continue
		}
		if _, ok := visited[url]; ok {
			// URL already visited
			continue
		}
		visited[url] = true

		// Add the urls to the queue
		links := crawl.ExtractLinks(body, url)
		for _, link := range links {
			if _, ok := visited[link]; ok {
				// URL already visited
				continue
			}
			waitingQueue = append(waitingQueue, link)
		}

		// Add the words to the index
		index.AddDoc(url, crawl.ExtractText(body))
	}

	err := index.Save()
	if err != nil {
		log.Fatal(err.Error())
	}

	//log.Print(waitingQueue)

}
