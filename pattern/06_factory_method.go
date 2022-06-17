package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать код.
	Когда нужно дать возможность пользователям расширять части фреймворка или библиотеки.
	Когда нужно сэкономить системные ресурсы, повторно используя уже созданные объекты, вместо
	создания новых.

	+ Избавляет класс от привязки к конкретным классам продуктов
	+ Выделяет код производства продуктов в одно место, упрощая поддержку кода
	+ Упрощает добавление новых продуктов в программу
	+ Реализует принцип открытости/закрытости
	- Может привести к созданию больших параллельных иерархий классов, так как для каждого класса
	  продукта надо создать свой подкласс создателя
*/

const (
	ServerType           = "server"
	PersonalComputerType = "personal"
	NotebookType         = "notebook"
)

// Интерфейс Computer

type Comp interface {
	GetType() string
	PrintDetails()
}

// Конструктор для интерфейса (наш главный метод, фабрика)

func New(typeName string) Comp {
	switch typeName {
	default:
		fmt.Printf("%s несуществующий тип объекта\n", typeName)
		return nil
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	case NotebookType:
		return NewNotebook()
	}
}

// Конкретный объект Server, реализующий интерфейс Computer

type Server struct {
	Type   string
	Core   int
	Memory int
}

// Конструктор для Server

func NewServer() Comp {
	return Server{
		Type:   ServerType,
		Core:   16,
		Memory: 256,
	}
}

func (s Server) GetType() string {
	return s.Type
}

func (s Server) PrintDetails() {
	fmt.Printf("%s Core:[%d] Mem:[%d]\n", s.Type, s.Core, s.Memory)
}

// Конкретный объект PersonalComputer, реализующий интерфейс Computer

type PersonalComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

// Конструктор для PersonalComputer

func NewPersonalComputer() Comp {
	return PersonalComputer{
		Type:    PersonalComputerType,
		Core:    8,
		Memory:  16,
		Monitor: true,
	}
}

func (p PersonalComputer) GetType() string {
	return p.Type
}

func (p PersonalComputer) PrintDetails() {
	fmt.Printf("%s Core:[%d] Mem:[%d] Monitor:[%t]\n", p.Type, p.Core, p.Memory, p.Monitor)
}

// Конкретный объект Notebook, реализующий интерфейс Computer

type Notebook struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

// Конструктор для Notebook

func NewNotebook() Comp {
	return Notebook{
		Type:    NotebookType,
		Core:    4,
		Memory:  8,
		Monitor: true,
	}
}

func (n Notebook) GetType() string {
	return n.Type
}

func (n Notebook) PrintDetails() {
	fmt.Printf("%s Core:[%d] Mem:[%d] Monitor:[%t]\n", n.Type, n.Core, n.Memory, n.Monitor)
}

// Поведение

var types = []string{PersonalComputerType, NotebookType, ServerType, "monoblock"}

func main() {
	for _, typeName := range types {
		computer := New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}
}
