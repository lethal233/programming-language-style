package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Column struct {
	Data interface{}
	Fn   func() interface{}
}

type KV struct {
	Word string
	Freq int
}

var AllColumns []*Column

func updateColumns() {
	for _, col := range AllColumns {
		if col.Fn != nil {
			col.Data = col.Fn()
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a file path as argument.")
		os.Exit(1)
	}

	filePath := os.Args[1]
	stopWordsPath := "../stop_words.txt"

	contentBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read file: %s\n", filePath)
		os.Exit(1)
	}

	allWords := regexp.MustCompile(`[a-z]{2,}`).FindAllString(strings.ToLower(string(contentBytes)), -1)

	stopWordsFile, err := ioutil.ReadFile(stopWordsPath)
	if err != nil {
		fmt.Printf("Failed to read file: %s\n", stopWordsPath)
		os.Exit(1)
	}

	stopWords := strings.Split(string(stopWordsFile), ",")

	allWordsColumn := &Column{Data: allWords}
	stopWordsColumn := &Column{Data: stopWords}
	nonStopWordsColumn := &Column{Fn: func() interface{} {
		nonStopWords := make([]string, 0)
		d := allWordsColumn.Data.([]string)
		f := stopWordsColumn.Data.([]string)
		for _, word := range d {
			if !contains(f, word) {
				nonStopWords = append(nonStopWords, word)
			}
		}
		return nonStopWords
	}}
	uniqueWordsColumn := &Column{Fn: func() interface{} {
		wordSet := make(map[string]struct{})
		d := nonStopWordsColumn.Data.([]string)
		for _, word := range d {
			wordSet[word] = struct{}{}
		}
		uniqueWords := make([]string, 0, len(wordSet))
		for word := range wordSet {
			uniqueWords = append(uniqueWords, word)
		}
		return uniqueWords
	}}
	countsColumn := &Column{Fn: func() interface{} {
		d := nonStopWordsColumn.Data.([]string)
		f := uniqueWordsColumn.Data.([]string)
		counts := make([]KV, 0, len(d))
		for i, s := range f {
			counts = append(counts, KV{Word: s, Freq: 0})
			for _, word := range d {
				if s == word {
					counts[i].Freq++
				}
			}
		}
		return counts
	}}
	sortedDataColumn := &Column{Fn: func() interface{} {
		d := countsColumn.Data.([]KV)
		sort.Slice(d, func(i, j int) bool {
			return d[i].Freq > d[j].Freq
		})
		return d
	}}

	AllColumns = []*Column{allWordsColumn, stopWordsColumn, nonStopWordsColumn, uniqueWordsColumn, countsColumn, sortedDataColumn}

	updateColumns()

	ans := sortedDataColumn.Data.([]KV)
	for i := 0; i < 25 && i < len(ans); i++ {
		fmt.Println(ans[i].Word, " - ", ans[i].Freq)
	}
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
