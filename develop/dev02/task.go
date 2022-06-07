package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Unpacker interface {
	Unpack() (string, error)
}

type PackedString string

func (s PackedString) Unpack() (string, error) {
	var lastRune, lastLetter rune
	var result, num strings.Builder
	var esc bool

	result.Reset()
	num.Reset()

	lastRune = 0
	lastLetter = 0

	for i, curRune := range s {
		// if first character is num, wrong input
		if i == 0 && unicode.IsDigit(curRune) {
			return "", errors.New("error: first character is digit")
		}
		// writing to result and unpacking previous sequence
		if unicode.IsLetter(curRune) {
			// if letter after digit
			if unicode.IsDigit(lastRune) {
				numRunes, err := strconv.Atoi(num.String())
				if err != nil {
					return "", err
				}
				for j := 0; j < numRunes-1; j++ {
					result.WriteRune(lastLetter)
				}
				num.Reset()
			}
			result.WriteRune(curRune)
			lastLetter = curRune
			lastRune = curRune
		}
		// write to num or write digit to result
		if unicode.IsDigit(curRune) {
			// escape sequence
			if esc {
				result.WriteRune(curRune)
				lastLetter = curRune
				lastRune = curRune
				esc = false
			} else {
				// if digit after letter
				if unicode.IsLetter(lastRune) {
					num.Reset()
				}
				num.WriteRune(curRune)
				lastRune = curRune
				// last digit in input string
				if i == utf8.RuneCountInString(string(s))-1 {
					numRunes, err := strconv.Atoi(num.String())
					if err != nil {
						return "", err
					}
					for j := 0; j < numRunes-1; j++ {
						result.WriteRune(lastLetter)
					}
				}
			}
		}
		if curRune == '\\' {
			if esc {
				result.WriteRune(curRune)
				lastLetter = curRune
				lastRune = curRune
				esc = false
			} else {
				if i == utf8.RuneCountInString(string(s))-1 {
					return "", errors.New("error: last character is \\")
				}
				esc = true
				lastRune = curRune
			}
		}
	}

	return result.String(), nil
}
