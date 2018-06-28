package main

import (
	"fmt"
)

func main() {

   //Четное или нечетное число 7
   if 7%2 == 0 {
	   fmt.Println("7 is even")
   } else {
	   fmt.Println("7 is odd")   
   }

   //Else с доп условием
   num := 9
   if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}


}