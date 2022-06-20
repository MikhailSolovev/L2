package pattern

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
	*** В отличие от стратегии контекст и конретные состояния могут знать друг о друге и инициировать переходы
		от одного состояния к другому

	Когда есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния. Причем типов
	состояний много и их код часто меняется.
	Когда код класса содержит множество больших, похожих друг на друга, условных операторов, которые выбирают поведения
	в зависимости от текущих значений полей класса.
	Когда сознательно используется табличная машина состояний, построенная на условных операторах, но вынуждены
	мириться с дублированием кода для похожих состояний и переходов.

	+ Избавляет от множества больших условных операторов машины состояний
	+ Концентрирует в одном месте код, связанный с определенным состоянием
	+ Упрощает код контекста
	- Может неоправданно усложнить код, если состояний мало и они редко меняются
*/

// Общий интерфейс состояния

type State interface {
	// AddItem - добавление товара
	AddItem(int) error
	// RequestItem - запрос ассортимента
	RequestItem() error
	// InsertMoney - занести деньги
	InsertMoney(money int) error
	// DispenseItem - получение товара пользователем
	DispenseItem() error
}

// Нет товара - конкретное состояние (доступна только функция AddItem)

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) RequestItem() error {
	return fmt.Errorf("item out of stock")
}

func (i *NoItemState) AddItem(count int) error {
	i.vendingMachine.IncrementItemCount(count)
	i.vendingMachine.SetState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *NoItemState) DispenseItem() error {
	return fmt.Errorf("item out of stock")
}

// Запрос товара - конкретное состояние (доступна только функция InsertMoney)

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *ItemRequestedState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *ItemRequestedState) InsertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less, please insert [%d]", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is OK")
	i.vendingMachine.SetState(i.vendingMachine.hasMoney)
	return nil
}

func (i *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("please insert money first")
}

// В автомат внесли деньги - конкретное состояние (доступная только функция DispenseItem)

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) RequestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing item")
	i.vendingMachine.itemCount -= 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.SetState(i.vendingMachine.hasItem)
	}
	return nil
}

// В автомате есть товар - конкретное состояние (доступны функции RequestItem и AddItem)

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (i *HasItemState) RequestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Println("Item requested")
	i.vendingMachine.SetState(i.vendingMachine.itemRequested)
	return nil
}

func (i *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.IncrementItemCount(count)
	return nil
}

func (i *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (i *HasItemState) DispenseItem() error {
	return fmt.Errorf("please select item first")
}

// Вендинговая машина с одним товаром (контекст)

type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	currentState  State
	itemCount     int
	itemPrice     int
}

// NewVendingMachine - консктруктор для вендингового автомата
func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	// Инициализация состояний
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}
	// Первоначальное состояние машины
	v.SetState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

// SetState - установка текущего состяния
func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

func (v *VendingMachine) IncrementItemCount(count int) {
	v.itemCount += count
}

// Поведение

func main() {
	vendingMachine := NewVendingMachine(1, 10)

	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()

	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
