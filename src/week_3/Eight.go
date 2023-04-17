package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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
	//sort.Slice(sWds, func(i, j int) bool { return wordFreq[sWds[i]] > wordFreq[sWds[j]] })

	quickSort(sWds, 0, len(sWds)-1, wordFreq)

	for i := 0; i < 25; i++ {
		fmt.Printf("%s - %d\n", sWds[i], wordFreq[sWds[i]])
	}
}

func quickSort(arr []string, low int, high int, wordFreq map[string]int) {
	if low < high {
		pivot := partition(arr, low, high, wordFreq)
		quickSort(arr, low, pivot-1, wordFreq)
		quickSort(arr, pivot+1, high, wordFreq)
	}
}

func partition(arr []string, low int, high int, wordFreq map[string]int) int {
	// Choose a random pivot index
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)
	pivotIndex := randGen.Intn(high-low+1) + low
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

	pivot := arr[high]
	i := low

	for j := low; j <= high-1; j++ {
		if wordFreq[arr[j]] > wordFreq[pivot] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}
