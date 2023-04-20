package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func extractWords(dataStorageObj map[string]interface{}, pathToFile string) {
	data, _ := ioutil.ReadFile(pathToFile)
	dataStr := strings.ToLower(string(data))
	reg := regexp.MustCompile(`[\W_]+`)
	dataStr = reg.ReplaceAllString(dataStr, " ")
	dataStorageObj["data"] = strings.Fields(dataStr)
}

func loadStopWords(stopWordsObj map[string]interface{}) {
	data, _ := ioutil.ReadFile("../stop_words.txt")
	stopWordsList := strings.Split(string(data), ",")
	stopWords := make(map[string]bool)
	for _, word := range stopWordsList {
		stopWords[word] = true
	}
	for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
		stopWords[string(letter)] = true
	}
	stopWordsObj["stop_words"] = stopWords
}

func incrementCount(wordFreqsObj map[string]interface{}, w string) {
	wordFreqs := wordFreqsObj["freqs"].(map[string]int)
	wordFreqs[w]++
}

var dataStorageObj = map[string]interface{}{
	"data":  []string{},
	"init":  nil,
	"words": nil,
}
var stopWordsObj = map[string]interface{}{
	"stop_words":   map[string]bool{},
	"init":         nil,
	"is_stop_word": nil,
}
var wordFreqsObj = map[string]interface{}{
	"freqs":           make(map[string]int),
	"increment_count": nil,
	"sorted":          nil,
}

func main() {
	dataStorageObj := map[string]interface{}{
		"data": []string{},
		"init": func(pathToFile string) {
			extractWords(dataStorageObj, pathToFile)
		},
		"words": func() []string {
			return dataStorageObj["data"].([]string)
		},
	}

	stopWordsObj := map[string]interface{}{
		"stop_words": map[string]bool{},
		"init": func() {
			loadStopWords(stopWordsObj)
		},
		"is_stop_word": func(word string) bool {
			return stopWordsObj["stop_words"].(map[string]bool)[word]
		},
	}

	wordFreqsObj := map[string]interface{}{
		"freqs":           make(map[string]int),
		"increment_count": func(w string) { incrementCount(wordFreqsObj, w) },
		"sorted": func() []struct {
			Word  string
			Count int
		} {
			wordFreqs := make([]struct {
				Word  string
				Count int
			}, 0, len(wordFreqsObj["freqs"].(map[string]int)))

			for word, count := range wordFreqsObj["freqs"].(map[string]int) {
				wordFreqs = append(wordFreqs, struct {
					Word  string
					Count int
				}{word, count})
			}

			sort.Slice(wordFreqs, func(i, j int) bool {
				return wordFreqs[i].Count > wordFreqs[j].Count
			})

			return wordFreqs
		},
	}

	dataStorageObj["init"].(func(string))(os.Args[1])
	stopWordsObj["init"].(func())()

	for _, w := range dataStorageObj["words"].(func() []string)() {
		if !stopWordsObj["is_stop_word"].(func(string) bool)(w) {
			wordFreqsObj["increment_count"].(func(string))(w)
		}
	}

	wordFreqs := wordFreqsObj["sorted"].(func() []struct {
		Word  string
		Count int
	})()
	for _, wordFreq := range wordFreqs[:25] {
		fmt.Println(wordFreq.Word, "-", wordFreq.Count)
	}
}
