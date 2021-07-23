package main

import (
	"fmt"
	"strings"
	"sync"
)

type Counter struct {
	mutex      sync.Mutex
	wordArray  []string
	letterFreq map[string]int
}

// Init initialises the frequency map for 26 characters
func (counter *Counter) Init() {
	counter.letterFreq = make(map[string]int)
	for char := 'a'; char <= 'z'; char++ {
		counter.letterFreq[string(char)] = 0
	}
}

// preProcess transforms the string to lower case
func (counter *Counter) preProcess() {
	for id, word := range counter.wordArray {
		counter.wordArray[id] = strings.ToLower(word)
	}
}

// calcLetterFrequency calculates the frequency of each letter
func (counter *Counter) calcLetterFrequency(wg *sync.WaitGroup) {
	for _, word := range counter.wordArray {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			for _, letter := range word {
				counter.mutex.Lock()
				counter.letterFreq[string(letter)]++
				counter.mutex.Unlock()
			}
		}(word)
	}
}

// prints the frequency of each letter
func (counter *Counter) printFrequency() {
	fmt.Println(counter.letterFreq)
}

func main() {

	var wg sync.WaitGroup
	counter := Counter{}
	counter.wordArray = []string{"quick", "dog", "dog", "dog", "dog"}
	counter.Init()
	counter.preProcess()

	counter.calcLetterFrequency(&wg)
	wg.Wait()
	counter.printFrequency()
}
