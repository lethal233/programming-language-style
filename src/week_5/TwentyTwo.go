package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var myLog = log.New(os.Stderr, "\033[0;31m[Error]\033[0m ", log.LstdFlags|log.Llongfile)

func extractWords22(pathToFile string) ([]string, error) {
	if pathToFile == "" {
		myLog.Fatalf("I need a non-empty string!\n")
	}

	content, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		fmt.Printf("I/O error when opening %s: %v! I quit!\n", pathToFile, err)
		return nil, err
	}

	strData := string(content)
	pattern := regexp.MustCompile(`[\W_]+`)
	wordList := strings.Split(strings.ToLower(pattern.ReplaceAllString(strData, " ")), " ")
	return wordList, nil
}

func removeStopWords22(wordList []string, e error) ([]string, error) {
	if e != nil {
		return nil, e
	}
	stopWordsFile := filepath.Join("..", "stop_words.txt")
	content, err := ioutil.ReadFile(stopWordsFile)
	if err != nil {
		fmt.Printf("I/O error when opening %s: %v! I quit!\n", stopWordsFile, err)
		return nil, err
	}

	stopWords := strings.Split(string(content), ",")
	for i := 'a'; i <= 'z'; i++ {
		stopWords = append(stopWords, string(i))
	}

	var filteredWords []string
	for _, word := range wordList {
		if !stringInSlice22(word, stopWords) {
			filteredWords = append(filteredWords, word)
		}
	}
	return filteredWords, nil
}

func stringInSlice22(s string, slice []string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func frequencies22(wordList []string, e error) (map[string]int, error) {
	if e != nil {
		return nil, e
	}
	wordFreqs := make(map[string]int)
	for _, word := range wordList {
		wordFreqs[word]++
	}
	return wordFreqs, nil
}

func sortWords(wordFreq map[string]int, e error) ([]wordCount, error) {
	if e != nil {
		return nil, e
	}
	words := make([]wordCount, 0, len(wordFreq))
	for word, count := range wordFreq {
		words = append(words, wordCount{word, count})
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i].count > words[j].count
	})
	return words, nil
}

type wordCount struct {
	word  string
	count int
}

func main() {
	if len(os.Args) < 2 {
		myLog.Fatalf("You idiot! I need an input file!\n")
	}
	pathToFile := os.Args[1]
	wordFreqs, err := sortWords(frequencies22(removeStopWords22(extractWords22(pathToFile))))
	if err != nil {
		myLog.Fatalf("Something Wrong: %v\n", err)
	}

	if len(wordFreqs) < 25 {
		myLog.Fatalf("SRSLY? Less than 25 words!\n")
	}

	for _, wc := range wordFreqs[:25] {
		fmt.Printf("%s  -  %d\n", wc.word, wc.count)
	}
}
