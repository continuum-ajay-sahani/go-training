package main

import "fmt"

type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

func vehicleTest() {
	prius := car{}
	tacoma := car{}
	bmw528 := car{}
	boeing747 := plane{}
	boeing757 := plane{}
	boeing767 := plane{}

	rides := []vehicles{prius, tacoma, bmw528, boeing747, boeing757, boeing767}

	for key, value := range rides {
		fmt.Println(key, " - ", value)
	}
}

func main() {
	//vehicleTest()
	animalTest()
}

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func specs(a interface{}) {
	fmt.Println(a)
}

func animalTest() {
	fido := dog{animal{"woof"}, true}
	fifi := cat{animal{"meow"}, true}
	specs(fido)
	specs(fifi)
}
