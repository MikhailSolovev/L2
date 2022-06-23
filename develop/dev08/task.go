package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stdout, err)
		}
		cmdStr = strings.TrimSuffix(cmdStr, "\n")
		outStr := runCommand(cmdStr)
		if outStr != `` {
			fmt.Fprintln(os.Stdout, outStr)
		}
	}
}

// runCommand - принимает на вход строку, разбивает ее по пробелам и запускает соответствующий обработчик для команды.
func runCommand(cmdStr string) string {
	argsCmdStr := strings.Split(cmdStr, ` `)
	switch argsCmdStr[0] {
	case "\\quit":
		os.Exit(0)
	case "cd":
		return runCd(argsCmdStr[1:])
	case "pwd":
		return runPwd(argsCmdStr[1:])
	case "echo":
		return runEcho(argsCmdStr[1:])
	case "ps":
		return runPs(argsCmdStr[1:])
	case "kill":
		return runKill(argsCmdStr[1:])
	case "cut":
		return runCut(argsCmdStr[1:])
	default:
		return "command not found: " + argsCmdStr[0]
	}
	return ``
}

// runKill - получает PID процессов, которые необходимо убить, затем просматривает все процессы кроме системных
// (PPID = 0 и PPID = 1). Убивает те, у которых PID совпадает с переданными аргументами.
func runKill(args []string) string {
	if len(args) == 0 {
		return "kill: not enough arguments"
	}
	processes, err := ps.Processes()
	if err != nil {
		return err.Error()
	}
	resStr := ""
	for _, arg := range args {
		flag := false
		for _, process := range processes {
			if process.PPid() > 1 && strconv.Itoa(process.Pid()) == arg {
				err = syscall.Kill(process.Pid(), syscall.SIGKILL)
				if err != nil {
					resStr += err.Error()
				}
				flag = true
				break
			}
		}
		if !flag {
			resStr += fmt.Sprintf("kill: kill %s failed: no such process\n", arg)
		}
	}
	return strings.TrimSuffix(resStr, "\n")
}

// runPs - получает PID процессов кроме системных (PPID = 0 и PPID = 1).
func runPs(args []string) string {
	if len(args) != 0 {
		return "ps: too many arguments"
	}
	processes, err := ps.Processes()
	if err != nil {
		return err.Error()
	}
	resStr := "PID\tPPid\tExec\n--------------------\n"
	for _, process := range processes {
		if process.PPid() > 1 {
			resStr += fmt.Sprintf("%d\t%d\t%s\n", process.Pid(), process.PPid(), process.Executable())
		}
	}
	return strings.TrimSuffix(resStr, "\n")
}

// runCd - меняет директорию на заданную или же, если не было передано аргументов, меняет на домашнюю.
func runCd(args []string) string {
	if len(args) > 1 {
		return "cd: too many arguments"
	} else if len(args) == 1 {
		err := os.Chdir(args[0])
		if err != nil {
			return err.Error()
		}
	} else {
		err := os.Chdir(`/Users/ion_mion`)
		if err != nil {
			return err.Error()
		}
	}
	return ``
}

// runPwd - показывает путь от корня до текущей директории.
func runPwd(args []string) string {
	if len(args) != 0 {
		return "pwd: too many arguments"
	}
	dir, err := os.Getwd()
	if err != nil {
		return err.Error()
	}
	return dir
}

// runEcho - вывод переданной строки.
func runEcho(args []string) string {
	return strings.Join(args, ` `)
}

func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

// runCut - нарезает строку по разделителю
// Use: runCut([]string{"-s" `-d"[delim]"` `-f[start-stop]` `"[string]"`})
// Поддерживает флаги:
// -f - "fields" - выбрать поля (колонки)
// -d - "delimiter" - использовать другой разделитель
// -s - "separated" - только строки с разделителем
func runCut(args []string) string {
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
