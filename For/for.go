package main

import (
	"fmt"
)

func main() {

	//Вариант 1: вывод цифр от 1 до 3
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//Вариант 2: вывод цифр от 7 до 9
	for j:= 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//Вариант 3: Цикл без условия. Повторяется пока не будет выполнена команда выхода. 
	for {
		fmt.Println("loop")
		break
	}

	//Вариант 3: Игнорирование команд после if
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}