package main

import "fmt"


func main() {
	multiplos3(100);
}

func multiplos3(x int) {
	for i := 1; i < x; i++ {
		if i % 3 == 0 {
			fmt.Println(i);
		}
	}
}