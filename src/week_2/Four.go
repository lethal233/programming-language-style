package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

// no library functions
// no abstractions

func main() {
	stopWords, err := os.Open("../stop_words.txt")
	if err != nil {
		os.Exit(1)
	}
	defer stopWords.Close()
	stopWordsSet := map[string]bool{}
	stopWordsR := bufio.NewReader(stopWords)
	stopWordsLine, err := stopWordsR.ReadBytes('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Error in opening the file that stores stop words......", err)
		return
	}
	stopWordsArr := strings.Split(string(stopWordsLine), ",")
	for i := 0; i < len(stopWordsArr); i++ {
		stopWordsSet[stopWordsArr[i]] = true
	}

	if len(os.Args) < 2 {
		fmt.Println("Please provide the path to the text file to be processed")
		os.Exit(1)
	}
	path := os.Args[1]
	fp, err := os.Open(path)
	if err != nil {
		fmt.Println("Error in opening the text file......")
		return
	}
	defer fp.Close()
	r := bufio.NewReader(fp)
	dict := map[string]int32{}
	for {
		lineBytes, err := r.ReadBytes('\n') // "word--word"
		if err != nil && err != io.EOF {
			fmt.Println("Unexpected error")
			return
		}
		if err == io.EOF {
			break
		}
		line := strings.TrimSpace(string(lineBytes))
		startChar := -1

		for i, c := range line {
			if startChar == -1 {
				if unicode.IsLetter(c) || unicode.IsNumber(c) {
					startChar = i
				}
			} else {
				if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
					word := strings.ToLower(line[startChar:i])
					if stopWordsSet[word] || len(word) < 2 {
						startChar = -1
						continue
					}
					v, ok := dict[word]
					if ok {
						dict[word] = v + 1
					} else {
						dict[word] = 1
					}
					startChar = -1
				}
			}
		}
		if startChar != -1 {
			word := strings.ToLower(line[startChar:])
			if stopWordsSet[word] || len(word) < 2 {
				continue
			}
			v, ok := dict[word]
			if ok {
				dict[word] = v + 1

			} else {
				dict[word] = 1
			}
		}
	}

	sortedWords := make([]string, 0, len(dict))
	for k := range dict {
		sortedWords = append(sortedWords, k)
	}

	//selection sorting algorithm
	for i := 0; i < len(sortedWords)-1; i++ {
		maxIdx := i
		for j := i + 1; j < len(sortedWords); j++ {
			if dict[sortedWords[j]] > dict[sortedWords[maxIdx]] {
				maxIdx = j
			}
		}
		if maxIdx != i {
			sortedWords[maxIdx], sortedWords[i] = sortedWords[i], sortedWords[maxIdx]
		}
	}
	for _, word := range sortedWords[:25] {
		fmt.Printf("%s  -  %d\n", word, dict[word])
	}
}
