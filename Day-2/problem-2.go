package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type teacher struct {
	rating float64
}

type student struct {
	rating    int
	timeTaken int
}

func (stud *student) giveRating() int {
	stud.rating = rand.Intn(10)
	return stud.rating
}

// calcRating calculates teacher ratings
func (tch *teacher) calcRating(students []student) (float64, error) {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	var total int64
	totalStudents := len(students)
	if totalStudents == 0 {
		return 0, errors.New("null value passed")
	}

	for id, st := range students {
		wg.Add(1)
		go func(st student, id int) {
			defer wg.Done()
			st.timeTaken = rand.Intn(10)
			time.Sleep(time.Duration(st.timeTaken))
			atomic.AddInt64(&total, int64(st.giveRating()))
			fmt.Printf("Student %d rating %d\n", id, st.rating)
		}(st, id)
	}

	wg.Wait()

	tch.rating = float64(total) / float64(totalStudents)
	return tch.rating, nil
}

func main() {
	students := make([]student, 10)
	tch := teacher{}
	_, err := tch.calcRating(students)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Average Rating of teacher %f\n", tch.rating)
}
