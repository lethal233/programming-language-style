package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func extractWords21(pathToFile string) []string {
	if pathToFile == "" {
		return []string{}
	}

	content, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		fmt.Printf("I/O error when opening %s: %v\n", pathToFile, err)
		return []string{}
	}

	strData := string(content)
	pattern := regexp.MustCompile(`[\W_]+`)
	wordList := strings.Split(strings.ToLower(pattern.ReplaceAllString(strData, " ")), " ")
	return wordList
}

func removeStopWords21(wordList []string) []string {
	stopWordsFile := filepath.Join("..", "stop_words.txt")
	content, err := ioutil.ReadFile(stopWordsFile)
	if err != nil {
		fmt.Printf("I/O error when opening %s: %v\n", stopWordsFile, err)
		return wordList
	}

	stopWords := strings.Split(string(content), ",")
	for i := 'a'; i <= 'z'; i++ {
		stopWords = append(stopWords, string(i))
	}

	var filteredWords []string
	for _, word := range wordList {
		if !stringInSlice21(word, stopWords) {
			filteredWords = append(filteredWords, word)
		}
	}
	return filteredWords
}

func stringInSlice21(s string, slice []string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func frequencies21(wordList []string) map[string]int {
	if len(wordList) == 0 {
		return map[string]int{}
	}
	wordFreqs := make(map[string]int)
	for _, word := range wordList {
		wordFreqs[word]++
	}
	return wordFreqs
}

func sortWords21(wordFreq map[string]int) []wordCount21 {
	if len(wordFreq) == 0 {
		return []wordCount21{}
	}
	words := make([]wordCount21, 0, len(wordFreq))
	for word, count := range wordFreq {
		words = append(words, wordCount21{word, count})
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i].count > words[j].count
	})
	return words
}

type wordCount21 struct {
	word  string
	count int
}

func main() {
	var fileName string
	if len(os.Args) < 2 {
		fileName = "../pride-and-prejudice.txt" // here the reasonable (default) value is ../pride-and-prejudice.txt
	} else {
		fileName = os.Args[1]
	}
	wordFreqs := sortWords21(frequencies21(removeStopWords21(extractWords21(fileName))))

	for _, wc := range wordFreqs[:25] {
		fmt.Printf("%s  -  %d\n", wc.word, wc.count)
	}
}
