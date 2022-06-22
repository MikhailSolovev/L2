package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type StrMatrix struct {
	data [][]string
}

// NewStrMatrix - получает на вход имя файла, возвращает ссылку на объект StrMatrix и ошибку
func NewStrMatrix(name string) (*StrMatrix, error) {
	newStrMatrix := &StrMatrix{}

	file, err := os.Open(name)
	if err != nil {
		return newStrMatrix, err
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
		newStrMatrix.data = append(newStrMatrix.data, strings.Split(scanner.Text(), " "))
	}

	return newStrMatrix, nil
}

// OutputFile - вывод матрицы в файл
func (s *StrMatrix) OutputFile(name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	for _, str := range s.data {
		_, err := file.WriteString(strings.Join(str, " ") + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// OutputStd - вывод матрицы в stdout
func (s *StrMatrix) OutputStd() {
	for _, str := range s.data {
		fmt.Println(strings.Join(str, " "))
	}
}

// Replace - меняет строчки по заданной маске, возвращает true, если строки менялись местами
func (s *StrMatrix) Replace(mask []int) bool {
	newdata := make([][]string, len(s.data))
	var flag bool

	for i, m := range mask {
		if i != m {
			flag = true
			newdata[i] = s.data[m]
		} else {
			newdata[i] = s.data[i]
		}
	}

	s.data = newdata
	return flag
}

// CheckStrInMatrix - проверяет присутсвует ли строка в матрице, если строка присутствует возвращает true
func (s *StrMatrix) CheckStrInMatrix(find []string) bool {
	for _, str := range s.data {
		if strings.Join(str, " ") == strings.Join(find, " ") {
			return true
		}
	}

	return false
}

// Uniq - убирает повторяющиеся строки из матрицы
func (s *StrMatrix) Uniq() {
	newStrMatrix := StrMatrix{}

	for _, str := range s.data {
		if !newStrMatrix.CheckStrInMatrix(str) {
			newStrMatrix.data = append(newStrMatrix.data, str)
		}
	}

	s.data = newStrMatrix.data
}

// sortAlph - сортировка в лексиграфическом порядке, на входе столбец, по которому происходит сортировка, а также
// флаг, если его значение true, то обратный порядок сортировки, если false, то прямой. Флаг trimleft отвечает за
// игнорирование пробелов в начале строки. Возвращает флаг, если его значение true, то матрица изначально была
// НЕ отсортирована, если false, то отсортирована
func (s *StrMatrix) sortAlph(column int, reverse bool, trimleft bool) (rep bool, err error) {
	// Генерация маски индекс-строка
	var mask [][]string
	for i := 0; i < len(s.data); i++ {
		mask = append(mask, []string{s.data[i][column], strconv.Itoa(i)})
	}

	// Прямой порядок сортировки
	if !reverse {
		sort.SliceStable(mask, func(i, j int) bool {
			if !trimleft {
				return mask[i][0] < mask[j][0]
			} else {
				return strings.TrimLeft(mask[i][0], " ") < strings.TrimLeft(mask[j][0], " ")
			}
		})
		// Обратный порядок сортировки
	} else {
		sort.SliceStable(mask, func(i, j int) bool {
			if !trimleft {
				return mask[j][0] <= mask[i][0]
			} else {
				return strings.TrimLeft(mask[j][0], " ") <= strings.TrimLeft(mask[i][0], " ")
			}

		})
	}

	// Генерация маски индекс-индекс
	intmask := make([]int, len(mask))
	for i, el := range mask {
		intmask[i], err = strconv.Atoi(el[1])
		if err != nil {
			return false, err
		}
	}

	// Обмен строчками по заданной маске индекс-индекс
	return s.Replace(intmask), nil
}

// Можно обойтись без этого метода

// sortNum - сортировка по числовому значению строки, на входе столбец, по которому происходит сортировка, а также
// флаг, если его значение true, то обратный порядок сортировки, если false, то прямой. Возвращает флаг, если его
// значение true, то матрица изначально была НЕ отсортирована, если false, то отсортирована
func (s *StrMatrix) sortNum(column int, reverse bool, trimleft bool) (rep bool, err error) {
	// Генерация маски индекс-строка
	var mask [][]string
	for i := 0; i < len(s.data); i++ {
		mask = append(mask, []string{s.data[i][column], strconv.Itoa(i)})
	}

	// Прямой порядок сортировки
	if !reverse {
		sort.SliceStable(mask, func(i, j int) bool {
			numi, err1 := strconv.Atoi(mask[i][0])
			numj, err2 := strconv.Atoi(mask[j][0])
			if err1 != nil && err2 == nil {
				return true
			} else if err1 == nil && err2 != nil {
				return false
			} else if err1 != nil && err2 != nil {
				if !trimleft {
					return mask[i][0] < mask[j][0]
				} else {
					return strings.TrimLeft(mask[i][0], " ") < strings.TrimLeft(mask[j][0], " ")
				}
			} else {
				if numi < numj {
					return true
				} else {
					return false
				}
			}
		})
		// Обратный порядок сортировки
	} else {
		sort.SliceStable(mask, func(i, j int) bool {
			numi, err1 := strconv.Atoi(mask[i][0])
			numj, err2 := strconv.Atoi(mask[j][0])
			if err1 != nil && err2 == nil {
				return false
			} else if err1 == nil && err2 != nil {
				return true
			} else if err1 != nil && err2 != nil {
				if !trimleft {
					return mask[j][0] <= mask[i][0]
				} else {
					return strings.TrimLeft(mask[j][0], " ") <= strings.TrimLeft(mask[i][0], " ")
				}
			} else {
				if numj <= numi {
					return true
				} else {
					return false
				}
			}
		})
	}

	// Генерация маски индекс-индекс
	intmask := make([]int, len(mask))
	for i, el := range mask {
		intmask[i], err = strconv.Atoi(el[1])
		if err != nil {
			return false, err
		}
	}

	// Обмен строчками по заданной маске индекс-индекс
	return s.Replace(intmask), nil
}

// Use sort [-flags] file
func main() {
	// номер столбца (-k int)
	var k int
	flag.IntVar(&k, "k", 1, "define column")
	// сортировка по числовому значению (-n)
	var n bool
	flag.BoolVar(&n, "n", false, "sort by numeric value")
	// сортировка в обратном направлении (-r)
	var r bool
	flag.BoolVar(&r, "r", false, "sort in reverse order")
	// уникальные строки (-u)
	var u bool
	flag.BoolVar(&u, "u", false, "keep only uniq lines")
	// проверить отсортированы ли данные (-c)
	var c bool
	flag.BoolVar(&c, "c", false, "check if data sorted")
	// игнорировать лидирующие пробелы у строк (-b)
	var b bool
	flag.BoolVar(&b, "b", false, "ignore leading spaces")
	// вывод в файл (-o string)
	var o string
	flag.StringVar(&o, "o", "", "output into file")

	flag.Parse()

	m, err := NewStrMatrix(flag.Arg(0))
	if err != nil {
		log.Fatalf(err.Error())
	}

	if u {
		m.Uniq()
	}

	var rep bool
	if n {
		rep, err = m.sortNum(k-1, r, b)
		if err != nil {
			log.Fatalf(err.Error())
		}
	} else {
		rep, err = m.sortAlph(k-1, r, b)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	if c {
		fmt.Println(rep)
	} else {
		if o != "" {
			err = m.OutputFile(o)
			if err != nil {
				log.Fatalf(err.Error())
			}
		} else {
			m.OutputStd()
		}
	}
}
