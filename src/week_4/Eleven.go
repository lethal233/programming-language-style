package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

// TFExercise is the base interface for the classes
type TFExercise interface {
	Info() string
}

// DataStorageManager models the contents of the file
type DataStorageManager struct {
	data string
}

func NewDataStorageManager(filePath string) *DataStorageManager {
	data, _ := ioutil.ReadFile(filePath)
	dataStr := strings.ToLower(string(data))
	reg := regexp.MustCompile(`[\W_]+`)
	dataStr = reg.ReplaceAllString(dataStr, " ")
	return &DataStorageManager{data: dataStr}
}

func (d *DataStorageManager) Words() []string {
	return strings.Fields(d.data)
}

func (d *DataStorageManager) Info() string {
	return fmt.Sprintf("DataStorageManager: My major data structure is a %T", d.data)
}

// StopWordManager models the stop word filter
type StopWordManager struct {
	stopWords map[string]bool
}

func NewStopWordManager(filePath string) *StopWordManager {
	data, _ := ioutil.ReadFile(filePath)
	stopWordsList := strings.Split(string(data), ",")
	stopWords := make(map[string]bool)
	for _, word := range stopWordsList {
		stopWords[word] = true
	}
	for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
		stopWords[string(letter)] = true
	}
	return &StopWordManager{stopWords: stopWords}
}

func (s *StopWordManager) IsStopWord(word string) bool {
	return s.stopWords[word]
}

func (s *StopWordManager) Info() string {
	return fmt.Sprintf("StopWordManager: My major data structure is a %T", s.stopWords)
}

// WordFrequencyManager keeps the word frequency data
type WordFrequencyManager struct {
	wordFreqs map[string]int
}

func NewWordFrequencyManager() *WordFrequencyManager {
	return &WordFrequencyManager{wordFreqs: make(map[string]int)}
}

func (w *WordFrequencyManager) IncrementCount(word string) {
	w.wordFreqs[word]++
}

func (w *WordFrequencyManager) Sorted() []struct {
	Word  string
	Count int
} {
	wordFreqs := make([]struct {
		Word  string
		Count int
	}, 0, len(w.wordFreqs))

	for word, count := range w.wordFreqs {
		wordFreqs = append(wordFreqs, struct {
			Word  string
			Count int
		}{word, count})
	}

	sort.Slice(wordFreqs, func(i, j int) bool {
		return wordFreqs[i].Count > wordFreqs[j].Count
	})

	return wordFreqs
}

func (w *WordFrequencyManager) Info() string {
	return fmt.Sprintf("WordFrequencyManager: My major data structure is a %T", w.wordFreqs)
}

// WordFrequencyController controls the workflow
type WordFrequencyController struct {
	storageManager  *DataStorageManager
	stopWordManager *StopWordManager
	wordFreqManager *WordFrequencyManager
}

func NewWordFrequencyController(filePath string, stopWordsFilePath string) *WordFrequencyController {
	return &WordFrequencyController{
		storageManager:  NewDataStorageManager(filePath),
		stopWordManager: NewStopWordManager(stopWordsFilePath),
		wordFreqManager: NewWordFrequencyManager(),
	}
}

func (c *WordFrequencyController) Run() {
	for _, word := range c.storageManager.Words() {
		if !c.stopWordManager.IsStopWord(word) {
			c.wordFreqManager.IncrementCount(word)
		}
	}
	wordFreqs := c.wordFreqManager.Sorted()
	for _, wordFreq := range wordFreqs[0:25] {
		fmt.Printf("%s  -  %d\n", wordFreq.Word, wordFreq.Count)
	}
}

func main() {
	NewWordFrequencyController(os.Args[1], "../stop_words.txt").Run()
}
