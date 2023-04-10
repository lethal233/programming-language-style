package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
	"unicode"
)

var data []rune
var words []string
var wordFrequencies = map[string]int{}
var sortedWords []string

func readFile(path2file string) {
	file, err := os.Open(path2file)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, []rune(scanner.Text())...)
		data = append(data, ' ')
	}
}

func filterCharsAndNormalize() {
	for i, c := range data {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			data[i] = ' '
		} else {
			data[i] = unicode.ToLower(c)
		}
	}
}

func scan() {
	dataStr := string(data)
	words = append(words, strings.Fields(dataStr)...)
}

func removeStopWords() {
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
	for i := 0; i < len(words); i++ {
		if stopWordsSet[words[i]] {
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
}

func frequencies() {
	for w := range words {
		if v, ok := wordFrequencies[words[w]]; ok {
			wordFrequencies[words[w]] = v + 1
		} else {
			wordFrequencies[words[w]] = 1
		}
	}
}

func sort2() {
	sortedWords = make([]string, 0, len(wordFrequencies))
	for k := range wordFrequencies {
		sortedWords = append(sortedWords, k)
	}

	sort.Slice(sortedWords, func(i, j int) bool {
		return wordFrequencies[sortedWords[i]] > wordFrequencies[sortedWords[j]]
	})
}

func main() {
	readFile(os.Args[1])
	filterCharsAndNormalize()
	scan()
	removeStopWords()
	frequencies()
	sort2()
	for i := 0; i < 25; i++ {
		println(sortedWords[i], " - ", wordFrequencies[sortedWords[i]])
	}
}
