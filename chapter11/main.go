package main

import "fmt"

type Car struct{}
type Aircraft struct{}

func (c Car) move() string {
	return "Автомобиль едет"
}
func (a Aircraft) move() string {
	return "Самолет летит"
}

func main() {

	var tesla Car = Car{}
	var boing Aircraft = Aircraft{}
	fmt.Println(tesla.move())
	fmt.Println(boing.move())
}
