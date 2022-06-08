package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Когда нужно параметризовать объекты выполняемым действием.
	Когда нужно ставить операции в очередь, выполнять их по расписанию или передавать по сети.
	Когда нужна операция отмены.

	+ Убирает прямую зависимость между объектами, вызывающими операции и объектами, которые их непосредственно
	  выполняют
	+ Позволяет реализовать простую отмену и повтор операции
	+ Позволяет реализовать отложенный запуск команд
	+ Позволяет собирать сложные команды из простых
	+ Реализует принцип открытости/закрытости
	- Усложняет код программы за счет дополнительных классов
*/

// Отправитель

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

// Интерфейс команды

type command interface {
	execute()
}

// Конкретные команды

// Включение TV

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

// Выключение TV

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

// Интерфейс получателя

type device interface {
	on()
	off()
}

// Конкретный получатель

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning TV on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning TV off")
}

// Поведение

func main() {
	tv := &tv{}

	onCommand := &onCommand{device: tv}
	offCommand := &offCommand{device: tv}

	onButton := &button{command: onCommand}
	onButton.press()

	offButton := &button{command: offCommand}
	offButton.press()
}
