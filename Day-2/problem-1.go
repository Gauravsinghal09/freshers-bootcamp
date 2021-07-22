package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Init initialises the frequency map for 26 characters
func Init(letterFreq map[string]int) {
	for char := 'a'; char <= 'z'; char++ {
		letterFreq[string(char)] = 0
	}
}

// preProcess transforms the string to lower case
func preProcess(wordArray []string) {
	for _, word := range wordArray {
		word = strings.ToLower(word)
	}
}

// calcLetterFrequency calculates the frequency of each letter
func calcLetterFrequency(wordArray []string, letterFreq map[string]int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	for _, word := range wordArray {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			for _, letter := range word {
				mutex.Lock()
				letterFreq[string(letter)]++
				mutex.Unlock()
				time.Sleep(time.Millisecond)
			}
		}(word)
	}
}

func main() {

	var wg sync.WaitGroup
	var mutex = sync.Mutex{}
	var wordArray = []string{"quick", "dog", "dog", "dog", "dog"}
	var letterFreq = make(map[string]int)

	Init(letterFreq)
	preProcess(wordArray)

	calcLetterFrequency(wordArray, letterFreq, &mutex, &wg)
	wg.Wait()

	fmt.Println(letterFreq)
}
