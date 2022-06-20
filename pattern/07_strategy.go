package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
	Когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
	Когда есть множество похожих классов, отличающихся только некоторым поведением.
	Когда не хочется обнажать детали реализации алгоритмов для других классов.
	Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора. Каждая ветка такого
	оператора представляет вариацию алгоритма.

	+ Горячая замена алгоритмов на лету
	+ Изолирует код и данные алгоритмов от остальных классов
	+ Уход от наследования к делегированию
	+ Реализует принцип открытости/закрытости
	- Усложняет программу за счет дополнительных классов
	- Клиент должен знать, в чем разница между стратегиями, чтобы выбрать подходящую
*/

// Общий интерфейс для стратегий

type Strategy interface {
	Route(startPoint int, endPoint int)
}

// Передвижение на машине - конкретная реализиаця стратегии

type RouteStrategy struct {
}

func (r *RouteStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	total := endPoint - startPoint
	totalTime := total * 40 * trafficJam
	fmt.Printf("Road A:[%d] to B: [%d] Avg speed: [%d] Traffic jam: [%d] Total: [%d] Total time: [%d] min\n",
		startPoint, endPoint, avgSpeed, trafficJam, total, totalTime)
}

// Передвижение на общественном транспорте - конкретная реализиаця стратегии

type PublicTransportStrategy struct {
}

func (p *PublicTransportStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 40
	total := endPoint - startPoint
	totalTime := total * 40
	fmt.Printf("PublicTransport A:[%d] to B: [%d] Avg speed: [%d] Total: [%d] Total time: [%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}

// Передвижение пешком - конкретная реализиаця стратегии

type WalkStrategy struct {
}

func (p *WalkStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 4
	total := endPoint - startPoint
	totalTime := total * 60
	fmt.Printf("Walk A:[%d] to B: [%d] Avg speed: [%d] Total: [%d] Total time: [%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}

// Навигатор - общий контекст

type Navigator struct {
	Strategy
}

func (n *Navigator) SetStrategy(str Strategy) {
	n.Strategy = str
}

// Поведение

var (
	start      = 10
	end        = 100
	strategies = []Strategy{
		&PublicTransportStrategy{},
		&RouteStrategy{},
		&WalkStrategy{},
	}
)

func main() {
	// Контекст
	nav := Navigator{}
	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)
	}
}
