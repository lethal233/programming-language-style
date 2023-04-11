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
		fmt.Println("Error in opening the file that stores stop words......", err)
		os.Exit(1)
	}
	defer stopWords.Close()
	stopWordsR := bufio.NewReader(stopWords)
	stopWordsLine, err := stopWordsR.ReadBytes('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Error in opening the file that stores stop words......", err)
		return
	}
	stopWordsArr := strings.Split(string(stopWordsLine), ",")

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
	var wordsDict []string
	var freqDict []int32
	for {
		lineBytes, err := r.ReadBytes('\n')
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
				flag := false
				if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
					word := strings.ToLower(line[startChar:i])
					if len(word) < 2 {
						startChar = -1
						continue
					}
					for _, stopWord := range stopWordsArr {
						if word == stopWord {
							startChar = -1
							flag = true
						}
					}
					for ind, w := range wordsDict {
						if w == word {
							freqDict[ind]++
							startChar = -1
							flag = true
							break
						}
					}
					if !flag {
						wordsDict = append(wordsDict, word)
						freqDict = append(freqDict, 1)
						startChar = -1
					}
				}
			}
		}
		flag := false
		if startChar != -1 {
			word := strings.ToLower(line[startChar:])
			if len(word) < 2 {
				startChar = -1
				flag = true
			}
			for _, stopWord := range stopWordsArr {
				if word == stopWord {
					startChar = -1
					flag = true
				}
			}
			for ind, w := range wordsDict {
				if w == word {
					freqDict[ind]++
					startChar = -1
					break
				}
			}
			if !flag {
				wordsDict = append(wordsDict, word)
				freqDict = append(freqDict, 1)
				startChar = -1
			}
		}
	}

	for i := 0; i < len(wordsDict)-1; i++ {
		maxIdx := i
		for j := i + 1; j < len(wordsDict); j++ {
			if freqDict[j] > freqDict[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx != i {
			wordsDict[maxIdx], wordsDict[i] = wordsDict[i], wordsDict[maxIdx]
			freqDict[maxIdx], freqDict[i] = freqDict[i], freqDict[maxIdx]
		}
	}
	for ind, word := range wordsDict[:25] {
		fmt.Printf("%s  -  %d\n", word, freqDict[ind])
	}
}
