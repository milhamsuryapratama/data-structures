package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Posting struct {
	DocumentID string
	Frequency  int
}

type PostingList map[string]*Posting

type InvertedIndex map[string]PostingList

func NewInvertedIndex(docs map[string]string) InvertedIndex {
	index := make(InvertedIndex)
	for id, doc := range docs {
		words := strings.Fields(doc)
		for _, word := range words {
			word = strings.ToLower(word)
			_, ok := index[word]
			if !ok {
				index[word] = make(PostingList)
			}
			_, ok = index[word][id]
			if !ok {
				index[word][id] = &Posting{DocumentID: id, Frequency: 1}
				continue
			}
			index[word][id].Frequency++
		}
	}

	ko, _ := json.MarshalIndent(index, "  ", "  ")
	fmt.Println(string(ko))

	return index
}

func (index InvertedIndex) Search(word string) PostingList {
	word = strings.ToLower(word)
	return index[word]
}

func Intersection(p1, p2 PostingList) PostingList {
	intersectionPL := make(PostingList)
	for id := range p1 {
		if p2[id] != nil {
			intersectionPL[id] = p1[id]
		}
	}
	return intersectionPL
}

func (index InvertedIndex) SearchTwoWords(word1, word2 string) PostingList {
	return Intersection(index.Search(word1), index.Search(word2))
}

func main() {
	// sample documents
	docs := map[string]string{
		"doc1": "The sky is blue.",
		"doc2": "The sun is bright today.",
		"doc3": "The sun is bright and the sky is beautiful and today.",
	}

	// initialise the inverted index with documents
	idx := NewInvertedIndex(docs)

	fmt.Printf("Index: \n%v\n\n", idx)

	// search single word
	postingList := idx.Search("sky")
	fmt.Printf("Search term: sky\nResults:\n")
	for documentId, posting := range postingList {
		fmt.Printf("Document ID: %s, Frequency: %d\n", documentId, posting.Frequency)
	}

	// search two words
	postingList = idx.SearchTwoWords("sun", "sky")
	fmt.Printf("\nSearch terms: sun and sky\nResults:\n")
	for documentId, posting := range postingList {
		fmt.Printf("Document ID: %s, Frequency: %d\n", documentId, posting.Frequency)
	}

}
