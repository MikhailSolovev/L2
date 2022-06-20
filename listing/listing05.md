Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Error интерфейс, при вызове функции test() в него попадает
тип *customError со значением nil. Поэтому при проверке на
!= nil получается true.

Вывод: error
```
