package main

import "fmt"

func texto(c chan string) {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			c <- "Ping";
		} else {
			c <- "Pong";
		}
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c;
		fmt.Println(msg);
	}	
}

func main() {
	var c chan string = make(chan string);

	go texto(c);
	go imprimir(c);

}

