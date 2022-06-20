package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Когда программа содержит несколько объектов, способных обработать тот или иной запрос, однако заранее неизвестно
	какой запрос придет и какой обработчик понадобится.
	Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
	Когда набор объектов, способных обработать запрос, должен задаваться динамически.

	+ Уменьшает зависимость между клиентом и обработчиками
	+ Реализует принцип единственной обязанности
	+ Реализует принцип открытости/закрытости
	- Запрос может остаться никем не обработанным
*/

// Интерфейс сервиса

type Service interface {
	Execute(*Data)
	SetNext(Service)
}
type Data struct {
	// Отметка о приеме данных
	GetSource bool
	// Отметка об обработке
	UpdateSource bool
}

// Конкретная реализация сервиса, сервис отправки данных

type Device struct {
	Name string
	Next Service
}

func (d *Device) Execute(data *Data) {
	// Если данные уже были приняты
	if data.GetSource {
		fmt.Printf("Data from device [%s] already get.\n", d.Name)
		d.Next.Execute(data)
		return
	}
	fmt.Printf("Get data from device [%s].\n", d.Name)
	data.GetSource = true
	d.Next.Execute(data)
}

func (d *Device) SetNext(svc Service) {
	d.Next = svc
}

// Конкретная реализация сервиса, сервис получения/обработки данных

type UpdateDataService struct {
	Name string
	Next Service
}

func (u *UpdateDataService) Execute(data *Data) {
	// Если данные уже были обработаны
	if data.UpdateSource {
		fmt.Printf("Data in service [%s] is already update.\n", u.Name)
		u.Next.Execute(data)
		return
	}
	fmt.Printf("Update data from service [%s].\n", u.Name)
	data.UpdateSource = true
	u.Next.Execute(data)
}

func (u *UpdateDataService) SetNext(svc Service) {
	u.Next = svc
}

// Конкретная реализация сервиса, сервис сохранения данных

type DataService struct {
	Next Service
}

func (d *DataService) Execute(data *Data) {
	// Если данные не были обработаны
	if !data.UpdateSource {
		fmt.Printf("Data not update.\n")
		return
	}
	fmt.Printf("Data save.\n")
}

func (d *DataService) SetNext(svc Service) {
	d.Next = svc
}

// Поведение

func main() {
	// Сервис получения
	device := &Device{Name: "Device-1"}
	// Сервис обработки
	updateSvc := &UpdateDataService{Name: "Update-1"}
	// Сервис сохранения
	dataSvc := &DataService{}

	// Установление цепочки обязаностей
	// Отправка данных с device, передача выполнения на сервис обработки
	device.SetNext(updateSvc)
	// Обновление данных, передача выполнения на сервис сохранения
	updateSvc.SetNext(dataSvc)

	// Создание пакета данных
	data := &Data{}

	// Инициализация цепочки обязанностей
	device.Execute(data)
}
