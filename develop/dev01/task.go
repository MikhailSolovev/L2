package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func GetPreciseTime(host string) (time.Time, error) {
	t, err := ntp.Time(host)
	if err != nil {
		return t, err
	}
	return t, nil
}

func main() {
	cur, err := GetPreciseTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("%v\nLocal time: %v\n", err, cur.Format(time.UnixDate))
	}
	fmt.Println(cur.Format(time.UnixDate))
}
