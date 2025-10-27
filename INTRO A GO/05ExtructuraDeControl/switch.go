package main

import (
	"fmt"
	"time"
)

func MostrarSwitch() {
	x := 13

	switch {
	case x < 0:
		fmt.Println("Negativo")
	case x == 0:
		fmt.Println("Cero")
	case x > 0 && x < 100:
		fmt.Println("Positivo menor que 100")
	default:
		fmt.Println("Grande")
	}
}

func QueTiempoEs() {
	hora := time.Now().Hour()

	fmt.Println("Hora actual:", hora)

	switch {
	case hora >= 6 && hora < 12:
		fmt.Println("Buenos dias")
	case hora >= 12 && hora < 18:
		fmt.Println("Buenas tardes")
	default:
		fmt.Println("Buenas Noches	")
	}
}
