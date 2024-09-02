package main

import "fmt"

type Subsystem1 struct {
}

func (s Subsystem1) operation1() {
	fmt.Println("Subsystem1 do something")
}

type Subsystem2 struct {
}

func (s Subsystem2) operation2() {
	fmt.Println("Subsystem2 do something")
}

type Facade struct {
	Subsystem1
	Subsystem2
}

func NewFacade() *Facade {
	return &Facade{
		Subsystem1: Subsystem1{},
		Subsystem2: Subsystem2{},
	}

}

func (f *Facade) difficultOperation() {
	f.operation1()
	f.operation2()
}

func main() {
	facade := NewFacade()
	facade.difficultOperation()
}
