package main

import "fmt"

// salary Interface contains a method calc_salary
type salary interface{
	calc_salary() float64
}

// employee struct contains basic which denotes basic pay per unit duration
// and the total duration worked
type employee struct{
	basic float64
	duration float64
}

// calculates salary of an employee on the basis of basic and duration worked
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
