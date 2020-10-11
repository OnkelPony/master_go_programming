package main

import "fmt"

type car struct {
	brand string
	price int
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	c.price = newPrice
}

func main() {
	myCar := car{brand: "Lexus", price: 80000}
	myCar.changeCar1("Peugeot", 20000)
	fmt.Println(myCar)

	(&myCar).changeCar2("Acura", 78000)
	fmt.Println(myCar)

	var yourCar *car
	yourCar = &myCar
	fmt.Println(*yourCar)

	yourCar.changeCar2("Maserati", 90000)
	fmt.Println(*yourCar)
	(*yourCar).changeCar2("Maserati", 108000)
	fmt.Println(myCar)
}
