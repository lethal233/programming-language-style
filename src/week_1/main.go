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
	stopWordsSet, err := readStopWords("../stop_words.txt")
	if err != nil {
		fmt.Println("Error in opening the file that stores stop words......", err)
		return
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
	regExp := regexp.MustCompile("[0-9a-zA-Z]+")
	for {
		lineBytes, err := r.ReadBytes('\n') // "word--word"
		line := strings.TrimSpace(string(lineBytes))
		if err != nil && err != io.EOF {
			fmt.Println("Unexpected error")
			return
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
		fmt.Printf("%s  -  %d\n", lstWords[i].Word, lstWords[i].Freq)
	}
}

func readStopWords(stopWordsPath string) (map[string]bool, error) {
	stopWords, err := os.Open(stopWordsPath)
	if err != nil {
		return nil, err
	}
	defer stopWords.Close()
	stopWordsSet := map[string]bool{}
	stopWordsR := bufio.NewReader(stopWords)
	stopWordsLine, err := stopWordsR.ReadBytes('\n')
	stopWordsArr := strings.Split(string(stopWordsLine), ",")
	for i := 0; i < len(stopWordsArr); i++ {
		stopWordsSet[stopWordsArr[i]] = true
	}
	return stopWordsSet, nil
}

func min(i int, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
