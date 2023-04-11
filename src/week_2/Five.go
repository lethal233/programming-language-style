package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

var data []rune
var words []string
var wdDict []string
var frqDict []int32

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
	var stopWordsList []string
	file, err := os.Open("../stop_words.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stopWordsList = strings.Split(scanner.Text(), ",")
	}
	for i := 0; i < len(words); i++ {
		//println(words[i])
		flag := false
		if len(words[i]) < 2 {
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}
		for _, stopWord := range stopWordsList {
			if words[i] == stopWord {
				flag = true
				break
			}
		}
		if flag {
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
}

func frequencies() {
	for w := range words {
		flg := false
		for ind, wd := range wdDict {
			if wd == words[w] {
				frqDict[ind]++
				flg = true
				break
			}
		}
		if !flg {
			wdDict = append(wdDict, words[w])
			frqDict = append(frqDict, 1)
		}
	}
}

func sort2() {
	for i := 0; i < len(wdDict)-1; i++ {
		maxIdx := i
		for j := i + 1; j < len(frqDict); j++ {
			if frqDict[j] > frqDict[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx != i {
			wdDict[maxIdx], wdDict[i] = wdDict[i], wdDict[maxIdx]
			frqDict[maxIdx], frqDict[i] = frqDict[i], frqDict[maxIdx]
		}
	}
}

func main() {
	readFile(os.Args[1])
	filterCharsAndNormalize()
	scan()
	removeStopWords()
	frequencies()
	sort2()
	for i := 0; i < 25; i++ {
		println(wdDict[i], " - ", frqDict[i])
	}
}
