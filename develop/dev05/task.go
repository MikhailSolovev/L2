package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Data struct {
	data []string
}

// contains - вспомогательная функция проверки: содержится ли элемент в срезе
func contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

// Out - вывод подходящих строк в STDOUT
func (d Data) Out(c, n bool, aft, bef int, idxs []int) {
	var alreadyOutput []int
	// Обработка с, только вывод кол-ва подходящих строк
	if c {
		fmt.Println(len(idxs))
	} else {
		for _, i := range idxs {
			ia := i + aft
			if ia >= len(d.data) {
				ia = len(d.data) - 1
			}
			ib := i - bef
			if ib < 0 {
				ib = 0
			}
			for j := ib; j <= ia; j++ {
				if contains(alreadyOutput, j) {
					continue
				}
				// Обработка n, вывод номера строки
				if n {
					fmt.Printf("%d:%s\n", j+1, d.data[j])
				} else {
					fmt.Println(d.data[j])
				}
				alreadyOutput = append(alreadyOutput, j)
			}
		}
	}
}

// Search - функция поиска принимает на вход ключи: i - игнорирует регистр, v - ищет строки, НЕ удовлетворяющие
// параметрам поиска, f - нахождение подстроки в строке (без обработки регулярных выражений). Также принимает
// pattern, который ищется в строках. Возвращает срез индексов строк, которые удовлетворяют условиям поиска.
func (d Data) Search(i, v, f bool, pattern string) ([]int, error) {
	var idxs []int
	for idx, line := range d.data {
		// Обработка i, строки приводятся к нижнему регистру, а также сам паттерн
		if i {
			line = strings.ToLower(line)
			pattern = strings.ToLower(pattern)
		}
		// Обработка паттерна как обычной подстроки
		if f {
			if v {
				if !strings.Contains(line, pattern) {
					idxs = append(idxs, idx)
				}
			} else {
				if strings.Contains(line, pattern) {
					idxs = append(idxs, idx)
				}
			}
			// Обработка регулярок
		} else {
			re, err := regexp.Compile(pattern)
			if err != nil {
				return idxs, err
			}
			if v {
				if !re.MatchString(line) {
					idxs = append(idxs, idx)
				}
			} else {
				if re.MatchString(line) {
					idxs = append(idxs, idx)
				}
			}
		}
	}
	return idxs, nil
}

// NewData - получает на вход имя файла, возвращает ссылку на объект Data и ошибку
func NewData(name string) (*Data, error) {
	newData := &Data{}

	file, err := os.Open(name)
	if err != nil {
		return newData, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		newData.data = append(newData.data, scanner.Text())
	}

	return newData, nil
}

// Use grep [-flags] "pattern" file
func main() {
	// игнорирование регистра (-i)
	var i bool
	flag.BoolVar(&i, "i", false, "ignore register")
	// инверсия вывода (-v)
	var v bool
	flag.BoolVar(&v, "v", false, "invert output")
	// поиск по "чистым" строкам (-F)
	var F bool
	flag.BoolVar(&F, "F", false, "off regex")
	// вывод только кол-ва строк (-с)
	var c bool
	flag.BoolVar(&c, "c", false, "count of matched strings")
	// вывод номеров строк (-n)
	var n bool
	flag.BoolVar(&n, "n", false, "print index of line")
	// вывод строк ДО (-B int)
	var B int
	flag.IntVar(&B, "B", 0, "print lines before match")
	// вывод строк ПОСЛЕ (-A int)
	var A int
	flag.IntVar(&A, "A", 0, "print lines after match")
	// вывод контекста (-C int)
	var C int
	flag.IntVar(&C, "C", 0, "print context around match")

	flag.Parse()

	d, err := NewData(flag.Arg(1))
	if err != nil {
		log.Fatalf(err.Error())
	}

	idxs, err := d.Search(i, v, F, flag.Arg(0))
	if err != nil {
		log.Fatalf(err.Error())
	}

	if C != 0 {
		d.Out(c, n, C, C, idxs)
	} else {
		d.Out(c, n, A, B, idxs)
	}
}
