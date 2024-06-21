package main

import "fmt"

type HouseBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

type House struct {
	windowType string
	doorType   string
	floor      int
}

type Director struct {
	builder HouseBuilder
}

func newDirector(b HouseBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b HouseBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

type NormalBuilder struct {
	House
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type DiamondBuilder struct {
	House
}

func newDiamondBuilder() *DiamondBuilder {
	return &DiamondBuilder{}
}

func (b *DiamondBuilder) setWindowType() {
	b.windowType = "Diamond Window"
}

func (b *DiamondBuilder) setDoorType() {
	b.doorType = "Diamond Door"
}

func (b *DiamondBuilder) setNumFloor() {
	b.floor = 5
}

func (b *DiamondBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

func main() {
	normalBuilder := newNormalBuilder()
	diamondBuilder := newDiamondBuilder()

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Println("================================================================")
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(diamondBuilder)
	diamondHouse := director.buildHouse()

	fmt.Println("================================================================")
	fmt.Printf("Diamond House Door Type: %s\n", diamondHouse.doorType)
	fmt.Printf("Diamond House Window Type: %s\n", diamondHouse.windowType)
	fmt.Printf("Diamond House Num Floor: %d\n", diamondHouse.floor)
}
