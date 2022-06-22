package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“done after %v”, time.Since(start))
*/

// asChan - создает канал (односторонний на чтение) из произвольного числа аргументов
func asChan(vs ...interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		for _, v := range vs {
			c <- v
		}
		close(c)
	}()
	return c
}

// merge - сливает n-каналов в один
//
//	go1	c1---->out
//	go2	c2---->out
//	go3 c3---->out
//	...
// Каждая отдельная горутина сливает данные из своего канала в один общий канал
func merge(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		var wg sync.WaitGroup

		wg.Add(len(channels))
		for _, c := range channels {
			go func(c <-chan interface{}) {
				for v := range c {
					out <- v
				}
				wg.Done()
			}(c)
		}

		wg.Wait()
		close(out)
	}()

	return out
}

// mergeReflect - реализация merge через пакет рефлексии, используется только одна горутина
func mergeReflect(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		var cases []reflect.SelectCase

		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		for len(cases) > 0 {
			// v - reflect value
			i, v, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface()
		}
	}()

	return out
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	done := make(chan struct{})

	wg.Add(1)
	for _, c := range channels {
		go func(c <-chan interface{}, done chan struct{}) {
			for {
				select {
				case <-done:
					return
				case _, ok := <-c:
					if !ok {
						done <- struct{}{}
					}
				default:
					time.Sleep(1 * time.Second)
				}
			}
		}(c, done)
	}

	go func() {
		for {
			select {
			case <-done:
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()
	return merge(channels...)
}

func main() {
	start := time.Now()
	a := asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := asChan(13, 14, 15, 18, 19)
	c := asChan(20, 21, 22, 23, 24, 25, 26, 27, 28, 29)

	for v := range or(a, b, c) {
		fmt.Println(v)
	}
	fmt.Printf("done after %v", time.Since(start))
}
