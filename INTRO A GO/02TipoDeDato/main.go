package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//bool
	var esMayorDeEdad bool = true
	fmt.Println("Es mayor de edad:", esMayorDeEdad)

	// string
	Nombre := "Joel"
	fmt.Println("Su nombre es:", Nombre)

	// int
	var edad, edad2 int = 19, 22
	fmt.Println(Nombre, "Tiene", edad, "y Jose tiene", edad2)

	// float

	const altura float32 = 1.50
	fmt.Println("Jose mide", altura, "metros")

	var pi float64 = 3.1416
	fmt.Println("pi", pi)

	// runes (funciona para emojis tambien )
	const inicialNombre rune = 'J'
	fmt.Println(inicialNombre)

	// saber el length de un string
	length := utf8.RuneCountInString(Nombre)
	fmt.Println(length)

	// el len es para saber la cantidad de byte
	fmt.Print(length)

	// para profundizar mas: https://go.dev/ref/spec

}
