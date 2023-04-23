package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type WordFrequencyFramework struct {
	loadEventHandlers   []func(string)
	doworkEventHandlers []func()
	endEventHandlers    []func()
}

func (w *WordFrequencyFramework) registerForLoadEvent(handler func(string)) {
	w.loadEventHandlers = append(w.loadEventHandlers, handler)
}

func (w *WordFrequencyFramework) registerForDoworkEvent(handler func()) {
	w.doworkEventHandlers = append(w.doworkEventHandlers, handler)
}

func (w *WordFrequencyFramework) registerForEndEvent(handler func()) {
	w.endEventHandlers = append(w.endEventHandlers, handler)
}

func (w *WordFrequencyFramework) run(path2file string) {
	for _, h := range w.loadEventHandlers {
		h(path2file)
	}
	for _, h := range w.doworkEventHandlers {
		h()
	}
	for _, h := range w.endEventHandlers {
		h()
	}
}

type DataStorage struct {
	data              string
	stopWordFilter    *StopWordFilter
	wordEventHandlers []func(string)
}

func dataStorageConst(w *WordFrequencyFramework, s *StopWordFilter) *DataStorage {
	ds := &DataStorage{stopWordFilter: s, wordEventHandlers: []func(string){}}
	w.registerForLoadEvent(ds.load)
	w.registerForDoworkEvent(ds.produceWords)
	return ds
}

func (d *DataStorage) load(path2file string) {
	var da []rune
	file, err := os.Open(path2file)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		da = append(da, []rune(scanner.Text())...)
		da = append(da, ' ')
	}
	d.data = string(da)
	re := regexp.MustCompile(`[\W_]+`)
	d.data = strings.ToLower(re.ReplaceAllString(d.data, " "))
}

func (d *DataStorage) produceWords() {
	words := strings.Fields(d.data)
	for _, w := range words {
		if !d.stopWordFilter.isStopWord(w) {
			for _, handler := range d.wordEventHandlers {
				handler(w)
			}
		}
	}
}

func (d *DataStorage) registerForWordEvent(handler func(string)) {
	d.wordEventHandlers = append(d.wordEventHandlers, handler)
}

type StopWordFilter struct {
	stopWords []string
}

func stopWordFilterConst(w *WordFrequencyFramework) *StopWordFilter {
	swf := &StopWordFilter{
		stopWords: []string{},
	}
	w.registerForLoadEvent(swf.load)
	return swf
}

func (s *StopWordFilter) load(ignore string) {
	file, _ := os.Open("../stop_words.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s.stopWords = strings.Split(scanner.Text(), ",")

	for _, l := range "abcdefghijklmnopqrstuvwxyz" {
		s.stopWords = append(s.stopWords, string(l))
	}
}

func (s *StopWordFilter) isStopWord(word string) bool {
	for _, stopWord := range s.stopWords {
		if word == stopWord {
			return true
		}
	}
	return false
}

type WordFrequencyCounter struct {
	wordFreqs map[string]int32
}

func wordFrequencyCounterConst(w *WordFrequencyFramework, ds *DataStorage) *WordFrequencyCounter {
	wfc := &WordFrequencyCounter{
		wordFreqs: map[string]int32{},
	}
	ds.registerForWordEvent(wfc.incrementCount)
	w.registerForEndEvent(wfc.printFreqs)
	return wfc
}

func (wfc *WordFrequencyCounter) incrementCount(word string) {
	_, exist := wfc.wordFreqs[word]
	if exist {
		wfc.wordFreqs[word]++
	} else {
		wfc.wordFreqs[word] = 1
	}
}

func (wfc *WordFrequencyCounter) printFreqs() {
	sWds := make([]string, 0, len(wfc.wordFreqs))
	for k := range wfc.wordFreqs {
		sWds = append(sWds, k)
	}

	sort.Slice(sWds, func(i, j int) bool {
		return wfc.wordFreqs[sWds[i]] > wfc.wordFreqs[sWds[j]]
	})

	for _, w := range sWds[:25] {
		fmt.Printf("%s  -  %d\n", w, wfc.wordFreqs[w])
	}
}

func main() {
	wfapp := &WordFrequencyFramework{
		loadEventHandlers:   []func(string){},
		doworkEventHandlers: []func(){},
		endEventHandlers:    []func(){},
	}
	swf := stopWordFilterConst(wfapp)
	ds := dataStorageConst(wfapp, swf)
	wordFrequencyCounterConst(wfapp, ds)
	wfapp.run(os.Args[1])
}
