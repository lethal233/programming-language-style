package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
	"unicode"
)

// decomposed in functional abstractions
// pipelined

func readFileString(path2file string) []rune {
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
	return da
}

func filterCharsNormalize(data []rune) []rune {
	for i, c := range data {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			data[i] = ' '
		} else {
			data[i] = unicode.ToLower(c)
		}
	}
	return data
}

func scanWords(data []rune) []string {
	dataStr := string(data)
	var w []string
	w = append(w, strings.Fields(dataStr)...)
	return w
}

func removeSWords(words []string) []string {
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

func freq(words []string) map[string]int {
	wFreq := map[string]int{}
	for i := 0; i < len(words); i++ {
		wFreq[words[i]]++
	}
	return wFreq
}

func sortFreq(wFreq map[string]int) ([]string, map[string]int) {
	var sWords []string
	for k := range wFreq {
		sWords = append(sWords, k)
	}

	sort.Slice(sWords, func(i, j int) bool {
		return wFreq[sWords[i]] > wFreq[sWords[j]]
	})
	return sWords, wFreq
}

func printFreq(sWords []string, wFreq map[string]int) {
	for i := 0; i < 25; i++ {
		println(sWords[i], " - ", wFreq[sWords[i]])
	}
}

func main() {
	printFreq(sortFreq(freq(removeSWords(scanWords(filterCharsNormalize(readFileString(os.Args[1])))))))
}
