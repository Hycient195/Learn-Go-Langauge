package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Methods in Go Langauge");

	var card Rectangle = Rectangle{ Height: 4, Width: 6 };
	fmt.Println(card.GetArea());
	fmt.Println(card.GetPerimeter());

	// Comparing between value recievers and pointer recievers;
	card.ModifyCopy();
	fmt.Println(card.Width, card.Height);
	card.ModifyReference();
	fmt.Println(card.Width, card.Height);
}

type Rectangle struct {
	Height float32
	Width  float32
}
func (r Rectangle) GetArea() float32 {
	return r.Height * r.Width;
}
func (r Rectangle) GetPerimeter() float32 {
	return 2*r.Height + 2*r.Width;
}
func (r Rectangle) ModifyCopy() {
	r.Height = 2;
	r.Width = 3;
}
func (r *Rectangle) ModifyReference() {
	(*r).Height = 8;
	(*r).Width = 10;
}
