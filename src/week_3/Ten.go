package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type kv struct {
	Word string
	Freq int
}

type TFTheOne struct {
	Value interface{}
}

func (t *TFTheOne) bind(f func(interface{}) interface{}) *TFTheOne {
	t.Value = f(t.Value)
	return t
}

func (t *TFTheOne) printMe() {
	fmt.Print(t.Value)
}

func main() {
	tf := &TFTheOne{Value: os.Args[1]}
	tf.bind(readFile1).
		bind(filterChars1).
		bind(normalize1).
		bind(scan1).
		bind(removeStopWords1).
		bind(frequencies1).
		bind(sortFrequencies).
		bind(top25Freqs).
		printMe()
}

func readFile1(v interface{}) interface{} {
	var da []rune
	file, err := os.Open(v.(string))
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		da = append(da, []rune(scanner.Text())...)
		da = append(da, ' ')
	}
	return string(da)
}

func filterChars1(v interface{}) interface{} {
	strData := v.(string)
	reg, err := regexp.Compile("[\\W_]+")
	if err != nil {
		panic(err)
	}
	return reg.ReplaceAllString(strData, " ")
}

func normalize1(v interface{}) interface{} {
	strData := v.(string)
	return strings.ToLower(strData)
}

func scan1(v interface{}) interface{} {
	strData := v.(string)
	return strings.Fields(strData)
}

func removeStopWords1(v interface{}) interface{} {
	words := v.([]string)
	stopWordsSet := map[string]bool{}
	file, err := os.Open("../stop_words.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stopWords := strings.Split(scanner.Text(), ",")
		for i := 0; i < len(stopWords); i++ {
			stopWordsSet[stopWords[i]] = true
		}
	}
	for i := 'a'; i <= 'z'; i++ {
		stopWordsSet[string(i)] = true
	}
	var filteredWords []string
	for i := 0; i < len(words); i++ {
		if !stopWordsSet[words[i]] {
			filteredWords = append(filteredWords, words[i])
		}
	}
	return filteredWords
}

func frequencies1(v interface{}) interface{} {
	wordList := v.([]string)
	wf := make(map[string]int)
	for _, w := range wordList {
		wf[w]++
	}
	return wf
}

func sortFrequencies(v interface{}) interface{} {
	wordFreqs := v.(map[string]int)
	var sortedFreqs []kv
	for k, v := range wordFreqs {
		sortedFreqs = append(sortedFreqs, kv{k, v})
	}
	sort.Slice(sortedFreqs, func(i, j int) bool {
		return sortedFreqs[i].Freq > sortedFreqs[j].Freq
	})
	return sortedFreqs
}

func top25Freqs(v interface{}) interface{} {
	wordFreqs := v.([]kv)
	top25 := ""
	for i := 0; i < 25 && i < len(wordFreqs); i++ {
		top25 += fmt.Sprintf("%s - %d\n", wordFreqs[i].Word, wordFreqs[i].Freq)
	}
	return top25
}
