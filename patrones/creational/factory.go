package main

import "errors"

// Patron Creacional

// Interface que define el comportamiento de un producto
type IProduct interface {
	setStocked(stock int)
	getStocked() int

	getName() string
	setName(name string)
}

// Implementacion de la interfaz IProduct para el producto de tipo "Computadora"
type Computer struct {
	name  string
	stock int
}

// Implementando de forma implicita la interfaz IProduct
func (c *Computer) setStocked(stock int) {
	c.stock = stock
}

func (c *Computer) getStocked() int {
	return c.stock
}

func (c *Computer) getName() string {
	return c.name
}

func (c *Computer) setName(name string) {
	c.name = name
}

// Creando clase base de computadora, por composicion sobre herencia
type Laptop struct {
	Computer
}

func NewLaptop() IProduct {
	return &Laptop{Computer{"Laptop", 25}}
}

type Desktop struct {
	Computer
}

func NewDesktop() IProduct {
	return &Desktop{Computer{"Desktop", 35}}
}

// Creando fabrica de productos: Factory pattern
func GetComputerFactory(computerType string) (IProduct, error) {
	switch computerType {
	case "Laptop":
		return NewLaptop(), nil
	case "Desktop":
		return NewDesktop(), nil
	default:
		return nil, errors.New("invalid computer type")
	}
}

// Trying polymorphism
func PrintNameAndStock(product IProduct) {
	println("Name:", product.getName(), "Stock:", product.getStocked())
}

func main() {
	laptop, _ := GetComputerFactory("Laptop")
	desktop, _ := GetComputerFactory("Desktop")

	PrintNameAndStock(laptop)
	PrintNameAndStock(desktop)
}
