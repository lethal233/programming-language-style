package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

type words struct {
	Word string
	Freq int32
}

func main() {
	path := os.Args[1]
	fp, err := os.Open(path)
	stopWordsPath := "/Volumes/mac-hdd/projects/APL-go/stop_words.txt"
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(fp)
	dict := map[string]int32{}
	stopWords, err := os.Open(stopWordsPath)
	if err != nil {
		panic(err)
	}
	stopWordsSet := map[string]bool{}
	stopWordsR := bufio.NewReader(stopWords)
	stopWordsLine, err := stopWordsR.ReadBytes('\n')
	stopWordsArr := strings.Split(string(stopWordsLine), ",")
	for i := 0; i < len(stopWordsArr); i++ {
		stopWordsSet[stopWordsArr[i]] = true
	}
	regExp := regexp.MustCompile("[0-9a-zA-Z]+")
	for {
		lineBytes, err := r.ReadBytes('\n') // "word--word"
		line := strings.TrimSpace(string(lineBytes))
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		arr := strings.Fields(line)
		for i := 0; i < len(arr); i++ {
			w := regExp.FindAllString(arr[i], -1)
			for i := 0; i < len(w); i++ {
				w[i] = strings.ToLower(w[i])
				if stopWordsSet[w[i]] || len(w[i]) < 2 {
					continue
				}
				v, ok := dict[w[i]]
				if ok {
					dict[w[i]] = v + 1
				} else {
					dict[w[i]] = 1
				}
			}
		}
	}

	var lstWords []words
	for k, v := range dict {
		lstWords = append(lstWords, words{k, v})
	}

	sort.Slice(lstWords, func(i, j int) bool {
		if lstWords[i].Freq == lstWords[j].Freq {
			return lstWords[i].Word < lstWords[j].Word
		} else {
			return lstWords[i].Freq > lstWords[j].Freq
		}
	})

	minNum := 25
	for i := 0; i < min(minNum, len(lstWords)); i++ {
		fmt.Printf("%s  -  %d", lstWords[i].Word, lstWords[i].Freq)
		if i != min(minNum, len(lstWords))-1 {
			fmt.Println()
		}
	}
}

func min(i int, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
