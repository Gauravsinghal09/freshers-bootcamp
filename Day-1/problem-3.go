package main

import "fmt"

// Employee Interface contains a method calc_salary
// calcSalary calculates the salary of an employee
type Employee interface {
	calcSalary() float64
}

// FullTime struct contains basic which denotes basic pay per unit duration
// duration in days
type FullTime struct {
	basic    float64
	duration float64
}

func (fte *FullTime) calcSalary() float64 {
	return fte.basic * fte.duration
}

// Contractor struct contains basic which denotes basic pay per unit duration
// duration in days
type Contractor struct {
	basic    float64
	duration float64
}

func (c *Contractor) calcSalary() float64 {
	return c.basic * c.duration
}

// Freelancer struct contains basic which denotes basic pay per unit duration
// duration in hours
type Freelancer struct {
	basic    float64
	duration float64
}

func (f *Freelancer) calcSalary() float64 {
	return f.basic * f.duration
}

// calcTotalExpense calculates the total expense the employer will bear
func calcTotalExpense(employees []Employee) float64 {
	expense := 0.0
	for _, employee := range employees {
		expense += employee.calcSalary()
	}
	return expense
}

func main() {
	var fte = FullTime{500, 28}
	var contractor = Contractor{100, 28}
	var freelancer = Freelancer{10, 20}
	var employees = []Employee{&fte, &contractor, &freelancer}
	fmt.Println(calcTotalExpense(employees))
}
