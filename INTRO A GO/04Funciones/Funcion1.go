package main

import "fmt"

func sub(a, b int) (int, int) /*el tipo de retorno*/ {
	return a + b, a - b
}

func main() {

	suma, resta := sub(3, 10)
	fmt.Println(suma)
	fmt.Println(resta)

	//calcularArea

	ancho := 5.0
	alto := 10.0

	area := CalcularArea(ancho, alto)
	fmt.Println("El area del rectangulo es:", area)

}
