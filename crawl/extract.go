package crawl

import (
	"html"
	"net/url"
	"regexp"
	"strings"
)

func ExtractText(text string) string {
	replaceScriptStyle := regexp.MustCompile("(<style[^<>]*>.*?<\\/style>|<script[^<>]*>.*?<\\/script>)")
	replaceTags := regexp.MustCompile("<.*?>")
	replaceMultipleSpace := regexp.MustCompile("\\s\\s+")

	text = strings.ReplaceAll(text, "\n", " ")
	text = replaceScriptStyle.ReplaceAllString(text, "")
	text = replaceTags.ReplaceAllString(text, " ")
	text = replaceMultipleSpace.ReplaceAllString(text, " ")
	text = html.UnescapeString(text)
	text = strings.ToLower(text)

	return text
}

func ExtractLinks(text string, base string) []string {
	findTags := regexp.MustCompile("<a[^>]* href=\".*?\"[^>]*>")
	findHref := regexp.MustCompile("href=\".*?\"")
	baseurl, _ := url.Parse(base)

	// TODO: parse the url to verify that the are valid

	// TODO: remove all hashbangs from urls
	// TODO: remove duplicate links
	tags := findTags.FindAllString(text, -1)
	links := make([]string, 0, len(tags))
	for _, tag := range tags {
		link := findHref.FindString(tag)
		link = link[6 : len(link)-1]
		linkurl, err := url.Parse(link)
		if err != nil {
			continue
		}
		links = append(links, baseurl.ResolveReference(linkurl).String())
	}
	return links
}
