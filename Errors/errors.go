package main

import (
	"errors"
	"fmt"
)

//Вариант 1
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

//Вариант 2
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	//Вариант 1
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 faled:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	//Ответ:
	//f1 worked: 10
	//f1 faled: can't work with 42

	//Вариант 2.1
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	//Ответ:
	//f2 worked: 10
	//f2 failed: 42 - can't work with it

	//Вариант 2.2
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
	//Ответ:
	//42
	//can't work with it
}
