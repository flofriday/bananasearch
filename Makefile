all: extract links

extract:
	go build -o extract cmd/extract/main.go   

links:
	go build -o links cmd/links/main.go  

clean: 
	rm -f extract links