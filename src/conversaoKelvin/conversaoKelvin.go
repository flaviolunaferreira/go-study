package main

import "fmt"

const kelvin = 0
const celso = 273.15

func main() {
	converte(float64(20))
}

func converte(kelvin float64) {
	fmt.Println(celso - kelvin)
}
