package store

import (
	"bufio"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"
)

type Index struct {
	path    string
	words   map[string][]int
	revDocs []string
}

func NewIndex(path string) *Index {
	return &Index{
		path:    path,
		words:   make(map[string][]int),
		revDocs: make([]string, 0, 64),
	}
}

func (i *Index) AddDoc(url string, text string) {
	id := len(i.revDocs)
	i.revDocs = append(i.revDocs, url)

	words := strings.Fields(text)
	for _, word := range words {
		word = strings.TrimFunc(word, unicode.IsPunct)
		if _, ok := i.words[word]; !ok {
			// The word does not yet exist so we need to create it
			i.words[word] = []int{id}
			continue
		}

		// The word already exists, lets see if this document is already in
		// there
		found := false
		for _, doc := range i.words[word] {
			if id == doc {
				found = true
				break
			}
		}

		if found {
			continue
		}

		// Add this document to the word
		i.words[word] = append(i.words[word], id)
	}
}

// TODO: more errorhandling
func (i *Index) Save() error {
	// Write the document index
	docPath := path.Join(i.path, "docs.txt")
	docFile, err := os.Create(docPath)
	if err != nil {
		return err
	}

	docWriter := bufio.NewWriter(docFile)
	for _, doc := range i.revDocs {
		docWriter.WriteString(doc)
		docWriter.WriteRune('\n')
	}
	docWriter.Flush()

	// Write the word index
	wordPath := path.Join(i.path, "words.txt")
	wordFile, err := os.Create(wordPath)
	if err != nil {
		return err
	}

	wordWriter := bufio.NewWriter(wordFile)
	for word, docs := range i.words {
		wordWriter.WriteString(word)
		wordWriter.WriteString(": ")
		for _, doc := range docs {
			wordWriter.WriteString(strconv.FormatInt(int64(doc), 10))
			wordWriter.WriteString(", ")
		}
		wordWriter.WriteRune('\n')
	}
	wordWriter.Flush()

	return nil
}
