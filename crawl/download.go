package crawl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

type Downloader interface {
	Load(url string) (string, string, error)
}

type DefaultDownloader struct {
	client http.Client
}

func NewDefaultDownloader() *DefaultDownloader {
	cl := http.Client{
		Timeout: time.Second * 10,
	}
	return &DefaultDownloader{client: cl}
}

func (d *DefaultDownloader) Load(url string) (string, string, error) {
	resp, err := d.client.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body), resp.Request.URL.String(), nil
}

type CachedDownloader struct {
	Downloader Downloader
	path       string
}

func NewCachedDownloader(downloader Downloader, cachePath string) *CachedDownloader {
	return &CachedDownloader{path: cachePath, Downloader: downloader}
}

func (d *CachedDownloader) Load(url string) (string, string, error) {
	cachePath := path.Join(d.path, strings.ReplaceAll(url, "/", ">"))

	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		body, url, err := d.Downloader.Load(url)
		if err != nil {
			return "", "", err
		}

		fmt.Println("Write cache at: %s", cachePath)
		ioutil.WriteFile(cachePath, []byte(body), 0644)
		return body, url, err
	}

	body, err := ioutil.ReadFile(cachePath)
	if err != nil {
		return "", "", err
	}
	return string(body), url, nil
}
