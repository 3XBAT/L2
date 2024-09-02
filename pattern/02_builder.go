package main

import "fmt"

type Product struct { // сложностроящийся продукт
	Property1 string
	Property2 string
}

type IBuilder interface { // интерфейс билдера
	setProperty1()
	setProperty2()
	getProduct() Product
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}
	if builderType == "unusual" {
		return newUnusualBuilder()
	}
	return nil
}

type NormalBuilder struct {
	Property1 string
	Property2 string
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setProperty1() {
	b.Property1 = "normal 1"
}
func (b *NormalBuilder) setProperty2() {
	b.Property2 = "normal 2"
}

func (b *NormalBuilder) getProduct() Product {
	return Product{
		Property1: b.Property1,
		Property2: b.Property2,
	}

}

type UnusualBuilder struct {
	Property1 string
	Property2 string
}

func newUnusualBuilder() *UnusualBuilder {
	return &UnusualBuilder{}
}

func (b *UnusualBuilder) setProperty1() {
	b.Property1 = "unusual 1"
}

func (b *UnusualBuilder) setProperty2() {
	b.Property2 = "unusual 2"
}

func (b *UnusualBuilder) getProduct() Product {
	return Product{
		Property1: b.Property1,
		Property2: b.Property2,
	}
}

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildProduct() Product {
	d.builder.setProperty1()
	d.builder.setProperty2()
	return d.builder.getProduct()
}

func main() {

	director := newDirector(getBuilder("normal"))
	normalProduct := director.buildProduct()
	fmt.Println("normalProduct:", normalProduct)

	director.setBuilder(getBuilder("unusual"))
	unusualProduct := director.buildProduct()
	fmt.Println("unusualProduct:", unusualProduct)
}
