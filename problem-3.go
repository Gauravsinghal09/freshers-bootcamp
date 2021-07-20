package main

import "fmt"

type salary interface{
	calc_salary() float64
}

type employee struct{
	basic float64
	duration float64
}

func (emp *employee) calc_salary() float64{
	return emp.basic * emp.duration
}

func main(){
	var fulltime = employee{500 ,28}
	var contractor = employee{100, 28}
	var freelancer = employee{50, 20}
	fmt.Println(fulltime.calc_salary())
	fmt.Println(contractor.calc_salary())
	fmt.Println(freelancer.calc_salary())
}
