package pattern

import (
	"errors"
	"fmt"
	"time"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Магазин

type Product struct {
	Name  string
	Price float32
}

type Shop struct {
	Name     string
	Products []Product
}

// Sell - фасад над всеми остальными объектами
func (s Shop) Sell(user User, product string) error {
	fmt.Println("[Магазин] Запрос к пользователю для получения остатка по карте")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка - может ли [%s] пользователь купить товар\n", user.Name)
	time.Sleep(time.Millisecond * 500)
	for _, prod := range s.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средств для покупки товара")
		}
		fmt.Printf("[Магазин] Товар [%s] - куплен\n", prod.Name)
	}
	return nil
}

// Карта

type Card struct {
	Name    string
	Balance float32
	Bank    *Bank
}

func (c Card) CheckBalance() error {
	fmt.Println("[Карта] Запрос в банк для проверки остатка")
	time.Sleep(time.Millisecond * 800)
	return c.Bank.CheckBalance(c.Name)
}

// Банк

type Bank struct {
	Name  string
	Cards []Card
}

func (b Bank) CheckBalance(CardNumber string) error {
	fmt.Printf("[Банк] Получение остатка по карте %s\n", CardNumber)
	time.Sleep(time.Millisecond * 300)
	for _, card := range b.Cards {
		if card.Name != CardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств")
		}
	}
	fmt.Println("[Банк] Остаток положительный")
	return nil
}

// Пользователь

type User struct {
	Name string
	Card *Card
}

func (u User) GetBalance() float32 {
	return u.Card.Balance
}

// Поведение

var (
	bank = Bank{
		Name:  "БАНК",
		Cards: []Card{},
	}
	card1 = Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = Card{
		Name:    "CRD-2",
		Balance: 5,
		Bank:    &bank,
	}
	user1 = User{
		Name: "Покупатель-1",
		Card: &card1,
	}
	user2 = User{
		Name: "Покупатель-2",
		Card: &card2,
	}
	prod = Product{
		Name:  "Сыр",
		Price: 150,
	}
	shop = Shop{
		Name: "SHOP",
		Products: []Product{
			prod,
		},
	}
)

func main() {
	fmt.Println("[Банк] Выпуск карт")
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Printf("[%s]\n", user1.Name)
	err := shop.Sell(user1, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("[%s]\n", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
