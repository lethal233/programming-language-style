package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

type kv struct {
	Word string
	Freq int
}
type TFQuarantine struct {
	funcs []func(interface{}) interface{}
}

func TFQuarantineConstructor(fc func(interface{}) interface{}) *TFQuarantine {
	tf := &TFQuarantine{funcs: []func(interface{}) interface{}{fc}}
	return tf
}

func (tf *TFQuarantine) bind(fc func(interface{}) interface{}) *TFQuarantine {
	tf.funcs = append(tf.funcs, fc)
	return tf
}
func (tf *TFQuarantine) execute() {
	guardCallable := func(v interface{}) interface{} {
		if reflect.TypeOf(v).Kind() == reflect.Func {
			return v.(func() interface{})()
		}
		return v
	}
	var value interface{} = func() interface{} { return nil }

	for _, f := range tf.funcs {
		value = f(guardCallable(value))
	}
	fmt.Print(guardCallable(value))
}

func getInput(arg interface{}) interface{} {
	fc := func() interface{} { return os.Args[1] }
	return fc
}

func extractWords(path2file interface{}) interface{} {
	fc := func() interface{} {
		var da []rune
		file, err := os.Open(path2file.(string))
		if err != nil {
			os.Exit(1)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			da = append(da, []rune(scanner.Text())...)
			da = append(da, ' ')
		}
		reg, err := regexp.Compile("[\\W_]+")
		if err != nil {
			panic(err)
		}
		return strings.Fields(reg.ReplaceAllString(strings.ToLower(string(da)), " "))
	}
	return fc
}

func removeStopWords(words interface{}) interface{} {
	fc := func() interface{} {
		word := words.([]string)
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
		for i := 0; i < len(word); i++ {
			if !stopWordsSet[word[i]] {
				filteredWords = append(filteredWords, word[i])
			}
		}
		return filteredWords
	}
	return fc
}

func frequencies(words interface{}) interface{} {
	wf := make(map[string]int)
	for _, w := range words.([]string) {
		wf[w]++
	}
	return wf
}
func sorted(wordFreqs interface{}) interface{} {
	var sortedFreqs []kv
	for k, v := range wordFreqs.(map[string]int) {
		sortedFreqs = append(sortedFreqs, kv{k, v})
	}
	sort.Slice(sortedFreqs, func(i, j int) bool {
		return sortedFreqs[i].Freq > sortedFreqs[j].Freq
	})
	return sortedFreqs
}
func top25Freq(wordFreqs interface{}) interface{} {
	top25 := ""
	for _, freq := range wordFreqs.([]kv)[:25] {
		top25 += fmt.Sprintf("%s  -  %d\n", freq.Word, freq.Freq)
	}
	return top25
}
func main() {
	tf := TFQuarantineConstructor(getInput)
	tf.bind(extractWords).
		bind(removeStopWords).
		bind(frequencies).
		bind(sorted).
		bind(top25Freq).
		execute()
}
