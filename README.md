# websearch

![Screenshot](screenshot.png)
Here is a screenshot with the word "dog" as a query on an index that was built
on 100 pages starting from [Wikipedia - Linux](https://en.wikipedia.org/wiki/Linux)

This is just a simple project where I try to implement a web search engine like
Google.

At the moment you can crawl the web (inpolite) create an simple inverted index
and search for single words. However the results are not even ranked.

## How to use
First, you need to run the crawler to index the 100 websites (yes this is hardcoded at the moment).
Here we use the [Wikipedia - Linux](https://en.wikipedia.org/wiki/Linux) page as the seed but any Website
with enough links will do.
```bash
make all
./crawler https://en.wikipedia.org/wiki/Linux
```

Now, we can just start the server and open [http://localhost:8000](http://localhost:8000)
```bash
./server
```
