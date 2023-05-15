package main

//
//import (
//	"bufio"
//	"fmt"
//	"io/ioutil"
//	"os"
//	"sort"
//	"strings"
//	"unicode"
//)
//
//func characters(filename string, ch chan rune) {
//	file, _ := os.Open(filename)
//	defer file.Close()
//	reader := bufio.NewReader(file)
//	for {
//		char, _, err := reader.ReadRune()
//		if err != nil {
//			close(ch)
//			return
//		}
//		ch <- char
//	}
//}
//
//func allWords(filename string, ch chan string) {
//	charCh := make(chan rune)
//	go characters(filename, charCh)
//	word := ""
//	for char := range charCh {
//		if unicode.IsLetter(char) {
//			word += strings.ToLower(string(char))
//		} else if word != "" {
//			ch <- word
//			word = ""
//		}
//	}
//	if word != "" {
//		ch <- word
//	}
//	close(ch)
//}
//
//func nonStopWords(filename string, ch chan string) {
//	wordCh := make(chan string)
//	go allWords(filename, wordCh)
//	stopWordsBytes, _ := ioutil.ReadFile("../stop_words.txt")
//	stopWords := strings.Split(string(stopWordsBytes), ",")
//	stopWordsMap := make(map[string]struct{})
//	for _, word := range stopWords {
//		stopWordsMap[word] = struct{}{}
//	}
//	for i := 'a'; i < 'z'; i++ {
//		stopWordsMap[string(i)] = struct{}{}
//	}
//	for word := range wordCh {
//		if _, ok := stopWordsMap[word]; !ok {
//			ch <- word
//		}
//	}
//	close(ch)
//}
//
//func countAndSort(filename string, ch chan []struct {
//	word string
//	freq int
//}) {
//	freqMap := make(map[string]int)
//	nonStopWordsCh := make(chan string)
//	go nonStopWords(filename, nonStopWordsCh)
//	i := 1
//	for word := range nonStopWordsCh {
//		freqMap[word]++
//		if i%5000 == 0 {
//			ch <- getSortedWordFreqs(freqMap)
//		}
//		i++
//	}
//	ch <- getSortedWordFreqs(freqMap)
//	close(ch)
//}
//
//func getSortedWordFreqs(freqMap map[string]int) []struct {
//	word string
//	freq int
//} {
//	wordFreqs := make([]struct {
//		word string
//		freq int
//	}, len(freqMap))
//	i := 0
//	for word, freq := range freqMap {
//		wordFreqs[i] = struct {
//			word string
//			freq int
//		}{word, freq}
//		i++
//	}
//	sort.Slice(wordFreqs, func(i, j int) bool {
//		return wordFreqs[i].freq > wordFreqs[j].freq
//	})
//	return wordFreqs
//}
//
//func main() {
//	wordFreqsCh := make(chan []struct {
//		word string
//		freq int
//	})
//	go countAndSort(os.Args[1], wordFreqsCh)
//	for wordFreqs := range wordFreqsCh {
//		fmt.Println("-----------------------------")
//		for _, wordFreq := range wordFreqs[:25] {
//			fmt.Printf("%s - %d\n", wordFreq.word, wordFreq.freq)
//		}
//	}
//}
