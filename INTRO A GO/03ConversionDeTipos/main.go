package main

import "fmt"

func main() {

	/* conversiones
	se pierden bits que no caben basicamente se desborda
	*/

	var sueldo uint16 = 300
	var b uint8 = uint8(sueldo) //conversion

	var sueldoPersona2 uint8 = 250
	var c int8 = int8(sueldoPersona2)

	//.....

	fmt.Println(b)
	fmt.Println(c)
}
