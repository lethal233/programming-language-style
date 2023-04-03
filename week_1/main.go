package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {

	path := os.Args[1]
	fp, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(fp)
	dict := map[string]int32{}
	for {
		lineBytes, err := r.ReadBytes('\n')
		line := strings.TrimSpace(string(lineBytes))
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		arr := strings.Split(line, "[^0-9a-zA-Z]+")
		for i := 0; i < len(arr); i++ {
			v, ok := dict[arr[i]]
			if ok {
				dict[arr[i]] = v + 1
			} else {
				dict[arr[i]] = 1
			}
		}
	}

	type words struct {
		Word string
		Freq int32
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
	var minNum = int32(25)
	for i := 0; i < min(minNum, int32(len(lstWords))); i++ {
		fmt.Printf("%s  -  %d\n", lstWords[i].Word, lstWords[i].Freq)
	}
}

func min(i int32, j int32) int32 {
	if i > j {
		return j
	} else {
		return i
	}
}
