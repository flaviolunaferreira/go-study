package main

import "fmt";


func main() {
	lacoFor(10);
}

func lacoFor(x int) {
	for i := 1; i <= x; i++ {
		if i % 2 == 0 {
			fmt.Println(i, " ", "par")
		} else {
			fmt.Println(i, " ", "impar")
		}
	} 
}