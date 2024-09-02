package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/
// Посетитель должен расширяет функционал сущесвтующих структур
// в чем смысл
// мы делаем интерфейс Visitor, который может как-то взаимодействовать с объектами
// также нужно сделать метод accept в самих структурах, через него мы будем запускать методы, которые и расширяют функционал

// ЗАДАЧА: расширить функционал интефейса Shape(там только getType) так чтобы можно было посчитать площадь фигуры

type Shape interface {
	getType()
	accept(Visitor)
}

type Square struct {
	side int
}

func (sq *Square) getType() {
	fmt.Println("Square")
}

func (sq *Square) accept(v Visitor) {
	v.visitForSquare(sq)
}

type Circle struct {
	radius int
}

func (c *Circle) getType() {
	fmt.Println("Circle")
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

type Rectangle struct {
	a int
	b int
}

func (r *Rectangle) getType() {
	fmt.Println("Rectangle")
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {

	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(c *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaCalculator) visitForRectangle(r *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}

func main() {
	square := Square{5}
	circle := Circle{10}
	rectangle := Rectangle{10, 20}

	ac := AreaCalculator{}

	square.accept(&ac)
	circle.accept(&ac)
	rectangle.accept(&ac)
}
