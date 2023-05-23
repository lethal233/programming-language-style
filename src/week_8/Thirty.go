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

var (
	wordChan  = make(chan string, 2<<16)
	freqChan  = make(chan map[string]int, 10)
	wg        sync.WaitGroup
	stopwords map[string]struct{}
)

func processWords() {
	defer wg.Done()
	wordFreqs := make(map[string]int)
	for i := 0; ; i++ {
		word := <-wordChan
		if word == "" {
			break
		}
		if _, ok := stopwords[word]; !ok {
			wordFreqs[word]++
		}
	}
	freqChan <- wordFreqs
}

func main() {
	stopwords = make(map[string]struct{})
	fileBytes, _ := ioutil.ReadFile("../stop_words.txt")
	for _, word := range strings.Split(string(fileBytes), ",") {
		stopwords[word] = struct{}{}
	}
	for i := 'a'; i < 'z'; i++ {
		stopwords[string(i)] = struct{}{}
	}

	fileBytes, _ = ioutil.ReadFile(os.Args[1])
	wordFreqs1 := make(map[string]int)
	words := regexp.MustCompile(`[a-z]{2,}`).FindAllString(strings.ToLower(string(fileBytes)), -1)
	for _, word := range words {
		wordFreqs1[word] += 1
		wordChan <- word
	}
	close(wordChan)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go processWords()
	}
	wg.Wait()
	close(freqChan)

	wordFreqs := make(map[string]int)
	size := len(freqChan)
	for i := 0; i < size; i++ {
		fm := <-freqChan
		if fm == nil {
			break
		}
		for word, freq := range fm {
			wordFreqs[word] += freq
		}
	}

	type Pair struct {
		Word  string
		Count int
	}
	pairs := make([]Pair, 0, len(wordFreqs))
	for word, count := range wordFreqs {
		pairs = append(pairs, Pair{word, count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	for _, pair := range pairs[:25] {
		fmt.Println(pair.Word, "-", pair.Count)
	}

}
