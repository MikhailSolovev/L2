package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	Когда нужно выполнить операцию над всеми элементами сложной структуры объектов.
	Когда над объектами сложной структуры объектов надо выполнять некоторые, не связанные между собой операции, и не
	хочется засорять классы такими операциями.
	Когда новое поведения имеет смысл только для некоторых классов из существующей иерархии.

	+ Упрощает добавление новых операций над всей связанной структурой объектов
	+ Объединяет родственные операции в одном классе
	+ Посетитель может накапливать состояние при обходе структуры компонентов
	- Паттерн неоправдан, если иерархия компонентов часто меняется
	- Может привести к нарушению инкапсуляции компонентов
*/

// Общий интерфейс фигур

type shape interface {
	getType() string
	accept(visitor)
}

// Структура квадрат

type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "Square"
}

//Структура круг

type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}

// Структура прямоугольник

type rectangle struct {
	l int
	b int
}

func (r *rectangle) accept(v visitor) {
	v.visitForRectangle(r)
}

func (r *rectangle) getType() string {
	return "Rectangle"
}

// Общий интерфейс посетителя

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForRectangle(*rectangle)
}

// Конкретные посетитель - areaCalculator

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(c *circle) {
	fmt.Println("Calculating area for circle")
}

func (a *areaCalculator) visitForRectangle(r *rectangle) {
	fmt.Println("Calculating area for rectangle")
}

// Конкретный посетитель - middleCoordinates

type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s *square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

func (a *middleCoordinates) visitForRectangle(r *rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}

// Поведение

func main() {
	square := &square{side: 2}
	circle := &circle{radius: 3}
	rectangle := &rectangle{l: 2, b: 3}

	areaCalculator := &areaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()

	middleCoordinates := &middleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
