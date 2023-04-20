package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

// DataStorageManager12 models the contents of the file
type DataStorageManager12 struct {
	data string
}

func NewDataStorageManager12() *DataStorageManager12 {
	return &DataStorageManager12{}
}

func (d *DataStorageManager12) Dispatch(message []interface{}) interface{} {
	switch message[0] {
	case "init":
		return d.init(message[1].(string))
	case "words":
		return d.words()
	default:
		panic("Message not understood: " + message[0].(string))
	}
}

func (d *DataStorageManager12) init(filePath string) interface{} {
	data, _ := ioutil.ReadFile(filePath)
	dataStr := strings.ToLower(string(data))
	reg := regexp.MustCompile(`[\W_]+`)
	dataStr = reg.ReplaceAllString(dataStr, " ")
	d.data = dataStr
	return d.data
}

func (d *DataStorageManager12) words() []string {
	return strings.Fields(d.data)
}

// StopWordManager12 models the stop word filter
type StopWordManager12 struct {
	stopWords map[string]bool
}

func NewStopWordManager12() *StopWordManager12 {
	return &StopWordManager12{}
}

func (s *StopWordManager12) Dispatch(message []interface{}) interface{} {
	switch message[0] {
	case "init":
		return s.init()
	case "is_stop_word":
		return s.isStopWord(message[1].(string))
	default:
		panic("Message not understood: " + message[0].(string))
	}
}

func (s *StopWordManager12) init() interface{} {
	data, _ := ioutil.ReadFile("../stop_words.txt")
	stopWordsList := strings.Split(string(data), ",")
	stopWords := make(map[string]bool)
	for _, word := range stopWordsList {
		stopWords[word] = true
	}
	for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
		stopWords[string(letter)] = true
	}
	s.stopWords = stopWords
	return s.stopWords
}

func (s *StopWordManager12) isStopWord(word string) bool {
	return s.stopWords[word]
}

// WordFrequencyManager12 keeps the word frequency data
type WordFrequencyManager12 struct {
	wordFreqs map[string]int
}

func NewWordFrequencyManager12() *WordFrequencyManager12 {
	return &WordFrequencyManager12{wordFreqs: make(map[string]int)}
}

func (w *WordFrequencyManager12) Dispatch(message []interface{}) interface{} {
	switch message[0] {
	case "increment_count":
		return w.incrementCount(message[1].(string))
	case "sorted":
		return w.sorted()
	default:
		panic("Message not understood: " + message[0].(string))
	}
}

func (w *WordFrequencyManager12) incrementCount(word string) interface{} {
	w.wordFreqs[word]++
	return w.wordFreqs
}

func (w *WordFrequencyManager12) sorted() []struct {
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

// WordFrequencyController12 controls the workflow
type WordFrequencyController12 struct {
	storageManager  *DataStorageManager12
	stopWordManager *StopWordManager12
	wordFreqManager *WordFrequencyManager12
}

func NewWordFrequencyController12() *WordFrequencyController12 {
	return &WordFrequencyController12{}
}

func (w *WordFrequencyController12) Dispatch(message []interface{}) interface{} {
	switch message[0] {
	case "init":
		return w.init(message[1].(string))
	case "run":
		return w.run()
	default:
		panic("Message not understood: " + message[0].(string))
	}
}

func (w *WordFrequencyController12) init(filePath string) interface{} {
	w.storageManager = NewDataStorageManager12()
	w.stopWordManager = NewStopWordManager12()
	w.wordFreqManager = NewWordFrequencyManager12()
	w.storageManager.Dispatch([]interface{}{"init", filePath})
	w.stopWordManager.Dispatch([]interface{}{"init"})
	return w
}

func (w *WordFrequencyController12) run() interface{} {
	for _, word := range w.storageManager.Dispatch([]interface{}{"words"}).([]string) {
		if !w.stopWordManager.Dispatch([]interface{}{"is_stop_word", word}).(bool) {
			w.wordFreqManager.Dispatch([]interface{}{"increment_count", word})
		}
	}
	wordFreqs := w.wordFreqManager.Dispatch([]interface{}{"sorted"}).([]struct {
		Word  string
		Count int
	})
	for _, wf := range wordFreqs[:25] {
		println(wf.Word, " - ", wf.Count)
	}
	return w
}

func main() {
	wfcontroller := NewWordFrequencyController12()
	wfcontroller.Dispatch([]interface{}{"init", os.Args[1]})
	wfcontroller.Dispatch([]interface{}{"run"})
}
