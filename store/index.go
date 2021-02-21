package store

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"strings"
	"unicode"
)

type Index struct {
	Path  string `json:",omitempty"`
	Words map[string][]int
	Docs  []string
}

func NewIndex(path string) *Index {
	return &Index{
		Path:  path,
		Words: make(map[string][]int),
		Docs:  make([]string, 0, 64),
	}
}

func LoadIndex(storePath string) (*Index, error) {
	indexPath := path.Join(storePath, "index.json")
	data, err := ioutil.ReadFile(indexPath)
	if err != nil {
		return nil, err
	}

	index := Index{}
	err = json.Unmarshal(data, &index)
	if err != nil {
		return nil, err
	}

	index.Path = storePath
	return &index, nil
}

func (i *Index) AddDoc(url string, text string) {
	id := len(i.Docs)
	i.Docs = append(i.Docs, url)

	words := strings.Fields(text)
	for _, word := range words {
		word = strings.TrimFunc(word, unicode.IsPunct)
		if _, ok := i.Words[word]; !ok {
			// The word does not yet exist so we need to create it
			i.Words[word] = []int{id}
			continue
		}

		// The word already exists, lets see if this document is already in
		// there
		found := false
		for _, doc := range i.Words[word] {
			if id == doc {
				found = true
				break
			}
		}

		if found {
			continue
		}

		// Add this document to the word
		i.Words[word] = append(i.Words[word], id)
	}
}

// TODO: more errorhandling
func (i *Index) Save() error {
	data, err := json.Marshal(i)
	if err != nil {
		return err
	}

	indexPath := path.Join(i.Path, "index.json")
	err = ioutil.WriteFile(indexPath, data, 0666)
	return err
}

func (i *Index) GetDocs(query string) []string {
	ids, ok := i.Words[query]
	if !ok {
		return []string{}
	}
	docs := make([]string, 0, len(ids))
	for _, id := range ids {
		docs = append(docs, i.Docs[id])
	}
	return docs
}

func (i *Index) NumDocs() int {
	return len(i.Docs)
}

func (i *Index) NumWords() int {
	return len(i.Words)
}
