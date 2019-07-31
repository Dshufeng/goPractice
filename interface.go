package main

import "fmt"

// page:https://www.golangtc.com/t/5cce80a9b17a82478bd85f3c
type Human struct {
	Name string
	Age  int
}

type Student struct {
	Human
	School string
	Loan   float32
}

type Employee struct {
	Human
	Company string
	Money   float32
}

func (h Human) SayHi() {
	fmt.Println("human say hi")
}
func (h Human) Sing(song string) {
	fmt.Println("sing...", song)
}
func (s Student) SayHi() {
	fmt.Println("student say hi")
}

func (s Student) Learn() {
	fmt.Println("student learn")
}

func (e Employee) Sing(song string) {
	fmt.Println("sing...", song)
}
func (e *Employee) Offer() {
	e.Money = 1111
	fmt.Println(e.Money)
}

// interface
type Men interface {
	SayHi()
	Sing(song string)
}
type Stu interface {
	SayHi()
	Sing(song string)
	Learn()
}

type Emp interface {
	SayHi()
	Sing(song string)
	Offer()
}

func main() {
	//wang := Human{"wang", 19}

	student := Student{Human{"li", 20}, "bd", 100.00}
	var i Men
	i = student // 实现了interface 里的全部方法
	i.Sing("eee")
	i.SayHi()

	emp := Employee{Human{"wang", 18}, "jd", 10000.00}

	i = emp // 实现了Men interface 里的全部方法
	i.SayHi()
	i.Sing("aaa")

	(&emp).Offer()
	emp.Sing("aaa")
	fmt.Println(emp)
}
