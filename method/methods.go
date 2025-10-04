package main

import (
	"fmt"
)

type MyType string

// sayHi is unexported
func (m MyType) sayHi() {
	fmt.Println("Hi")
}

// sayHello is unexported
func (m MyType) sayHello() {
	fmt.Println("Hi from", m)
}

// MethodWithParameter is exported
func (m MyType) MethodWithParameter(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

// WithReturn is exported
func (m MyType) WithReturn() int {
	return len(m)
}

type Number int

func (n *Number) Double() {
	*n *= 2
}

func (n *Number) Display() {
	fmt.Println(*n)
}

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	value := MyType("A value")
	value.sayHi()

	MyType("A different value").sayHi()

	val := MyType("a WhyType value")
	val.sayHello()

	MyType("a different WhyType value").sayHello()

	MyType("MyType value:").MethodWithParameter(4, true)

	fmt.Println(MyType("MyType value").WithReturn())

	number := Number(4)
	fmt.Printf("number before Double method: %v\n", number)
	number.Double()
	fmt.Printf("number after Double method: %v\n", number)

	val2 := MyType("a value")
	pointer := &val2
	val2.method() // pointer of val2 automatically taken
	val2.pointerMethod()
	pointer.method() // value at pointer automatically retrieved
	pointer.pointerMethod()
	// This breaks convention
	// For consistency, all the methods of a type should either
	// take a value or a pointer; should avoid mixing.

	// Pointers can only be taken from values stored in variables.
	// Same limitation applies to calling methods with pointer receivers.

	number = Number(6)
	number.Double()
	number.Display()
}
