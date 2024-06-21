package main

import "fmt"

// Subject interface
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

// Observer interface
type Observer interface {
	Update(subject Subject)
}

// GameCharacter là subject
type GameCharacter struct {
	observers []Observer
	health    int
}

func (g *GameCharacter) RegisterObserver(o Observer) {
	g.observers = append(g.observers, o)
}

func (g *GameCharacter) RemoveObserver(o Observer) {
	var indexToRemove int
	for i, observer := range g.observers {
		if observer == o {
			indexToRemove = i
			break
		}
	}
	g.observers = append(g.observers[:indexToRemove], g.observers[indexToRemove+1:]...)
}

func (g *GameCharacter) NotifyObservers() {
	for _, observer := range g.observers {
		observer.Update(g)
	}
}

func (g *GameCharacter) TakeDamage(amount int) {
	g.health -= amount
	g.NotifyObservers()
}

// HealthSystem là một observer
type HealthSystem struct{}

func (h *HealthSystem) Update(subject Subject) {
	// Cập nhật hệ thống sức khỏe dựa trên trạng thái của nhân vật chính
	if character, ok := subject.(*GameCharacter); ok {
		fmt.Printf("HealthSystem: Character has %d health remaining\n", character.health)
	}
}

func main() {
	character := &GameCharacter{health: 100}
	healthSystem := &HealthSystem{}

	character.RegisterObserver(healthSystem)

	// Nhân vật chính nhận sát thương và hệ thống sức khỏe được thông báo
	character.TakeDamage(10)
}
