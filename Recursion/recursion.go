//Функция fact вызывает сама себя до тех пор пока аргуентом не станет 0: fact(0)
package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7)) //5040
}
