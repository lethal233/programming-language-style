package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	filePath := os.Args[1]
	stopWordsFile := "../stop_words.txt"

	stopWordsData, _ := ioutil.ReadFile(stopWordsFile)
	stopWords := strings.Split(string(stopWordsData), ",")

	data, _ := ioutil.ReadFile(filePath)
	dataStr := string(data)

	splits := mapSplitWords(partition(dataStr, 200), stopWords)
	splitsPerWord := regroup(splits)
	wordFreqs := sortWordFreqs(mapCountWords(splitsPerWord))

	for _, wf := range wordFreqs[:25] {
		fmt.Println(wf.word, "-", wf.count)
	}
}

type wordFreq32 struct {
	word  string
	count int
}

// generator like yield
func partition(dataStr string, nlines int) chan string {
	lines := strings.Split(dataStr, "\n")
	ch := make(chan string)
	go func() {
		for i := 0; i < len(lines); i += nlines {
			ch <- strings.Join(lines[i:min(i+nlines, len(lines))], "\n")
		}
		close(ch)
	}()
	return ch
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func mapSplitWords(partitions chan string, stopWords []string) [][]wordFreq32 {
	var result [][]wordFreq32
	for p := range partitions {
		if p != "" {
			result = append(result, splitWords(p, stopWords))
		} else {
			break
		}
	}
	return result
}

func splitWords(dataStr string, stopWords []string) []wordFreq32 {
	wordRegexp := regexp.MustCompile(`[a-z]{2,}`)
	words := wordRegexp.FindAllString(strings.ToLower(dataStr), -1)
	var wordFreqs []wordFreq32
	for _, word := range words {
		if !contains32(stopWords, word) {
			wordFreqs = append(wordFreqs, wordFreq32{word, 1})
		}
	}
	return wordFreqs
}

func regroup(pairsList [][]wordFreq32) map[string][]wordFreq32 {
	mapping := make(map[string][]wordFreq32)
	for _, pairs := range pairsList {
		for _, p := range pairs {
			mapping[p.word] = append(mapping[p.word], p)
		}
	}
	return mapping
}

func mapCountWords(splitsPerWord map[string][]wordFreq32) []wordFreq32 {
	var wordFreqs []wordFreq32
	for _, v := range splitsPerWord {
		wordFreqs = append(wordFreqs, countWords(v))
	}
	return wordFreqs
}

func countWords(mapping []wordFreq32) wordFreq32 {
	count := 0
	for _, pair := range mapping {
		count += pair.count
	}
	return wordFreq32{mapping[0].word, count}
}

func sortWordFreqs(wordFreqs []wordFreq32) []wordFreq32 {
	sort.Slice(wordFreqs, func(i, j int) bool {
		return wordFreqs[i].count > wordFreqs[j].count
	})
	return wordFreqs
}

func contains32(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
