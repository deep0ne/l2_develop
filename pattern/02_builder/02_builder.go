package main

import "fmt"

type Computer struct {
	GPU       string
	Kernels   int
	Processor string
	Mouse     string
	Keyboard  string
	OS        string
}

type Builder interface {
	MakeInners(kernels int, gpu, processor string)
	MakeOuters(mouse, keyboard string)
	MakeSettings(os string)
}

type Director struct {
	builder Builder
}

func (d *Director) Construct(kernels int, gpu, processor, mouse, keyboard, os string) {
	d.builder.MakeInners(kernels, gpu, processor)
	d.builder.MakeOuters(mouse, keyboard)
	d.builder.MakeSettings(os)
	fmt.Println("Computer is ready to use")
}

type ConcreteBuilder struct {
	product *Computer
}

func (b *ConcreteBuilder) MakeInners(kernels int, gpu, processor string) {
	b.product.Kernels = kernels
	b.product.GPU = gpu
	b.product.Processor = processor
	fmt.Println("Inner parts of computer are set")
}

func (b *ConcreteBuilder) MakeOuters(mouse, keyboard string) {
	b.product.Keyboard = keyboard
	b.product.Mouse = mouse
	fmt.Println("Keyboard and mouse are set")
}

func (b *ConcreteBuilder) MakeSettings(os string) {
	b.product.OS = os
	fmt.Printf("Installing %s...\n", os)
}

func (c *Computer) Show() {
	fmt.Printf("Our computer is working on %s. ", c.OS)
	fmt.Printf("We have %d kernels, %s proc and %s gpu. ", c.Kernels, c.Processor, c.GPU)
	fmt.Printf("We bought %s and %s for our computer", c.Keyboard, c.Mouse)
}

func main() {
	computer := Computer{}
	director := Director{&ConcreteBuilder{&computer}}
	director.Construct(8, "1050ti", "i7", "SteelSeries", "Razor", "Windows")
	computer.Show()
}
