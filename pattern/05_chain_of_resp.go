package main

import "fmt"

type Task struct {
	name       string
	condition1 bool
	condition2 bool
	condition3 bool
}

type Handler interface {
	execute(*Task)
	setNext(handler Handler)
}

type handler1 struct {
	next Handler
}

func (h *handler1) execute(t *Task) {
	if t.condition1 {
		fmt.Println("first handler already executed")
		h.next.execute(t)
		return
	}
	fmt.Println("first handler is working")
	t.condition1 = true
	h.next.execute(t)
}

func (h *handler1) setNext(handler Handler) {
	h.next = handler
}

type handler2 struct {
	next Handler
}

func (h *handler2) execute(t *Task) {
	if t.condition2 {
		fmt.Println("second handler already executed")
		h.next.execute(t)
		return
	}
	fmt.Println("second handler is working")
	t.condition2 = true
	h.next.execute(t)
}

func (h *handler2) setNext(handler Handler) {
	h.next = handler
}

type handler3 struct {
	next Handler
}

func (h *handler3) execute(t *Task) {
	if t.condition3 {
		fmt.Println("third handler already executed")

	}
	fmt.Println("third handler is working")

}

func (h *handler3) setNext(handler Handler) {
	h.next = handler
}

func main() {
	h3 := &handler3{}

	h2 := &handler2{}
	h2.setNext(h3)

	h1 := &handler1{}
	h1.setNext(h2)

	task := &Task{
		name: "task",
	}

	h1.execute(task)
}
