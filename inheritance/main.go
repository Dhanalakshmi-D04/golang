package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Intro() {
	fmt.Println(p.Age, p.Name)
}

type Employee struct {
	Person
	EmployeeId int
}

func (e Employee) Intro() {
	fmt.Println(e.EmployeeId, e.Name, e.Age)
}

func main() {
	p := Person{Name: "abc", Age: 21}
	ee := Employee{Person: p, EmployeeId: 12}
	ee.Person.Intro()
	ee.Intro()
	p.Intro()
}
