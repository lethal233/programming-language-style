package main

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strings"
)

type nop func()
type pt func([]string, map[string]int, nop)
type fq func([]string, sf)
type rsw func([]string, fq)
type sc func(string, rsw)
type n func(string, sc)
type fc func([]rune, n)
type sf func(map[string]int, pt)

func readFile(path2File string, cb fc) {
	var da []rune
	file, err := os.Open(path2File)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		da = append(da, []rune(scanner.Text())...)
		da = append(da, ' ')
	}
	cb(da, normalize) // filerChars(da, normalize)
}

func filterChars(data []rune, cb n) {
	pattern, err := regexp.Compile("[\\W_]+")
	if err != nil {
		os.Exit(1)
	}
	dataStr := string(data)
	dataStr = pattern.ReplaceAllString(dataStr, " ")
	cb(dataStr, scan) // normalize(dataStr, scan)
}

func normalize(data string, cb sc) {
	data = strings.ToLower(data)
	cb(data, removeStopWords) // scan(data, removeStopWords)
}

func scan(data string, cb rsw) {
	var words []string
	words = append(words, strings.Fields(data)...)
	cb(words, frequencies) // removeStopWords(words, frequencies)
}

func removeStopWords(words []string, cb fq) {
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
	cb(filteredWords, sortFreq) // frequencies(filteredWords, sortFreq)
}

func frequencies(words []string, cb sf) {
	var wordFreq = map[string]int{}
	for _, w := range words {
		wordFreq[w]++
	}
	cb(wordFreq, printText) // sortFreq(wordFreq, printText)
}

func sortFreq(wordFreq map[string]int, cb pt) {
	var sWords []string
	for k := range wordFreq {
		sWords = append(sWords, k)
	}

	sort.Slice(sWords, func(i, j int) bool {
		return wordFreq[sWords[i]] > wordFreq[sWords[j]]
	})
	cb(sWords, wordFreq, noOp) // printText(sWords, wordFreq, noOp)
}

func printText(sWords []string, wFreq map[string]int, cb nop) {
	for i := 0; i < 25; i++ {
		println(sWords[i], " - ", wFreq[sWords[i]])
	}
	cb()
}

func noOp() {
	return
}

func main() {
	readFile(os.Args[1], filterChars)
}
