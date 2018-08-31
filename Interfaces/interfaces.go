package main

import (
	"fmt"
	"math"
)

//Интерфейс для геометрических фигур
type geometry interface {
	area() float64
	perim() float64
}

//Структуры
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

//Методы
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//Функция
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r) 	//{3 4}
				//12
				//14
	measure(c) 	//{5}
				//78.53981633974483
				//31.41592653589793
}
