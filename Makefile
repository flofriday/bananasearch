all: extract links crawler server

crawler: cmd/crawler/*.go crawl/*.go store/*.go
	go build -o crawler cmd/crawler/main.go   

server: cmd/server/*.go store/*.go app/*.go 
	go build -o server cmd/server/main.go   

extract: cmd/extract/*.go crawl/*.go
	go build -o extract cmd/extract/main.go   

links: cmd/links/*.go crawl/*.go
	go build -o links cmd/links/main.go  

clean: cmd/crawler/*.go crawl/*.go store/*.go
	rm -f extract links crawler server