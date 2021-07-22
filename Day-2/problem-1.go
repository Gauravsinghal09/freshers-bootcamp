package main

import (
	"fmt"
	"sync"
)

func calcLetterFrequency(wordArray []string, letterFreq map[string]int, done chan bool) {
	var wg sync.WaitGroup

	fmt.Println("Counting Starts")
	var mutex = &sync.Mutex{}
	for _, word := range wordArray {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			for _, letter := range word {
				go func(letter int32) {
					mutex.Lock()
					letterFreq[string(letter)]++
					mutex.Unlock()
				}(letter)
			}
		}(word)
	}

	wg.Wait()
	fmt.Println("Counting Ends")
	done <- true
}

func main() {

	var wordArray = []string{"quick", "brown", "fox", "lazy", "dog"}
	var letterFreq = make(map[string]int)

	done := make(chan bool)

	go calcLetterFrequency(wordArray, letterFreq, done)
	<-done

	fmt.Println(letterFreq)
}
