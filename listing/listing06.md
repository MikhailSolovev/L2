Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
Передается ссылка на slice с len = 3, cap = 3, в функции
добавляется новый элемент в slice, так как размер
нижележащего массива не позволяет добавить четвертый
элемент, то создается новый массив и возвращается ссылка,
дальше изменяется именно новый массив. При выходе из функции
новый массив удаляется GC, так как нет ссылок из outer scope. 

Вывод: [3 2 3]

Слайс состоит из трех частей: ссылки на нижележащий
массив, кол-ва элементов в слайсе, размера нижележащего
массива.

      & (reference on first element of slice)
     /
slice - len (length of slice)  
     \
      cap (size of underlaying array)
```
