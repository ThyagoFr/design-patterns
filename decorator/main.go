package main

import (
	"fmt"
)

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f",
		c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

// Decorator - Adicionando um novo atributo na interface shape, sem precisar mexer na classe shape
// Tambem adicionando um novo pointer receiver method

// Problema -
// Como trabalhamos com interfaces, mesmo utilizando um Circle, é perdido o acesso a todas as funcionalidades que existem apenas em Circle, mas nao na interface Shape
type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s",
		c.Shape.Render(), c.Color)
}

func (c *ColoredShape) NewFunctionality() string {
	return fmt.Sprintf("New functionality")
}

func main() {
	redCircle := ColoredShape{Shape: &Circle{}, Color: "Red"}
	fmt.Println(redCircle.NewFunctionality())
	fmt.Println(redCircle.Render())
}
