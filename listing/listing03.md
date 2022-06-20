Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
Интерфейс будет равен nil в том и только том случае, если
и тип интерфейса и его значение будут равны nil. Тип
возвращаемого интерфейса - *fs.PathError. Значение - nil.

Вывод: <nil>
       false
       
           type
          /
interface 
          \
           value
            
В пустом интерфейсе и type и value равны nil.
```
