package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

// RunCut - нарезает строку по разделителю
// Use: runCut([]string{"-s" `-d"[delim]"` `-f[start-stop]` `"[string]"`})
// Поддерживает флаги:
// -f - "fields" - выбрать поля (колонки)
// -d - "delimiter" - использовать другой разделитель
// -s - "separated" - только строки с разделителем
func RunCut(args []string) string {
	data := ""
	suppress := false
	delim := "\t"
	var start, stop int64
	var err error

	// Проверки на количество аргументов
	if len(args) > 4 {
		return fmt.Sprintf("cut: too much args")
	}
	if len(args) < 2 {
		return fmt.Sprintf("cut: too less args")
	}

	for _, arg := range args {
		switch {
		// Парсим "[string]"
		case regexp.MustCompile(`(^").*("$)`).MatchString(arg):
			data = strings.TrimSuffix(arg, `"`)[1:]
		// Парсим -d"[delim]"
		case regexp.MustCompile(`-d".*"`).MatchString(arg):
			delim = strings.TrimSuffix(arg, `"`)[3:]
		// Парсим ключ -s
		case arg == "-s":
			suppress = true
		// Парсим -f[start-stop]
		case regexp.MustCompile(`-f.*`).MatchString(arg):
			temp := arg[2:]
			if regexp.MustCompile(`^\d+$`).MatchString(temp) {
				start, err = strconv.ParseInt(temp, 10, 64)
				if err != nil {
					return fmt.Sprintf("cut: %s", err.Error())
				}
				stop = start
			} else if regexp.MustCompile(`^\d+-$`).MatchString(temp) {
				temp = strings.TrimSuffix(temp, `-`)
				start, err = strconv.ParseInt(temp, 10, 64)
				if err != nil {
					return fmt.Sprintf("cut: %s", err.Error())
				}
				stop = math.MaxInt64
			} else if regexp.MustCompile(`^-\d+$`).MatchString(temp) {
				temp = temp[1:]
				start = 1
				stop, err = strconv.ParseInt(temp, 10, 64)
				if err != nil {
					return fmt.Sprintf("cut: %s", err.Error())
				}
			} else if regexp.MustCompile(`^\d+-\d+$`).MatchString(temp) {
				start, err = strconv.ParseInt(strings.Split(temp, "-")[0], 10, 64)
				if err != nil {
					return fmt.Sprintf("cut: %s", err.Error())
				}
				stop, err = strconv.ParseInt(strings.Split(temp, "-")[1], 10, 64)
				if err != nil {
					return fmt.Sprintf("cut: %s", err.Error())
				}
			}
		// Если отдан какой-то другой аргумент
		default:
			return fmt.Sprintf("cut: illegal option or input %s", arg)
		}
	}

	res := strings.Split(data, delim)

	// Если строка не имеет заданного разделителя, то мы ее выкидываем
	if suppress && len(res) == 1 {
		return ""
	}

	return strings.Join(res[Min(start-1, int64(len(res))):Min(stop, int64(len(res)))], delim)
}
