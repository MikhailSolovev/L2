package main

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// P.s хэш решил вычислять, чтобы не использовать метод reflect.DeepEqual(m1, m2)

// asSha256 - считает hash sum для любого объекта
func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

// CountLetters - подсчитывает кол-во каждого из символов в строке
func CountLetters(str string) map[rune]uint16 {
	res := map[rune]uint16{}
	word := []rune(str)
	for _, letter := range word {
		if _, ok := res[letter]; ok {
			res[letter] += 1
		} else {
			res[letter] = 1
		}
	}
	return res
}

// Unique - убирает повторы в slice
func Unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func FindAnagrams(words []string) map[string][]string {
	anagrams := map[string][]string{}
	for _, word := range words {
		word = strings.ToLower(word)
		hash := asSha256(CountLetters(word))
		anagrams[hash] = append(anagrams[hash], word)
	}
	// Замена hash sum на первый элемент списка, удаление ключей, значения которых пустой слайс,
	// убирает повторы, cортирует по возврастанию
	temp := map[string][]string{}
	for _, v := range anagrams {
		if len(v[1:]) != 0 {
			newV := Unique(v[1:])
			sort.Strings(newV)
			temp[v[0]] = newV
		}
	}
	anagrams = temp

	return anagrams
}
