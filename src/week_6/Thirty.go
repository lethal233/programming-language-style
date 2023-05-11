package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type wordFreq struct {
	word  string
	count int
}

func main() {
	filePath := os.Args[1]
	stopWordsFile := "../stop_words.txt"

	stopWordsData, _ := ioutil.ReadFile(stopWordsFile)
	stopWords := strings.Split(string(stopWordsData), ",")

	data, _ := ioutil.ReadFile(filePath)
	dataStr := strings.ToLower(string(data))

	wordRegexp := regexp.MustCompile(`[a-z]{2,}`)
	words := wordRegexp.FindAllString(dataStr, -1)

	wordSpace := make(chan string, len(words))
	freqSpace := make(chan map[string]int, 5)

	for _, word := range words {
		wordSpace <- word
	}
	close(wordSpace)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go processWords(wordSpace, freqSpace, stopWords, &wg)
	}

	wg.Wait()
	close(freqSpace)

	wordFreqs := make(map[string]int)
	for freq := range freqSpace {
		for k, v := range freq {
			wordFreqs[k] += v
		}
	}

	wordFreqSlice := make([]wordFreq, 0, len(wordFreqs))
	for k, v := range wordFreqs {
		wordFreqSlice = append(wordFreqSlice, wordFreq{k, v})
	}

	sort.Slice(wordFreqSlice, func(i, j int) bool {
		return wordFreqSlice[i].count > wordFreqSlice[j].count
	})

	for _, wf := range wordFreqSlice[:25] {
		fmt.Println(wf.word, "-", wf.count)
	}
}

func processWords(wordSpace <-chan string, freqSpace chan<- map[string]int, stopWords []string, wg *sync.WaitGroup) {
	defer wg.Done()

	wordFreqs := make(map[string]int)

	for word := range wordSpace {
		if !contains30(stopWords, word) {
			wordFreqs[word]++
		}
	}

	freqSpace <- wordFreqs
}

func contains30(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
