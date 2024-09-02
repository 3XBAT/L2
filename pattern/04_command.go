package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Switch struct {
	command Command
}

func (s *Switch) press() {
	s.command.execute()
}

type Command interface {
	execute()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

type Device interface {
	on()
	off()
}

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Tv off")
}

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}
	offCommand := &OffCommand{
		device: tv,
	}

	onSwitch := &Switch{
		command: onCommand,
	}
	onSwitch.press()

	offSwitch := &Switch{
		command: offCommand,
	}
	offSwitch.press()

}
