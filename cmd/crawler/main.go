package main

import (
	"log"
	"os"

	"github.com/flofriday/websearch/crawl"
	"github.com/flofriday/websearch/store"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("No seed url provided")
	}

	waitingQueue := os.Args[1:]
	visited := make(map[string]bool)
	index := store.NewIndex("index")
	limit := 10

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
			log.Printf("ERROR: %s", err.Error())
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

	// Statistics
	log.Print("")
	log.Print("--- Statistics ---")
	log.Printf("Indexed %d documents ", index.Docs())
	log.Printf("Indexed %d words", index.Words())
	log.Printf("%d documents in the waitingqueue", len(waitingQueue))

}
