package main

import (
	"errors"
	"fmt"
	"sync"
)

type Teacher struct {
	rating      float64
	totalRating int64
}

type Student struct {
	rating    int64
	timeTaken int
}

func (stud *Student) GiveRating() int64 {
	stud.rating = 10
	return stud.rating
}

// CalcRating calculates teacher ratings
func (tch *Teacher) CalcRating(students []Student) (float64, error) {

	var wg sync.WaitGroup
	totalStudents := len(students)
	if totalStudents == 0 {
		return 0, errors.New("null value passed")
	}

	studentRating := make(chan int64, totalStudents)
	var total int64

	for _, st := range students {
		wg.Add(1)
		go func(st Student) {
			defer wg.Done()
			//st.timeTaken = rand.Intn(10)
			//time.Sleep(time.Duration(st.timeTaken))
			studentRating <- st.GiveRating()
		}(st)
	}

	wg.Wait()
	close(studentRating)
	for rating := range studentRating {
		total += rating
	}

	tch.totalRating = total
	tch.rating = float64(total) / float64(totalStudents)
	return tch.rating, nil
}

func (tch *Teacher) PrintRating() {
	fmt.Printf("Total Rating: %d \n", tch.totalRating)
	fmt.Printf("Average Rating: %f \n", tch.rating)
}

func main() {
	students := make([]Student, 10)
	tch := Teacher{}
	_, err := tch.CalcRating(students)
	if err != nil {
		fmt.Println(err)
	}
	tch.PrintRating()
}
