package main

import "fmt"

type Animal interface {
	setTalk(talk string)
	setMoveType(moveType string)
	getTalk() string
	getMoveType() string
}

type Dog struct {
	talk     string
	moveType string
}

func newDog() Animal {
	return &Dog{}
}

func (d *Dog) setTalk(talk string) {
	d.talk = talk
}

func (d *Dog) setMoveType(moveType string) {
	d.moveType = moveType
}

func (d *Dog) getTalk() string {
	return d.talk
}

func (d *Dog) getMoveType() string {
	return d.moveType
}

func main() {
	animal := newDog()
	animal.setMoveType("run")
	animal.setTalk("gau gau")
	fmt.Println(animal.getMoveType())
	fmt.Println(animal.getTalk())
}
