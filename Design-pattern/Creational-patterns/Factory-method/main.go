package main

import (
	"fmt"
)

type ITransport interface {
	GetName() string
	SetName(name string)
	SetDestination(from string, to string)
	GetInfo() string
}

// Truck factory
type TruckTransport struct {
	name string
	from string
	to   string
}

func (tt *TruckTransport) GetName() string {
	return tt.name
}

func (tt *TruckTransport) SetName(name string) {
	tt.name = name
}

func (tt *TruckTransport) SetDestination(from string, to string) {
	tt.from = from
	tt.to = to
}

func (tt *TruckTransport) GetInfo() string {
	return fmt.Sprintf("%s is delivering a package from %s to %s", tt.name, tt.from, tt.to)
}

func NewTruck() ITransport {
	return &TruckTransport{}
}

// Ship factory
type ShipFactory struct {
	name string
	from string
	to   string
}

func (t *ShipFactory) GetName() string {
	return t.name
}

func (t *ShipFactory) SetName(name string) {
	t.name = name
}

func (t *ShipFactory) SetDestination(from string, to string) {
	t.from = from
	t.to = to
}

func (t *ShipFactory) GetInfo() string {
	return fmt.Sprintf("%s is delivering a container from %s to %s", t.name, t.from, t.to)
}

func NewShip() ITransport {
	return &ShipFactory{}
}

func getTransport(transportType string) (ITransport, error) {
	if transportType == "truck" {
		return NewTruck(), nil
	}
	if transportType == "ship" {
		return NewShip(), nil
	}
	return nil, fmt.Errorf("wrong trainsport type")
}

func main() {
	truck, _ := getTransport("truck")
	ship, _ := getTransport("ship")

	truck.SetName("Monster Truck")
	truck.SetDestination("Ho Chi Minh", "An Giang")

	ship.SetName("Battle Ship")
	ship.SetDestination("USA", "Australia")

	fmt.Println(truck.GetInfo())
	fmt.Println(ship.GetInfo())
}
