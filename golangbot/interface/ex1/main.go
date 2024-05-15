package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

type Contract struct {
	empId    int
	basicPay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicPay
}

// if you want another employee type just add here

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	p1 := Permanent{
		empId:    1,
		basicPay: 5000,
		pf:       20,
	}
	p2 := Permanent{
		empId:    2,
		basicPay: 6000,
		pf:       30,
	}
	c1 := Contract{
		empId:    3,
		basicPay: 3000,
	}
	employees := []SalaryCalculator{p1, p2, c1}
	totalExpense(employees)
}
