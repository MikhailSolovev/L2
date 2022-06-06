package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Когда нужно создавать разные представления какого-то объекта.
	Когда нужно собирать сложные составные объекты.

	+ Позволяет создавать продукты пошагово
	+ Позволяет использовать один и тот же код для создания различных продуктов
	+ Изолирует сложный код сборки продукта от его основной бизнес-логики
	- Усложняет код программы за счет дополнительных классов
	- Клиент будет привязан к конкретным классам строителей, так как в интерфейсе строителя
	  может не быть метода получения результата
*/

// Типы сборщиков/строителей

type AsusCollector struct {
	CPU     int
	Brand   string
	RAM     int
	GPU     int
	Monitor bool
}

func (collector *AsusCollector) SetCPU() {
	collector.CPU = 4
}

func (collector *AsusCollector) SetBrand() {
	collector.Brand = "Asus"
}

func (collector *AsusCollector) SetRAM() {
	collector.RAM = 8
}

func (collector *AsusCollector) SetGPU() {
	collector.GPU = 1
}

func (collector *AsusCollector) SetMonitor() {
	collector.Monitor = true
}

func (collector *AsusCollector) GetComputer() Computer {
	return Computer{
		CPU:     collector.CPU,
		Brand:   collector.Brand,
		RAM:     collector.RAM,
		GPU:     collector.GPU,
		Monitor: collector.Monitor,
	}
}

type HpCollector struct {
	CPU     int
	Brand   string
	RAM     int
	GPU     int
	Monitor bool
}

func (collector *HpCollector) SetCPU() {
	collector.CPU = 4
}

func (collector *HpCollector) SetBrand() {
	collector.Brand = "Hp"
}

func (collector *HpCollector) SetRAM() {
	collector.RAM = 16
}

func (collector *HpCollector) SetGPU() {
	collector.GPU = 2
}

func (collector *HpCollector) SetMonitor() {
	collector.Monitor = true
}

func (collector *HpCollector) GetComputer() Computer {
	return Computer{
		CPU:     collector.CPU,
		Brand:   collector.Brand,
		RAM:     collector.RAM,
		GPU:     collector.GPU,
		Monitor: collector.Monitor,
	}
}

// Интерфейс сборщика/строителя

const (
	AsusCollectorType = "asus"
	HpCollectorType   = "hp"
)

type Collector interface {
	SetCPU()
	SetBrand()
	SetRAM()
	SetGPU()
	SetMonitor()
	GetComputer() Computer
}

func GetCollector(CollectorType string) Collector {
	switch CollectorType {
	default:
		return nil
	case AsusCollectorType:
		return &AsusCollector{}
	case HpCollectorType:
		return &HpCollector{}
	}
}

// Объект

type Computer struct {
	CPU     int
	Brand   string
	RAM     int
	GPU     int
	Monitor bool
}

func (c Computer) Print() {
	fmt.Printf("%s CPU:[%d] RAM:[%d] GPU:[%d] Monitor:[%v]\n", c.Brand, c.CPU, c.RAM, c.GPU, c.Monitor)
}

// Фабрика/Директор

type Factory struct {
	Collector Collector
}

func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

func (factory *Factory) SetCollector(collector Collector) {
	factory.Collector = collector
}

// CreateComputer - основной метод по созданию объектов
func (factory *Factory) CreateComputer() Computer {
	factory.Collector.SetCPU()
	factory.Collector.SetBrand()
	factory.Collector.SetRAM()
	factory.Collector.SetGPU()
	factory.Collector.SetMonitor()
	return factory.Collector.GetComputer()
}

// Поведение

func main() {
	asusCollector := GetCollector("asus")
	hpCollector := GetCollector("hp")

	factory := NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	factory.SetCollector(hpCollector)
	hpComputer := factory.CreateComputer()
	hpComputer.Print()
}
