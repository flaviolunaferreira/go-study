package main

import "fmt"

var list = []float64{98, 77, 58, 89}

func main() {
	fmt.Println(media(list))
}

func media(lista []float64) float64 { 
	total := 0.0
	for _, valor := range lista {
		total += valor
	}
	return total / float64(len(lista))
}

