all: extract links crawler

crawler: cmd/crawler/*.go crawl/*.go store/*.go
	go build -o crawler cmd/crawler/main.go   

extract: cmd/extract/*.go crawl/*.go
	go build -o extract cmd/extract/main.go   

links: cmd/links/*.go crawl/*.go
	go build -o links cmd/links/main.go  

clean: cmd/crawler/*.go crawl/*.go store/*.go
	rm -f extract links crawler