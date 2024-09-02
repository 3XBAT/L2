package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "Ak 47",
			power: 500,
		},
	}
}

type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket",
			power: 150,
		},
	}
}

func getGun(gunType string) (IGun, error) {
	if gunType == "Ak47" {
		return newAk47(), nil
	}
	if gunType == "Musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("wrong type of gun")
}

func main() {
	ak47, _ := getGun("Ak47")
	musket, _ := getGun("Musket")

	fmt.Println(ak47)
	fmt.Println(musket)

}
