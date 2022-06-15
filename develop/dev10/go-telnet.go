package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type Client interface {
	Connect() error
}

type TelnetClient struct {
	conn    net.Conn
	host    string
	port    string
	timeout time.Duration
}

func NewTelnetClient(host string, port string, timeout int) *TelnetClient {
	return &TelnetClient{
		host:    host,
		port:    port,
		timeout: time.Duration(timeout) * time.Second,
	}
}

func (t *TelnetClient) Connect() (err error) {
	// Установление соединения
	t.conn, err = net.DialTimeout("tcp", net.JoinHostPort(t.host, t.port), t.timeout)
	if err != nil {
		return err
	}
	// Закрытие соединения + обработка ошибки
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}(t.conn)

	return t.goClientWriter()
}

func (t *TelnetClient) goClientWriter() (err error) {
	clientReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(t.conn)
	for {
		// Запись в socket из STDIN
		fmt.Print(">> ")
		clientRequest, err := clientReader.ReadString('\n')

		switch err {
		case nil:
			if _, err = fmt.Fprintf(t.conn, clientRequest); err != nil {
				return err
			}
		// Клиент закрыл соединение CTRL+D
		case io.EOF:
			return err
		default:
			return err
		}

		// Получение ответа из socket в STDOUT
		fmt.Print("-> ")
		serverResponse, err := serverReader.ReadString('\n')

		switch err {
		case nil:
			fmt.Printf("%s", serverResponse)
		// Сервер закрыл соединение CTRL+C
		case io.EOF:
			return err
		default:
			return err
		}
	}
}

// Usage ./go-telnet.go -timeout=10 host port
func main() {
	var timeout int
	flag.IntVar(&timeout, "timeout", 10, "Timeout for connect")
	flag.Parse()

	client := NewTelnetClient(flag.Arg(0), flag.Arg(1), timeout)
	err := client.Connect()
	if err != nil {
		fmt.Println(err)
	}
}
