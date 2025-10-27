package main

import "fmt"

func MostrarFor() {
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}
}

/*
	realmente la palabra while
	no existe en go
	se usa el bucle for con una sola condicion, que
	basicamente seria igual que el while
*/
func MostrarWhile() {
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}
