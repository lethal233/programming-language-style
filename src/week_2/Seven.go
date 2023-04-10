package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	var da []rune
	for scanner.Scan() {
		da = append(da, []rune(strings.ToLower(scanner.Text()+" "))...)
	}
	file.Close()

	wds := strings.FieldsFunc(string(da), func(c rune) bool { return !unicode.IsLetter(c) && !unicode.IsNumber(c) })

	file, _ = os.Open("../stop_words.txt")
	scanner = bufio.NewScanner(file)
	stopWordsSet := map[string]bool{}
	scanner.Scan()
	for _, w := range strings.Split(scanner.Text(), ",") {
		stopWordsSet[w] = true
	}

	file.Close()

	var filteredWords []string
	for _, w := range wds {
		if !stopWordsSet[w] && len(w) > 1 {
			filteredWords = append(filteredWords, w)
		}
	}

	wordFreq := map[string]int{}
	for _, w := range filteredWords {
		wordFreq[w]++
	}

	sWds := make([]string, 0, len(wordFreq))
	for k := range wordFreq {
		sWds = append(sWds, k)
	}
	sort.Slice(sWds, func(i, j int) bool { return wordFreq[sWds[i]] > wordFreq[sWds[j]] })

	for i := 0; i < 25; i++ {
		fmt.Printf("%s - %d\n", sWds[i], wordFreq[sWds[i]])
	}
}
